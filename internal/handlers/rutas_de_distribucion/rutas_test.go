package rutas_de_distribucion_test

// Test — Handler con httptest + fake en memoria (usando testify).
//
// Qué comprueba:
//   a) Ruta protegida SIN token → 401 Unauthorized.
//   b) Token expirado → 401 Unauthorized.
//   c) POST /rutas con datos válidos y token → 201 Created.
//   d) POST /rutas con nombre vacío → 400 Bad Request.

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"Pesca_Directa_AplicacionesWeb_II/internal/handlers/rutas_de_distribucion"
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/models"

	// Imports con Dot para inyectar tus servicios y repositorios sin prefijos molestos
	. "Pesca_Directa_AplicacionesWeb_II/internal/service"
	. "Pesca_Directa_AplicacionesWeb_II/internal/service/rutas_de_distribucion"
	storage "Pesca_Directa_AplicacionesWeb_II/internal/storage/rutas_de_distribucion"
)

// ─── Fake de UserRepository ───────────────────────────────────────────────────

type usuarioFakeRutas struct {
	datos  []models.Usuario
	nextID int
}

func nuevoUsuarioFakeRutas() *usuarioFakeRutas {
	return &usuarioFakeRutas{nextID: 1}
}

func (u *usuarioFakeRutas) CrearUsuario(usr models.Usuario) (models.Usuario, error) {
	usr.ID = u.nextID
	u.nextID++
	u.datos = append(u.datos, usr)
	return usr, nil
}
func (u *usuarioFakeRutas) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	for _, usr := range u.datos {
		if usr.Email == email {
			return usr, true
		}
	}
	return models.Usuario{}, false
}

// ─── Helper: genera un token JWT válido para los tests ───────────────────────

func generarTokenTestRutas(t *testing.T, auth *AuthService) string {
	t.Helper()
	u := models.Usuario{ID: 1, Email: "test@rutas.com", PasswordHash: "x"}
	token, err := auth.GenerarToken(u)
	require.NoError(t, err, "no se pudo generar token de prueba")
	return "Bearer " + token
}

// ─── Helper: construye el router completo en memoria ─────────────────────────

func nuevoRouterTestRutas(t *testing.T) (http.Handler, *AuthService) {
	t.Helper()

	authService := NewAuthService(nuevoUsuarioFakeRutas())
	rutasService := NewRutasService(storage.NuevaMemoriaRutas())

	// Inicializamos el servidor global pasándole tus servicios locales simulados
	servidor := rutas_de_distribucion.NewServer0(rutasService, authService)

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			// Rutas
			r.Get("/rutas", servidor.ListarRutas)
			r.Post("/rutas", servidor.CrearRuta)
			r.Get("/rutas/{id}", servidor.ObtenerRuta)
			r.Put("/rutas/{id}", servidor.ActualizarRuta)
			r.Delete("/rutas/{id}", servidor.BorrarRuta)

			// Puntos
			r.Get("/puntos", servidor.ListarPuntos)
			r.Post("/puntos", servidor.CrearPunto)
			r.Get("/puntos/{id}", servidor.ObtenerPunto)
			r.Put("/puntos/{id}", servidor.ActualizarPunto)
			r.Delete("/puntos/{id}", servidor.BorrarPunto)

			// Transportistas
			r.Get("/transportistas", servidor.ListarTransportistas)
			r.Post("/transportistas", servidor.CrearTransportista)
			r.Get("/transportistas/{id}", servidor.ObtenerTransportista)
			r.Put("/transportistas/{id}", servidor.ActualizarTransportista)
			r.Delete("/transportistas/{id}", servidor.BorrarTransportista)

			// Entregas
			r.Get("/entregas", servidor.ListarEntregas)
			r.Post("/entregas", servidor.CrearEntrega)
			r.Get("/entregas/{id}", servidor.ObtenerEntrega)
			r.Put("/entregas/{id}", servidor.ActualizarEntrega)
			r.Delete("/entregas/{id}", servidor.BorrarEntrega)
		})
	})

	return r, authService
}

// ─── Tests ────────────────────────────────────────────────────────────────────

// Test 401 — GET /rutas SIN token debe devolver 401.
func TestHandler_GetRutas_SinToken_Devuelve401(t *testing.T) {
	router, _ := nuevoRouterTestRutas(t)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/rutas", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

// Test 401 — Token expirado debe devolver 401.
func TestHandler_RutasTokenExpirado_Devuelve401(t *testing.T) {
	router, _ := nuevoRouterTestRutas(t)

	secretJWT := []byte("rutas-distribucion-tarqui-secret-2026")
	claims := &Claims{
		UsuarioID: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)),
		},
	}
	tokenExpirado, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretJWT)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/rutas", nil)
	req.Header.Set("Authorization", "Bearer "+tokenExpirado)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

// Test 200 — GET /rutas/{id} con ID existente → 200 OK.
func TestHandler_ObtenerRuta_Existente_Devuelve200(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.Ruta{Nombre: "Ruta A", Origen: "Puerto", Destino: "Mercado"})
	reqCrear := httptest.NewRequest(http.MethodPost, "/api/v1/rutas", bytes.NewReader(cuerpo))
	reqCrear.Header.Set("Content-Type", "application/json")
	reqCrear.Header.Set("Authorization", token)
	recCrear := httptest.NewRecorder()
	router.ServeHTTP(recCrear, reqCrear)
	require.Equal(t, http.StatusCreated, recCrear.Code)

	var creada models.Ruta
	require.NoError(t, json.Unmarshal(recCrear.Body.Bytes(), &creada))

	req := httptest.NewRequest(http.MethodGet, "/api/v1/rutas/"+strconv.FormatUint(uint64(creada.ID), 10), nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "body: %s", rec.Body.String())
}

// Test 400 — GET /rutas/{id} con ID no numérico → 400 Bad Request.
func TestHandler_ObtenerRuta_IDInvalido_Devuelve400(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/rutas/abc", nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code, "body: %s", rec.Body.String())
}

// Test 404 — GET /rutas/{id} con ID inexistente → 404 Not Found.
func TestHandler_ObtenerRuta_Inexistente_Devuelve404(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/rutas/9999", nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code, "body: %s", rec.Body.String())
}

// Test 200 — PUT /rutas/{id} con datos válidos → 200 OK.
func TestHandler_ActualizarRuta_Valida_Devuelve200(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.Ruta{Nombre: "Ruta B", Origen: "Puerto", Destino: "Mercado"})
	reqCrear := httptest.NewRequest(http.MethodPost, "/api/v1/rutas", bytes.NewReader(cuerpo))
	reqCrear.Header.Set("Content-Type", "application/json")
	reqCrear.Header.Set("Authorization", token)
	recCrear := httptest.NewRecorder()
	router.ServeHTTP(recCrear, reqCrear)
	require.Equal(t, http.StatusCreated, recCrear.Code)

	var creada models.Ruta
	require.NoError(t, json.Unmarshal(recCrear.Body.Bytes(), &creada))

	actualizacion, _ := json.Marshal(models.Ruta{Nombre: "Ruta B Editada", Origen: "Puerto", Destino: "Mercado"})
	req := httptest.NewRequest(http.MethodPut, "/api/v1/rutas/"+strconv.FormatUint(uint64(creada.ID), 10), bytes.NewReader(actualizacion))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "body: %s", rec.Body.String())
}

// Test 404 — PUT /rutas/{id} con ID inexistente → 404 Not Found.
func TestHandler_ActualizarRuta_Inexistente_Devuelve404(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.Ruta{Nombre: "X", Origen: "Y", Destino: "Z"})
	req := httptest.NewRequest(http.MethodPut, "/api/v1/rutas/9999", bytes.NewReader(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code, "body: %s", rec.Body.String())
}

// Test 204 — DELETE /rutas/{id} con ID existente → 204 No Content.
func TestHandler_BorrarRuta_Existente_Devuelve204(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.Ruta{Nombre: "Ruta C", Origen: "Puerto", Destino: "Mercado"})
	reqCrear := httptest.NewRequest(http.MethodPost, "/api/v1/rutas", bytes.NewReader(cuerpo))
	reqCrear.Header.Set("Content-Type", "application/json")
	reqCrear.Header.Set("Authorization", token)
	recCrear := httptest.NewRecorder()
	router.ServeHTTP(recCrear, reqCrear)
	require.Equal(t, http.StatusCreated, recCrear.Code)

	var creada models.Ruta
	require.NoError(t, json.Unmarshal(recCrear.Body.Bytes(), &creada))

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/rutas/"+strconv.FormatUint(uint64(creada.ID), 10), nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code, "body: %s", rec.Body.String())
}

// Test 404 — DELETE /rutas/{id} con ID inexistente → 404 Not Found.
func TestHandler_BorrarRuta_Inexistente_Devuelve404(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/rutas/9999", nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code, "body: %s", rec.Body.String())
}

// ═══════════════════════════════ PUNTOS ═════════════════════════════════════

// Test 201 — POST /puntos con datos válidos → 201 Created.
func TestHandler_CrearPunto_Valido_Devuelve201(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.Punto{RutaID: 1, Nombre: "Muelle", Direccion: "Av. Principal"})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/puntos", bytes.NewReader(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code, "body: %s", rec.Body.String())

	var respuesta models.Punto
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &respuesta))
	assert.NotZero(t, respuesta.ID)
	assert.Equal(t, "activo", respuesta.Estado)
}

// Test 200 — GET /puntos → 200 OK.
func TestHandler_ListarPuntos_Devuelve200(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/puntos", nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "body: %s", rec.Body.String())
}

// Test 404 — GET /puntos/{id} con ID inexistente → 404 Not Found.
func TestHandler_ObtenerPunto_Inexistente_Devuelve404(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/puntos/9999", nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code, "body: %s", rec.Body.String())
}

// Test 200 — PUT /puntos/{id} con datos válidos → 200 OK.
func TestHandler_ActualizarPunto_Valido_Devuelve200(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.Punto{RutaID: 1, Nombre: "Muelle", Direccion: "Av. Principal"})
	reqCrear := httptest.NewRequest(http.MethodPost, "/api/v1/puntos", bytes.NewReader(cuerpo))
	reqCrear.Header.Set("Content-Type", "application/json")
	reqCrear.Header.Set("Authorization", token)
	recCrear := httptest.NewRecorder()
	router.ServeHTTP(recCrear, reqCrear)
	require.Equal(t, http.StatusCreated, recCrear.Code)

	var creado models.Punto
	require.NoError(t, json.Unmarshal(recCrear.Body.Bytes(), &creado))

	actualizacion, _ := json.Marshal(models.Punto{RutaID: 1, Nombre: "Muelle 2", Direccion: "Av. Secundaria"})
	req := httptest.NewRequest(http.MethodPut, "/api/v1/puntos/"+strconv.FormatUint(uint64(creado.ID), 10), bytes.NewReader(actualizacion))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "body: %s", rec.Body.String())
}

// Test 204 — DELETE /puntos/{id} con ID existente → 204 No Content.
func TestHandler_BorrarPunto_Existente_Devuelve204(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.Punto{RutaID: 1, Nombre: "Muelle", Direccion: "Av. Principal"})
	reqCrear := httptest.NewRequest(http.MethodPost, "/api/v1/puntos", bytes.NewReader(cuerpo))
	reqCrear.Header.Set("Content-Type", "application/json")
	reqCrear.Header.Set("Authorization", token)
	recCrear := httptest.NewRecorder()
	router.ServeHTTP(recCrear, reqCrear)
	require.Equal(t, http.StatusCreated, recCrear.Code)

	var creado models.Punto
	require.NoError(t, json.Unmarshal(recCrear.Body.Bytes(), &creado))

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/puntos/"+strconv.FormatUint(uint64(creado.ID), 10), nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code, "body: %s", rec.Body.String())
}

// ═══════════════════════════ TRANSPORTISTAS ═════════════════════════════════

// Test 201 — POST /transportistas con datos válidos → 201 Created.
func TestHandler_CrearTransportista_Valido_Devuelve201(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.Transportista{Nombre: "Carlos", Telefono: "0991234567", PlacaVehiculo: "ABC-123"})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/transportistas", bytes.NewReader(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code, "body: %s", rec.Body.String())

	var respuesta models.Transportista
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &respuesta))
	assert.Equal(t, "activo", respuesta.Estado)
}

// Test 200 — GET /transportistas → 200 OK.
func TestHandler_ListarTransportistas_Devuelve200(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/transportistas", nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "body: %s", rec.Body.String())
}

// Test 404 — GET /transportistas/{id} con ID inexistente → 404 Not Found.
func TestHandler_ObtenerTransportista_Inexistente_Devuelve404(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/transportistas/9999", nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code, "body: %s", rec.Body.String())
}

// Test 200 — PUT /transportistas/{id} con datos válidos → 200 OK.
func TestHandler_ActualizarTransportista_Valido_Devuelve200(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.Transportista{Nombre: "Carlos", Telefono: "0991234567", PlacaVehiculo: "ABC-123"})
	reqCrear := httptest.NewRequest(http.MethodPost, "/api/v1/transportistas", bytes.NewReader(cuerpo))
	reqCrear.Header.Set("Content-Type", "application/json")
	reqCrear.Header.Set("Authorization", token)
	recCrear := httptest.NewRecorder()
	router.ServeHTTP(recCrear, reqCrear)
	require.Equal(t, http.StatusCreated, recCrear.Code)

	var creado models.Transportista
	require.NoError(t, json.Unmarshal(recCrear.Body.Bytes(), &creado))

	actualizacion, _ := json.Marshal(models.Transportista{Nombre: "Carlos P.", Telefono: "0991234567", PlacaVehiculo: "ABC-123"})
	req := httptest.NewRequest(http.MethodPut, "/api/v1/transportistas/"+strconv.FormatUint(uint64(creado.ID), 10), bytes.NewReader(actualizacion))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "body: %s", rec.Body.String())
}

// Test 204 — DELETE /transportistas/{id} con ID existente → 204 No Content.
func TestHandler_BorrarTransportista_Existente_Devuelve204(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.Transportista{Nombre: "Carlos", Telefono: "0991234567", PlacaVehiculo: "ABC-123"})
	reqCrear := httptest.NewRequest(http.MethodPost, "/api/v1/transportistas", bytes.NewReader(cuerpo))
	reqCrear.Header.Set("Content-Type", "application/json")
	reqCrear.Header.Set("Authorization", token)
	recCrear := httptest.NewRecorder()
	router.ServeHTTP(recCrear, reqCrear)
	require.Equal(t, http.StatusCreated, recCrear.Code)

	var creado models.Transportista
	require.NoError(t, json.Unmarshal(recCrear.Body.Bytes(), &creado))

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/transportistas/"+strconv.FormatUint(uint64(creado.ID), 10), nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code, "body: %s", rec.Body.String())
}

// ═══════════════════════════════ ENTREGAS ═══════════════════════════════════

// Test 400 — POST /entregas con PedidoID vacío → 400 Bad Request.
func TestHandler_CrearEntrega_PedidoIDVacio_Devuelve400(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.EntregaPedido{PedidoID: 0, PuntoID: 1, TransportistaID: 1})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/entregas", bytes.NewReader(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code, "body: %s", rec.Body.String())
}

// Test 201 — POST /entregas con datos válidos → 201 Created.
func TestHandler_CrearEntrega_Valida_Devuelve201(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.EntregaPedido{PedidoID: 1, PuntoID: 1, TransportistaID: 1})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/entregas", bytes.NewReader(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code, "body: %s", rec.Body.String())

	var respuesta models.EntregaPedido
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &respuesta))
	assert.Equal(t, "pendiente", respuesta.Estado)
}

// Test 200 — GET /entregas → 200 OK.
func TestHandler_ListarEntregas_Devuelve200(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/entregas", nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "body: %s", rec.Body.String())
}

// Test 404 — GET /entregas/{id} con ID inexistente → 404 Not Found.
func TestHandler_ObtenerEntrega_Inexistente_Devuelve404(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/entregas/9999", nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code, "body: %s", rec.Body.String())
}

// Test 200 — PUT /entregas/{id} con datos válidos → 200 OK.
func TestHandler_ActualizarEntrega_Valida_Devuelve200(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.EntregaPedido{PedidoID: 1, PuntoID: 1, TransportistaID: 1})
	reqCrear := httptest.NewRequest(http.MethodPost, "/api/v1/entregas", bytes.NewReader(cuerpo))
	reqCrear.Header.Set("Content-Type", "application/json")
	reqCrear.Header.Set("Authorization", token)
	recCrear := httptest.NewRecorder()
	router.ServeHTTP(recCrear, reqCrear)
	require.Equal(t, http.StatusCreated, recCrear.Code)

	var creada models.EntregaPedido
	require.NoError(t, json.Unmarshal(recCrear.Body.Bytes(), &creada))

	actualizacion, _ := json.Marshal(models.EntregaPedido{PedidoID: 1, PuntoID: 1, TransportistaID: 1, Estado: "entregado"})
	req := httptest.NewRequest(http.MethodPut, "/api/v1/entregas/"+strconv.FormatUint(uint64(creada.ID), 10), bytes.NewReader(actualizacion))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "body: %s", rec.Body.String())
}

// Test 204 — DELETE /entregas/{id} con ID existente → 204 No Content.
func TestHandler_BorrarEntrega_Existente_Devuelve204(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.EntregaPedido{PedidoID: 1, PuntoID: 1, TransportistaID: 1})
	reqCrear := httptest.NewRequest(http.MethodPost, "/api/v1/entregas", bytes.NewReader(cuerpo))
	reqCrear.Header.Set("Content-Type", "application/json")
	reqCrear.Header.Set("Authorization", token)
	recCrear := httptest.NewRecorder()
	router.ServeHTTP(recCrear, reqCrear)
	require.Equal(t, http.StatusCreated, recCrear.Code)

	var creada models.EntregaPedido
	require.NoError(t, json.Unmarshal(recCrear.Body.Bytes(), &creada))

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/entregas/"+strconv.FormatUint(uint64(creada.ID), 10), nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code, "body: %s", rec.Body.String())
}

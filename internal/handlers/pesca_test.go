package handlers

// Test 2 — Handler con httptest + fake en memoria.
//
// Qué comprueba:
//   a) Ruta protegida SIN token → 401 Unauthorized.
//   b) Token expirado → 401 Unauthorized.
//   c) POST /especies con datos válidos y token → 201 Created.
//   d) POST /especies con nombre_comun vacío → 400 Bad Request.
//
// Diferencia mock vs fake:
//   Mock  → configuras exactamente qué devuelve cada llamada (para aislar).
//   Fake  → implementación real simple en RAM; aquí SÍ se guardan datos.

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

// ─── Fake de AlmacenPesca en memoria ─────────────────────────────────────────
// A diferencia del mock, el fake SÍ guarda datos reales en RAM.
// Permite verificar que lo que se crea luego se puede leer.

type especieFake struct {
	datos  []models.Especie
	nextID int
}

func nuevoEspecieFake() *especieFake { return &especieFake{nextID: 1} }

func (f *especieFake) ListarEspecies() []models.Especie { return f.datos }
func (f *especieFake) BuscarEspeciePorID(id int) (models.Especie, bool) {
	for _, e := range f.datos {
		if e.ID == id {
			return e, true
		}
	}
	return models.Especie{}, false
}
func (f *especieFake) CrearEspecie(e models.Especie) models.Especie {
	e.ID = f.nextID
	f.nextID++
	f.datos = append(f.datos, e)
	return e
}
func (f *especieFake) ActualizarEspecie(id int, d models.Especie) (models.Especie, bool) {
	for i, e := range f.datos {
		if e.ID == id {
			d.ID = id
			f.datos[i] = d
			return d, true
		}
	}
	return models.Especie{}, false
}
func (f *especieFake) BorrarEspecie(id int) bool {
	for i, e := range f.datos {
		if e.ID == id {
			f.datos = append(f.datos[:i], f.datos[i+1:]...)
			return true
		}
	}
	return false
}

// Resto de AlmacenPesca — no usados en estos tests.
func (f *especieFake) ListarPescadores() []models.Pescador { return nil }
func (f *especieFake) BuscarPescadorPorID(id int) (models.Pescador, bool) {
	return models.Pescador{}, false
}
func (f *especieFake) CrearPescador(p models.Pescador) models.Pescador { return p }
func (f *especieFake) ActualizarPescador(id int, d models.Pescador) (models.Pescador, bool) {
	return d, true
}
func (f *especieFake) BorrarPescador(id int) bool                { return true }
func (f *especieFake) ListarEmbarcaciones() []models.Embarcacion { return nil }
func (f *especieFake) BuscarEmbarcacionPorID(id int) (models.Embarcacion, bool) {
	return models.Embarcacion{}, false
}
func (f *especieFake) CrearEmbarcacion(e models.Embarcacion) models.Embarcacion { return e }
func (f *especieFake) ActualizarEmbarcacion(id int, d models.Embarcacion) (models.Embarcacion, bool) {
	return d, true
}
func (f *especieFake) BorrarEmbarcacion(id int) bool    { return true }
func (f *especieFake) ListarCapturas() []models.Captura { return nil }
func (f *especieFake) BuscarCapturaPorID(id int) (models.Captura, bool) {
	return models.Captura{}, false
}
func (f *especieFake) CrearCaptura(c models.Captura) models.Captura { return c }
func (f *especieFake) ActualizarCaptura(id int, d models.Captura) (models.Captura, bool) {
	return d, true
}
func (f *especieFake) BorrarCaptura(id int) bool                                      { return true }
func (f *especieFake) ListarBodegas() []models.Bodega                                 { return nil }
func (f *especieFake) BuscarBodegaPorID(id int) (models.Bodega, bool)                 { return models.Bodega{}, false }
func (f *especieFake) CrearBodega(b models.Bodega) models.Bodega                      { return b }
func (f *especieFake) ActualizarBodega(id int, d models.Bodega) (models.Bodega, bool) { return d, true }
func (f *especieFake) BorrarBodega(id int) bool                                       { return true }
func (f *especieFake) ListarStocks() []models.Stock                                   { return nil }
func (f *especieFake) BuscarStockPorID(id int) (models.Stock, bool)                   { return models.Stock{}, false }
func (f *especieFake) CrearStock(s models.Stock) models.Stock                         { return s }
func (f *especieFake) ActualizarStock(id int, d models.Stock) (models.Stock, bool)    { return d, true }
func (f *especieFake) BorrarStock(id int) bool                                        { return true }

// ─── Fake de UserRepository ───────────────────────────────────────────────────

type usuarioFake struct {
	datos  []models.Usuario
	nextID int
}

func nuevoUsuarioFake() *usuarioFake { return &usuarioFake{nextID: 1} }

func (u *usuarioFake) CrearUsuario(usr models.Usuario) (models.Usuario, error) {
	usr.ID = u.nextID
	u.nextID++
	u.datos = append(u.datos, usr)
	return usr, nil
}
func (u *usuarioFake) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	for _, usr := range u.datos {
		if usr.Email == email {
			return usr, true
		}
	}
	return models.Usuario{}, false
}

// ─── Helper: genera un token JWT válido para los tests ───────────────────────

func generarTokenTest(t *testing.T, auth *service.AuthService) string {
	t.Helper()
	u := models.Usuario{ID: 1, Email: "test@pesca.com", PasswordHash: "x"}
	token, err := auth.GenerarToken(u)
	require.NoError(t, err, "no se pudo generar token de prueba")
	return "Bearer " + token
}

// ─── Helper: construye el router completo en memoria ─────────────────────────

func nuevoRouterTest(t *testing.T) (http.Handler, *service.AuthService) {
	t.Helper()

	pescaService := service.NewPescaService(nuevoEspecieFake())
	authService := service.NewAuthService(nuevoUsuarioFake())
	pedidoService := service.NewPedidoService(storage.NewMemoria())
	rutasService := service.NewRutasService(storage.NuevaMemoriaRutas())

	servidor := NewServer(pescaService, pedidoService, rutasService, authService)

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))
			r.Get("/especies", servidor.ListarEspecies)
			r.Post("/especies", servidor.CrearEspecie)
			r.Get("/especies/{id}", servidor.ObtenerEspecie)
		})
	})

	return r, authService
}

// ─── Tests ────────────────────────────────────────────────────────────────────

// Test 401 — GET /especies SIN token debe devolver 401.
// Esto demuestra que el middleware.Auth protege la ruta correctamente.
func TestHandler_GetEspecies_SinToken_Devuelve401(t *testing.T) {
	// ── Preparar ──────────────────────────────────────────────────────────
	router, _ := nuevoRouterTest(t)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/especies", nil)
	// ← sin header Authorization
	rec := httptest.NewRecorder()

	// ── Ejecutar ──────────────────────────────────────────────────────────
	router.ServeHTTP(rec, req)

	// ── Verificar ─────────────────────────────────────────────────────────
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

// Test 401 — Token expirado debe devolver 401.
func TestHandler_TokenExpirado_Devuelve401(t *testing.T) {
	// ── Preparar ──────────────────────────────────────────────────────────
	router, _ := nuevoRouterTest(t)

	// Usamos la misma clave secreta del servicio para firmar el token expirado.
	// El test crea manualmente un token con ExpiresAt en el pasado.
	secretJWTTest := []byte("pesca-directa-tarqui-secret-2026")
	claims := &service.Claims{
		UsuarioID: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)), // ya expiró
		},
	}
	tokenExpirado, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretJWTTest)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/especies", nil)
	req.Header.Set("Authorization", "Bearer "+tokenExpirado)
	rec := httptest.NewRecorder()

	// ── Ejecutar ──────────────────────────────────────────────────────────
	router.ServeHTTP(rec, req)

	// ── Verificar ─────────────────────────────────────────────────────────
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

// Test 201 — POST /especies con datos válidos y token → 201 Created.
func TestHandler_CrearEspecie_Valida_Devuelve201(t *testing.T) {
	// ── Preparar ──────────────────────────────────────────────────────────
	router, auth := nuevoRouterTest(t)
	token := generarTokenTest(t, auth)

	cuerpo, _ := json.Marshal(models.Especie{
		NombreComun:  "Corvina",
		UnidadMedida: "kg",
		Temporada:    "enero a marzo",
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/especies", bytes.NewReader(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	// ── Ejecutar ──────────────────────────────────────────────────────────
	router.ServeHTTP(rec, req)

	// ── Verificar ─────────────────────────────────────────────────────────
	require.Equal(t, http.StatusCreated, rec.Code, "body: %s", rec.Body.String())

	var respuesta models.Especie
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &respuesta))
	assert.NotZero(t, respuesta.ID, "la especie creada debe tener un ID asignado")
	assert.Equal(t, "Corvina", respuesta.NombreComun)
}

// Test 400 — POST /especies con nombre_comun vacío → 400 Bad Request.
func TestHandler_CrearEspecie_NombreVacio_Devuelve400(t *testing.T) {
	// ── Preparar ──────────────────────────────────────────────────────────
	router, auth := nuevoRouterTest(t)
	token := generarTokenTest(t, auth)

	cuerpo, _ := json.Marshal(models.Especie{
		NombreComun:  "", // ← campo requerido vacío
		UnidadMedida: "kg",
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/especies", bytes.NewReader(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	// ── Ejecutar ──────────────────────────────────────────────────────────
	router.ServeHTTP(rec, req)

	// ── Verificar ─────────────────────────────────────────────────────────
	assert.Equal(t, http.StatusBadRequest, rec.Code, "body: %s", rec.Body.String())
}

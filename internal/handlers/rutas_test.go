package handlers

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

func generarTokenTestRutas(t *testing.T, auth *service.AuthService) string {
	t.Helper()
	u := models.Usuario{ID: 1, Email: "test@rutas.com", PasswordHash: "x"}
	token, err := auth.GenerarToken(u)
	require.NoError(t, err, "no se pudo generar token de prueba")
	return "Bearer " + token
}

// ─── Helper: construye el router completo en memoria ─────────────────────────

func nuevoRouterTestRutas(t *testing.T) (http.Handler, *service.AuthService) {
	t.Helper()

	authService := service.NewAuthService(nuevoUsuarioFakeRutas())
	rutasService := service.NewRutasService(storage.NuevaMemoriaRutas())

	// NewServer necesita los 4 services — los otros dos con nil no compilan,
	// así que pasamos fakes mínimos usando los storages en memoria del proyecto.
	pescaService := service.NewPescaService(storage.NuevaMemoriaPesca())
	pedidoService := service.NewPedidoService(storage.NewMemoria())

	servidor := NewServer(pescaService, pedidoService, rutasService, authService)

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))
			r.Get("/rutas", servidor.ListarRutas)
			r.Post("/rutas", servidor.CrearRuta)
			r.Get("/rutas/{id}", servidor.ObtenerRuta)
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
	claims := &service.Claims{
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

// Test 201 — POST /rutas con datos válidos y token → 201 Created.
func TestHandler_CrearRuta_Valida_Devuelve201(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.Ruta{
		Nombre:  "Ruta Norte",
		Origen:  "Puerto Tarqui",
		Destino: "Mercado Central",
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/rutas", bytes.NewReader(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code, "body: %s", rec.Body.String())

	var respuesta models.Ruta
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &respuesta))
	assert.NotZero(t, respuesta.ID, "la ruta creada debe tener un ID asignado")
	assert.Equal(t, "Ruta Norte", respuesta.Nombre)
	assert.Equal(t, "activo", respuesta.Estado)
}

// Test 400 — POST /rutas con nombre vacío → 400 Bad Request.
func TestHandler_CrearRuta_NombreVacio_Devuelve400(t *testing.T) {
	router, auth := nuevoRouterTestRutas(t)
	token := generarTokenTestRutas(t, auth)

	cuerpo, _ := json.Marshal(models.Ruta{
		Nombre:  "",
		Origen:  "Puerto Tarqui",
		Destino: "Mercado Central",
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/rutas", bytes.NewReader(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code, "body: %s", rec.Body.String())
}

package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

// ─── Fake de UserRepository ───────────────────────────────────────────────────

type usuarioRepoFakePedidos struct {
	porEmail map[string]models.Usuario
	nextID   int
}

func nuevoUsuarioRepoFakePedidos() *usuarioRepoFakePedidos {
	return &usuarioRepoFakePedidos{porEmail: map[string]models.Usuario{}, nextID: 1}
}

func (f *usuarioRepoFakePedidos) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	u.ID = f.nextID
	f.nextID++
	f.porEmail[u.Email] = u
	return u, nil
}

func (f *usuarioRepoFakePedidos) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	u, ok := f.porEmail[email]
	return u, ok
}

// ─── Helper: construye el router completo en memoria ─────────────────────────

// construirEntornoPedidos arma el mismo router que main.go pero con almacén
// en memoria y repo de usuarios fake. Devuelve el handler y un token válido.
func construirEntornoPedidos(t *testing.T) (http.Handler, string) {
	t.Helper()

	almacen := storage.NewMemoria()
	almacen.Seed()
	usuarios := nuevoUsuarioRepoFakePedidos()

	pedidoSvc := service.NewPedidoService(almacen)
	authSvc := service.NewAuthService(usuarios)
	pescaSvc := service.NewPescaService(storage.NuevaMemoriaPesca())
	rutasSvc := service.NewRutasService(storage.NuevaMemoriaRutas())

	srv := handlers.NewServer(pescaSvc, pedidoSvc, rutasSvc, authSvc)

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", srv.Registrar)
		r.Post("/auth/login", srv.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))
			r.Get("/clientes", srv.ListarClientes)
			r.Post("/clientes", srv.CrearCliente)
			r.Get("/clientes/{id}", srv.ObtenerCliente)
		})
	})

	token := registrarYObtenerTokenPedidos(t, r)
	return r, token
}

// registrarYObtenerTokenPedidos hace register + login contra el propio router
// para conseguir un JWT válido, igual que lo haría un cliente real.
func registrarYObtenerTokenPedidos(t *testing.T, h http.Handler) string {
	t.Helper()
	cred := `{"email":"michelle@pedidos.com","password":"pedidos123"}`

	reqReg := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", strings.NewReader(cred))
	h.ServeHTTP(httptest.NewRecorder(), reqReg)

	reqLogin := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(cred))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, reqLogin)
	require.Equal(t, http.StatusOK, rec.Code, "el login debería devolver 200")

	var resp struct {
		Token string `json:"token"`
	}
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	require.NotEmpty(t, resp.Token)
	return resp.Token
}

// ─── Tests ────────────────────────────────────────────────────────────────────

// TestCrearCliente_Exitoso: POST con token y cuerpo válido → 201 Created.
func TestCrearCliente_Exitoso(t *testing.T) {
	h, token := construirEntornoPedidos(t)
	body := `{"usuario_id":1,"tipo_cliente":"restaurante","nombre_negocio":"Sushi Koi","direccion":"Av. Flavio Reyes","telefono":"0991234567"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/clientes", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)
	var creado models.Cliente
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))
	assert.NotZero(t, creado.ID)
	assert.Equal(t, "Sushi Koi", creado.NombreNegocio)
	assert.Equal(t, "activo", creado.Estado)
}

// TestObtenerCliente_NoEncontrado: id inexistente → 404 Not Found.
func TestObtenerCliente_NoEncontrado(t *testing.T) {
	h, token := construirEntornoPedidos(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/clientes/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// TestCrearCliente_Invalido: cuerpo que viola la regla de negocio → 400.
func TestCrearCliente_Invalido(t *testing.T) {
	h, token := construirEntornoPedidos(t)
	body := `{"usuario_id":1,"tipo_cliente":"restaurante","nombre_negocio":"","direccion":"Av. 1","telefono":"0991234567"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/clientes", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// TestRutaProtegida_SinToken: sin header Authorization → 401 Unauthorized.
func TestRutaProtegida_SinToken(t *testing.T) {
	h, _ := construirEntornoPedidos(t)
	body := `{"usuario_id":1,"tipo_cliente":"restaurante","nombre_negocio":"Test","direccion":"Av. 1","telefono":"0991234567"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/clientes", strings.NewReader(body))
	// A propósito: NO seteamos Authorization.
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

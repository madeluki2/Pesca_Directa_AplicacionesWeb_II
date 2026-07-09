package gestion_pedidos_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	gp "Pesca_Directa_AplicacionesWeb_II/internal/handlers/gestion_pedidos"
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	svcgp "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pedidos"
	storagegp "Pesca_Directa_AplicacionesWeb_II/internal/storage/gestion_pedidos"
)

// ─── Helper: abre SQLite en memoria ──────────────────────────────────────────

func abrirDBEnMemoria(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	sqlDB, err := db.DB()
	require.NoError(t, err)
	sqlDB.SetMaxOpenConns(1)

	require.NoError(t, db.AutoMigrate(
		&models.Cliente{},
		&models.Pedido{},
		&models.DetallePedido{},
	))
	return db
}

// ─── Fake de UserRepository ───────────────────────────────────────────────────

type usuarioFakePedido struct {
	porEmail map[string]models.Usuario
	nextID   int
}

func nuevoUsuarioFakePedido() *usuarioFakePedido {
	return &usuarioFakePedido{porEmail: map[string]models.Usuario{}, nextID: 1}
}

func (f *usuarioFakePedido) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	u.ID = f.nextID
	f.nextID++
	f.porEmail[u.Email] = u
	return u, nil
}

func (f *usuarioFakePedido) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	u, ok := f.porEmail[email]
	return u, ok
}

// ─── Helper: construye el router ─────────────────────────────────────────────

func construirEntornoPedido(t *testing.T) (http.Handler, string) {
	t.Helper()

	db := abrirDBEnMemoria(t)
	almacen := storagegp.NuevoAlmacenSQLite(db)
	almacen.Seed()

	usuarios := nuevoUsuarioFakePedido()
	pedidoSvc := svcgp.NewPedidoService(almacen)
	authSvc := service.NewAuthService(usuarios)

	srv := &gp.Server{
		Pedidos: pedidoSvc,
		Auth:    authSvc,
	}

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", srv.Registrar)
		r.Post("/auth/login", srv.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))
			r.Post("/pedidos", srv.CrearPedido)
			r.Get("/pedidos/{id}", srv.ObtenerPedido)
			r.Get("/pedidos", srv.ListarPedidos)
			r.Put("/pedidos/{id}", srv.ActualizarPedido)  // añadido
			r.Delete("/pedidos/{id}", srv.EliminarPedido) // añadido
		})
	})

	token := obtenerTokenPedido(t, r)
	return r, token
}

func obtenerTokenPedido(t *testing.T, h http.Handler) string {
	t.Helper()
	cred := `{"email":"ilaria@pedidos.com","password":"pedidos123"}`

	reqReg := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", strings.NewReader(cred))
	h.ServeHTTP(httptest.NewRecorder(), reqReg)

	reqLogin := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(cred))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, reqLogin)
	require.Equal(t, http.StatusOK, rec.Code)

	var resp struct {
		Token string `json:"token"`
	}
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	require.NotEmpty(t, resp.Token)
	return resp.Token
}

// ─── Tests de Pedido ─────────────────────────────────────────────────────────

func TestCrearPedido_Exitoso(t *testing.T) {
	h, token := construirEntornoPedido(t)

	body := `{"cliente_id":1,"fecha":"` + time.Now().Format(time.RFC3339) + `"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/pedidos", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code, "body: %s", rec.Body.String())
	var creado models.Pedido
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))
	assert.NotZero(t, creado.ID)
	assert.Equal(t, "pendiente", creado.Estado)
}

func TestCrearPedido_SinClienteID(t *testing.T) {
	h, token := construirEntornoPedido(t)

	body := `{"cliente_id":0,"fecha":"` + time.Now().Format(time.RFC3339) + `"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/pedidos", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code, "body: %s", rec.Body.String())
}

func TestObtenerPedido_NoEncontrado(t *testing.T) {
	h, token := construirEntornoPedido(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/pedidos/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestPedido_SinToken(t *testing.T) {
	h, _ := construirEntornoPedido(t)

	body := `{"cliente_id":1,"fecha":"` + time.Now().Format(time.RFC3339) + `"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/pedidos", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestActualizarPedido_Exitoso(t *testing.T) {
	h, token := construirEntornoPedido(t)

	// Crear pedido
	body := `{"cliente_id":1,"fecha":"` + time.Now().Format(time.RFC3339) + `"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/pedidos", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	require.Equal(t, http.StatusCreated, rec.Code)

	var creado models.Pedido
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))

	// Actualizar pedido
	updateBody := `{"estado":"completado"}`
	reqUpd := httptest.NewRequest(http.MethodPut, "/api/v1/pedidos/"+strconv.Itoa(creado.ID), strings.NewReader(updateBody))
	reqUpd.Header.Set("Authorization", "Bearer "+token)
	reqUpd.Header.Set("Content-Type", "application/json")
	recUpd := httptest.NewRecorder()
	h.ServeHTTP(recUpd, reqUpd)

	require.Equal(t, http.StatusOK, recUpd.Code)
	var actualizado models.Pedido
	require.NoError(t, json.NewDecoder(recUpd.Body).Decode(&actualizado))
	assert.Equal(t, "completado", actualizado.Estado)
}

func TestEliminarPedido_Exitoso(t *testing.T) {
	h, token := construirEntornoPedido(t)

	// Crear pedido
	body := `{"cliente_id":1,"fecha":"` + time.Now().Format(time.RFC3339) + `"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/pedidos", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	require.Equal(t, http.StatusCreated, rec.Code)

	var creado models.Pedido
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))

	// Eliminar pedido
	reqDel := httptest.NewRequest(http.MethodDelete, "/api/v1/pedidos/"+strconv.Itoa(creado.ID), nil)
	reqDel.Header.Set("Authorization", "Bearer "+token)
	recDel := httptest.NewRecorder()
	h.ServeHTTP(recDel, reqDel)

	assert.Equal(t, http.StatusNoContent, recDel.Code)
}

func TestListarPedidos(t *testing.T) {
	h, token := construirEntornoPedido(t)

	// Crear dos pedidos
	for i := 0; i < 2; i++ {
		body := `{"cliente_id":1,"fecha":"` + time.Now().Format(time.RFC3339) + `"}`
		req := httptest.NewRequest(http.MethodPost, "/api/v1/pedidos", strings.NewReader(body))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusCreated, rec.Code)
	}

	// Listar pedidos
	reqList := httptest.NewRequest(http.MethodGet, "/api/v1/pedidos", nil)
	reqList.Header.Set("Authorization", "Bearer "+token)
	recList := httptest.NewRecorder()
	h.ServeHTTP(recList, reqList)

	require.Equal(t, http.StatusOK, recList.Code)
	var pedidos []models.Pedido
	require.NoError(t, json.NewDecoder(recList.Body).Decode(&pedidos))
	assert.GreaterOrEqual(t, len(pedidos), 2)
}

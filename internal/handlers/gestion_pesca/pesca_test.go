package gestion_pesca

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	authService "Pesca_Directa_AplicacionesWeb_II/internal/service"
	service "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pesca"
)

type fakeAlmacenParaHandlers struct {
	datos  []models.Especie
	nextID int
}

func nuevoFakeAlmacen() *fakeAlmacenParaHandlers { return &fakeAlmacenParaHandlers{nextID: 1} }

func (f *fakeAlmacenParaHandlers) ListarEspecies() []models.Especie { return f.datos }
func (f *fakeAlmacenParaHandlers) BuscarEspeciePorID(id int) (models.Especie, bool) {
	for _, e := range f.datos {
		if e.ID == id {
			return e, true
		}
	}
	return models.Especie{}, false
}
func (f *fakeAlmacenParaHandlers) CrearEspecie(e models.Especie) models.Especie {
	e.ID = f.nextID
	f.nextID++
	f.datos = append(f.datos, e)
	return e
}
func (f *fakeAlmacenParaHandlers) ActualizarEspecie(id int, e models.Especie) (models.Especie, bool) {
	for i, v := range f.datos {
		if v.ID == id {
			e.ID = id
			f.datos[i] = e
			return e, true
		}
	}
	return models.Especie{}, false
}
func (f *fakeAlmacenParaHandlers) BorrarEspecie(id int) bool {
	for i, v := range f.datos {
		if v.ID == id {
			f.datos = append(f.datos[:i], f.datos[i+1:]...)
			return true
		}
	}
	return false
}

// Stubs requeridos por compatibilidad estructural con la interfaz del proyecto
func (f *fakeAlmacenParaHandlers) ListarPescadores() []models.Pescador { return nil }
func (f *fakeAlmacenParaHandlers) BuscarPescadorPorID(id int) (models.Pescador, bool) {
	return models.Pescador{}, false
}
func (f *fakeAlmacenParaHandlers) CrearPescador(p models.Pescador) models.Pescador { return p }
func (f *fakeAlmacenParaHandlers) ActualizarPescador(id int, d models.Pescador) (models.Pescador, bool) {
	return d, false
}
func (f *fakeAlmacenParaHandlers) BorrarPescador(id int) bool                { return false }
func (f *fakeAlmacenParaHandlers) ListarEmbarcaciones() []models.Embarcacion { return nil }
func (f *fakeAlmacenParaHandlers) BuscarEmbarcacionPorID(id int) (models.Embarcacion, bool) {
	return models.Embarcacion{}, false
}
func (f *fakeAlmacenParaHandlers) CrearEmbarcacion(e models.Embarcacion) models.Embarcacion { return e }
func (f *fakeAlmacenParaHandlers) ActualizarEmbarcacion(id int, d models.Embarcacion) (models.Embarcacion, bool) {
	return d, false
}
func (f *fakeAlmacenParaHandlers) BorrarEmbarcacion(id int) bool    { return false }
func (f *fakeAlmacenParaHandlers) ListarCapturas() []models.Captura { return nil }
func (f *fakeAlmacenParaHandlers) BuscarCapturaPorID(id int) (models.Captura, bool) {
	return models.Captura{}, false
}
func (f *fakeAlmacenParaHandlers) CrearCaptura(c models.Captura) models.Captura { return c }
func (f *fakeAlmacenParaHandlers) ActualizarCaptura(id int, d models.Captura) (models.Captura, bool) {
	return d, false
}
func (f *fakeAlmacenParaHandlers) BorrarCaptura(id int) bool      { return false }
func (f *fakeAlmacenParaHandlers) ListarBodegas() []models.Bodega { return nil }
func (f *fakeAlmacenParaHandlers) BuscarBodegaPorID(id int) (models.Bodega, bool) {
	return models.Bodega{}, false
}
func (f *fakeAlmacenParaHandlers) CrearBodega(b models.Bodega) models.Bodega { return b }
func (f *fakeAlmacenParaHandlers) ActualizarBodega(id int, d models.Bodega) (models.Bodega, bool) {
	return d, false
}
func (f *fakeAlmacenParaHandlers) BorrarBodega(id int) bool     { return false }
func (f *fakeAlmacenParaHandlers) ListarStocks() []models.Stock { return nil }
func (f *fakeAlmacenParaHandlers) BuscarStockPorID(id int) (models.Stock, bool) {
	return models.Stock{}, false
}
func (f *fakeAlmacenParaHandlers) CrearStock(s models.Stock) models.Stock { return s }
func (f *fakeAlmacenParaHandlers) ActualizarStock(id int, d models.Stock) (models.Stock, bool) {
	return d, false
}
func (f *fakeAlmacenParaHandlers) BorrarStock(id int) bool { return false }

type fakeUsuariosAuth struct{}

func (fakeUsuariosAuth) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	return u, nil
}

func (fakeUsuariosAuth) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	return models.Usuario{}, false
}

func nuevoRouterTest(t *testing.T) (http.Handler, *authService.AuthService) {
	t.Helper()
	auth := authService.NewAuthService(fakeUsuariosAuth{})
	r := chi.NewRouter()

	fakeStore := nuevoFakeAlmacen()
	svc := service.NewPescaService(fakeStore)
	server := NewServer(Deps{Pesca: svc})

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.Auth(auth))
		r.Post("/especies", server.CrearEspecie)
		r.Get("/especies", server.ListarEspecies)
	})

	return r, auth
}

func generarTokenTest(t *testing.T, auth *authService.AuthService) string {
	t.Helper()
	token, err := auth.GenerarToken(models.Usuario{ID: 1, Email: "usuario_test"})
	require.NoError(t, err)
	return "Bearer " + token
}

// ═══════════════════════════ TESTS HANDLERS (ESPECIE) ═══════════════════════════

func TestHandler_CrearEspecie_SinToken_Devuelve401(t *testing.T) {
	router, _ := nuevoRouterTest(t)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/especies", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestHandler_CrearEspecie_Valida_Devuelve201(t *testing.T) {
	router, auth := nuevoRouterTest(t)
	token := generarTokenTest(t, auth)

	cuerpo, _ := json.Marshal(models.Especie{
		NombreComun:  "Corvina",
		UnidadMedida: "kg",
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/especies", bytes.NewReader(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)
	var respuesta models.Especie
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &respuesta))
	assert.Equal(t, "Corvina", respuesta.NombreComun)
}

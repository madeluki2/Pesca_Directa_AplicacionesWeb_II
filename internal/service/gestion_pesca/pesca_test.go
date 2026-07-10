package gestion_pesca_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	service "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pesca"
)

type mockAlmacenPesca struct {
	especies []models.Especie
	nextID   int
}

func nuevoMock() *mockAlmacenPesca { return &mockAlmacenPesca{nextID: 1} }

// ─── Implementación real de tu entidad individual ───
func (m *mockAlmacenPesca) ListarEspecies() []models.Especie { return m.especies }
func (m *mockAlmacenPesca) BuscarEspeciePorID(id int) (models.Especie, bool) {
	for _, e := range m.especies {
		if e.ID == id {
			return e, true
		}
	}
	return models.Especie{}, false
}

func (m *mockAlmacenPesca) CrearEspecie(e models.Especie) models.Especie {
	e.ID = m.nextID
	m.nextID++
	m.especies = append(m.especies, e)
	return e
}
func (m *mockAlmacenPesca) ActualizarEspecie(id int, datos models.Especie) (models.Especie, bool) {
	for i, e := range m.especies {
		if e.ID == id {
			datos.ID = id
			m.especies[i] = datos
			return datos, true
		}
	}
	return models.Especie{}, false
}
func (m *mockAlmacenPesca) BorrarEspecie(id int) bool {
	for i, e := range m.especies {
		if e.ID == id {
			m.especies = append(m.especies[:i], m.especies[i+1:]...)
			return true
		}
	}
	return false
}

// ─── Métodos Stub Dummy requeridos para satisfacer la interfaz AlmacenPesca compartida ───
func (m *mockAlmacenPesca) ListarPescadores() []models.Pescador { return nil }
func (m *mockAlmacenPesca) BuscarPescadorPorID(id int) (models.Pescador, bool) {
	return models.Pescador{}, false
}
func (m *mockAlmacenPesca) CrearPescador(p models.Pescador) models.Pescador { return p }
func (m *mockAlmacenPesca) ActualizarPescador(id int, d models.Pescador) (models.Pescador, bool) {
	return d, false
}
func (m *mockAlmacenPesca) BorrarPescador(id int) bool                { return false }
func (m *mockAlmacenPesca) ListarEmbarcaciones() []models.Embarcacion { return nil }
func (m *mockAlmacenPesca) BuscarEmbarcacionPorID(id int) (models.Embarcacion, bool) {
	return models.Embarcacion{}, false
}
func (m *mockAlmacenPesca) CrearEmbarcacion(e models.Embarcacion) models.Embarcacion { return e }
func (m *mockAlmacenPesca) ActualizarEmbarcacion(id int, d models.Embarcacion) (models.Embarcacion, bool) {
	return d, false
}
func (m *mockAlmacenPesca) BorrarEmbarcacion(id int) bool    { return false }
func (m *mockAlmacenPesca) ListarCapturas() []models.Captura { return nil }
func (m *mockAlmacenPesca) BuscarCapturaPorID(id int) (models.Captura, bool) {
	return models.Captura{}, false
}
func (m *mockAlmacenPesca) CrearCaptura(c models.Captura) models.Captura { return c }
func (m *mockAlmacenPesca) ActualizarCaptura(id int, d models.Captura) (models.Captura, bool) {
	return d, false
}
func (m *mockAlmacenPesca) BorrarCaptura(id int) bool      { return false }
func (m *mockAlmacenPesca) ListarBodegas() []models.Bodega { return nil }
func (m *mockAlmacenPesca) BuscarBodegaPorID(id int) (models.Bodega, bool) {
	return models.Bodega{}, false
}
func (m *mockAlmacenPesca) CrearBodega(b models.Bodega) models.Bodega { return b }
func (m *mockAlmacenPesca) ActualizarBodega(id int, d models.Bodega) (models.Bodega, bool) {
	return d, false
}
func (m *mockAlmacenPesca) BorrarBodega(id int) bool     { return false }
func (m *mockAlmacenPesca) ListarStocks() []models.Stock { return nil }
func (m *mockAlmacenPesca) BuscarStockPorID(id int) (models.Stock, bool) {
	return models.Stock{}, false
}
func (m *mockAlmacenPesca) CrearStock(s models.Stock) models.Stock { return s }
func (m *mockAlmacenPesca) ActualizarStock(id int, d models.Stock) (models.Stock, bool) {
	return d, false
}
func (m *mockAlmacenPesca) BorrarStock(id int) bool { return false }

// ═══════════════════════════ TESTS SERVICE (ESPECIE) ═══════════════════════════

func TestPescaService_CrearEspecie(t *testing.T) {
	mock := nuevoMock()
	svc := service.NewPescaService(mock)

	t.Run("nombre común vacío → error de validación", func(t *testing.T) {
		_, err := svc.CrearEspecie(models.Especie{NombreComun: "   ", UnidadMedida: "kg"})
		require.ErrorIs(t, err, service.ErrNombreComunVacio)
	})

	t.Run("datos válidos → registro exitoso", func(t *testing.T) {
		res, err := svc.CrearEspecie(models.Especie{NombreComun: "Atún", UnidadMedida: "kg"})
		require.NoError(t, err)
		assert.Equal(t, "Atún", res.NombreComun)
	})
}

func TestPescaService_ObtenerEspecie(t *testing.T) {
	mock := nuevoMock()
	svc := service.NewPescaService(mock)
	creada := mock.CrearEspecie(models.Especie{NombreComun: "Picudo", UnidadMedida: "kg"})

	t.Run("ID existente", func(t *testing.T) {
		e, err := svc.ObtenerEspecie(creada.ID)
		require.NoError(t, err)
		assert.Equal(t, "Picudo", e.NombreComun)
	})

	t.Run("ID no existente", func(t *testing.T) {
		_, err := svc.ObtenerEspecie(9999)
		require.ErrorIs(t, err, service.ErrEspecieNoEncontrada)
	})
}

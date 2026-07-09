package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// ─── Mock de AlmacenPesca ─────────────────────────────────────────────────────

type almacenMock struct {
	embarcaciones []models.Embarcacion
	especies      []models.Especie
	pescadores    []models.Pescador
	capturas      []models.Captura
	bodegas       []models.Bodega
	stocks        []models.Stock
	nextID        int
	llamoCrear    bool
}

func nuevoMock() *almacenMock { return &almacenMock{nextID: 1} }

// Especies
func (m *almacenMock) ListarEspecies() []models.Especie { return m.especies }
func (m *almacenMock) BuscarEspeciePorID(id int) (models.Especie, bool) {
	for _, e := range m.especies {
		if e.ID == id {
			return e, true
		}
	}
	return models.Especie{}, false
}
func (m *almacenMock) CrearEspecie(e models.Especie) models.Especie {
	m.llamoCrear = true
	e.ID = m.nextID
	m.nextID++
	m.especies = append(m.especies, e)
	return e
}
func (m *almacenMock) ActualizarEspecie(id int, d models.Especie) (models.Especie, bool) {
	for i, e := range m.especies {
		if e.ID == id {
			d.ID = id
			m.especies[i] = d
			return d, true
		}
	}
	return models.Especie{}, false
}
func (m *almacenMock) BorrarEspecie(id int) bool {
	for i, e := range m.especies {
		if e.ID == id {
			m.especies = append(m.especies[:i], m.especies[i+1:]...)
			return true
		}
	}
	return false
}

// Pescadores
func (m *almacenMock) ListarPescadores() []models.Pescador { return m.pescadores }
func (m *almacenMock) BuscarPescadorPorID(id int) (models.Pescador, bool) {
	for _, p := range m.pescadores {
		if p.ID == id {
			return p, true
		}
	}
	return models.Pescador{}, false
}
func (m *almacenMock) CrearPescador(p models.Pescador) models.Pescador {
	p.ID = m.nextID
	m.nextID++
	m.pescadores = append(m.pescadores, p)
	return p
}
func (m *almacenMock) ActualizarPescador(id int, d models.Pescador) (models.Pescador, bool) {
	for i, p := range m.pescadores {
		if p.ID == id {
			d.ID = id
			m.pescadores[i] = d
			return d, true
		}
	}
	return models.Pescador{}, false
}
func (m *almacenMock) BorrarPescador(id int) bool {
	for i, p := range m.pescadores {
		if p.ID == id {
			m.pescadores = append(m.pescadores[:i], m.pescadores[i+1:]...)
			return true
		}
	}
	return false
}

// Embarcaciones
func (m *almacenMock) ListarEmbarcaciones() []models.Embarcacion { return m.embarcaciones }
func (m *almacenMock) BuscarEmbarcacionPorID(id int) (models.Embarcacion, bool) {
	for _, e := range m.embarcaciones {
		if e.ID == id {
			return e, true
		}
	}
	return models.Embarcacion{}, false
}
func (m *almacenMock) CrearEmbarcacion(e models.Embarcacion) models.Embarcacion {
	e.ID = m.nextID
	m.nextID++
	m.embarcaciones = append(m.embarcaciones, e)
	return e
}
func (m *almacenMock) ActualizarEmbarcacion(id int, d models.Embarcacion) (models.Embarcacion, bool) {
	for i, e := range m.embarcaciones {
		if e.ID == id {
			d.ID = id
			m.embarcaciones[i] = d
			return d, true
		}
	}
	return models.Embarcacion{}, false
}
func (m *almacenMock) BorrarEmbarcacion(id int) bool {
	for i, e := range m.embarcaciones {
		if e.ID == id {
			m.embarcaciones = append(m.embarcaciones[:i], m.embarcaciones[i+1:]...)
			return true
		}
	}
	return false
}

// Capturas
func (m *almacenMock) ListarCapturas() []models.Captura { return m.capturas }
func (m *almacenMock) BuscarCapturaPorID(id int) (models.Captura, bool) {
	for _, c := range m.capturas {
		if c.ID == id {
			return c, true
		}
	}
	return models.Captura{}, false
}
func (m *almacenMock) CrearCaptura(c models.Captura) models.Captura {
	c.ID = m.nextID
	m.nextID++
	m.capturas = append(m.capturas, c)
	return c
}
func (m *almacenMock) ActualizarCaptura(id int, d models.Captura) (models.Captura, bool) {
	for i, c := range m.capturas {
		if c.ID == id {
			d.ID = id
			m.capturas[i] = d
			return d, true
		}
	}
	return models.Captura{}, false
}
func (m *almacenMock) BorrarCaptura(id int) bool {
	for i, c := range m.capturas {
		if c.ID == id {
			m.capturas = append(m.capturas[:i], m.capturas[i+1:]...)
			return true
		}
	}
	return false
}

// Bodegas
func (m *almacenMock) ListarBodegas() []models.Bodega { return m.bodegas }
func (m *almacenMock) BuscarBodegaPorID(id int) (models.Bodega, bool) {
	for _, b := range m.bodegas {
		if b.ID == id {
			return b, true
		}
	}
	return models.Bodega{}, false
}
func (m *almacenMock) CrearBodega(b models.Bodega) models.Bodega {
	b.ID = m.nextID
	m.nextID++
	m.bodegas = append(m.bodegas, b)
	return b
}
func (m *almacenMock) ActualizarBodega(id int, d models.Bodega) (models.Bodega, bool) {
	for i, b := range m.bodegas {
		if b.ID == id {
			d.ID = id
			m.bodegas[i] = d
			return d, true
		}
	}
	return models.Bodega{}, false
}
func (m *almacenMock) BorrarBodega(id int) bool {
	for i, b := range m.bodegas {
		if b.ID == id {
			m.bodegas = append(m.bodegas[:i], m.bodegas[i+1:]...)
			return true
		}
	}
	return false
}

// Stocks
func (m *almacenMock) ListarStocks() []models.Stock { return m.stocks }
func (m *almacenMock) BuscarStockPorID(id int) (models.Stock, bool) {
	for _, s := range m.stocks {
		if s.ID == id {
			return s, true
		}
	}
	return models.Stock{}, false
}
func (m *almacenMock) CrearStock(s models.Stock) models.Stock {
	s.ID = m.nextID
	m.nextID++
	m.stocks = append(m.stocks, s)
	return s
}
func (m *almacenMock) ActualizarStock(id int, d models.Stock) (models.Stock, bool) {
	for i, s := range m.stocks {
		if s.ID == id {
			d.ID = id
			m.stocks[i] = d
			return d, true
		}
	}
	return models.Stock{}, false
}
func (m *almacenMock) BorrarStock(id int) bool {
	for i, s := range m.stocks {
		if s.ID == id {
			m.stocks = append(m.stocks[:i], m.stocks[i+1:]...)
			return true
		}
	}
	return false
}

// ═══════════════════════════ ESPECIES ════════════════════════════════════════

func TestPescaService_CrearEspecie(t *testing.T) {
	casos := []struct {
		nombre    string
		entrada   models.Especie
		esperaErr error
		llamoRepo bool
	}{
		{"nombre_comun vacío → rechazado", models.Especie{NombreComun: "", UnidadMedida: "kg"}, ErrNombreComunVacio, false},
		{"unidad_medida vacía → rechazado", models.Especie{NombreComun: "Atún", UnidadMedida: ""}, ErrUnidadMedidaVacia, false},
		{"válida → persiste con Estado=true", models.Especie{NombreComun: "Corvina", UnidadMedida: "kg"}, nil, true},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			mock := nuevoMock()
			svc := NewPescaService(mock)
			res, err := svc.CrearEspecie(c.entrada)
			if c.esperaErr != nil {
				require.ErrorIs(t, err, c.esperaErr)
				assert.False(t, mock.llamoCrear)
			} else {
				require.NoError(t, err)
				assert.True(t, mock.llamoCrear)
				assert.True(t, res.Estado)
				assert.NotZero(t, res.ID)
			}
		})
	}
}

func TestPescaService_ObtenerEspecie(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creada := mock.CrearEspecie(models.Especie{NombreComun: "Atún", UnidadMedida: "kg", Estado: true})

	t.Run("ID existente → devuelve especie", func(t *testing.T) {
		e, err := svc.ObtenerEspecie(creada.ID)
		require.NoError(t, err)
		assert.Equal(t, "Atún", e.NombreComun)
	})
	t.Run("ID inexistente → ErrEspecieNoEncontrada", func(t *testing.T) {
		_, err := svc.ObtenerEspecie(9999)
		require.ErrorIs(t, err, ErrEspecieNoEncontrada)
	})
}

func TestPescaService_ActualizarEspecie(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creada := mock.CrearEspecie(models.Especie{NombreComun: "Atún", UnidadMedida: "kg", Estado: true})

	t.Run("datos inválidos → rechazado", func(t *testing.T) {
		_, err := svc.ActualizarEspecie(creada.ID, models.Especie{NombreComun: "", UnidadMedida: "kg"})
		require.ErrorIs(t, err, ErrNombreComunVacio)
	})
	t.Run("ID inexistente → ErrEspecieNoEncontrada", func(t *testing.T) {
		_, err := svc.ActualizarEspecie(9999, models.Especie{NombreComun: "X", UnidadMedida: "kg"})
		require.ErrorIs(t, err, ErrEspecieNoEncontrada)
	})
	t.Run("válido → actualiza", func(t *testing.T) {
		e, err := svc.ActualizarEspecie(creada.ID, models.Especie{NombreComun: "Corvina", UnidadMedida: "kg"})
		require.NoError(t, err)
		assert.Equal(t, "Corvina", e.NombreComun)
	})
}

func TestPescaService_BorrarEspecie(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creada := mock.CrearEspecie(models.Especie{NombreComun: "Atún", UnidadMedida: "kg", Estado: true})

	t.Run("ID inexistente → error", func(t *testing.T) {
		err := svc.BorrarEspecie(9999)
		require.ErrorIs(t, err, ErrEspecieNoEncontrada)
	})
	t.Run("ID existente → borra sin error", func(t *testing.T) {
		err := svc.BorrarEspecie(creada.ID)
		require.NoError(t, err)
	})
}

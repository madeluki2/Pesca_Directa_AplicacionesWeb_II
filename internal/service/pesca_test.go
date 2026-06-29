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

// ═══════════════════════════ PESCADORES ══════════════════════════════════════

func TestPescaService_CrearPescador(t *testing.T) {
	casos := []struct {
		nombre    string
		entrada   models.Pescador
		esperaErr error
	}{
		{"usuario_id=0 → rechazado", models.Pescador{UsuarioID: 0, Cedula: "123", PuertoBase: "Tarqui"}, ErrCredencialesInvalidas},
		{"cedula vacía → rechazado", models.Pescador{UsuarioID: 1, Cedula: "", PuertoBase: "Tarqui"}, ErrCedulaVacia},
		{"puerto_base vacío → rechazado", models.Pescador{UsuarioID: 1, Cedula: "123", PuertoBase: ""}, ErrPuertoVacio},
		{"válido → crea con Estado=true", models.Pescador{UsuarioID: 1, Cedula: "1310045678", PuertoBase: "Tarqui"}, nil},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			mock := nuevoMock()
			svc := NewPescaService(mock)
			res, err := svc.CrearPescador(c.entrada)
			if c.esperaErr != nil {
				require.ErrorIs(t, err, c.esperaErr)
			} else {
				require.NoError(t, err)
				assert.True(t, res.Estado)
				assert.NotZero(t, res.ID)
			}
		})
	}
}

func TestPescaService_ObtenerPescador(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creado := mock.CrearPescador(models.Pescador{UsuarioID: 1, Cedula: "123", PuertoBase: "Manta", Estado: true})

	t.Run("existente → devuelve pescador", func(t *testing.T) {
		p, err := svc.ObtenerPescador(creado.ID)
		require.NoError(t, err)
		assert.Equal(t, "123", p.Cedula)
	})
	t.Run("inexistente → ErrPescadorNoEncontrado", func(t *testing.T) {
		_, err := svc.ObtenerPescador(9999)
		require.ErrorIs(t, err, ErrPescadorNoEncontrado)
	})
}

func TestPescaService_BorrarPescador(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creado := mock.CrearPescador(models.Pescador{UsuarioID: 1, Cedula: "123", PuertoBase: "Manta", Estado: true})

	require.ErrorIs(t, svc.BorrarPescador(9999), ErrPescadorNoEncontrado)
	require.NoError(t, svc.BorrarPescador(creado.ID))
}

// ═══════════════════════════ CAPTURAS ════════════════════════════════════════

func TestPescaService_CrearCaptura(t *testing.T) {
	casos := []struct {
		nombre    string
		entrada   models.Captura
		esperaErr error
	}{
		{"embarcacion_id=0 → rechazado", models.Captura{EmbarcacionID: 0, EspecieID: 1, CantidadKG: 10, Fecha: "2026-06-28", EstadoFrescura: "fresco"}, ErrEmbarcacionNoEncontrada},
		{"especie_id=0 → rechazado", models.Captura{EmbarcacionID: 1, EspecieID: 0, CantidadKG: 10, Fecha: "2026-06-28", EstadoFrescura: "fresco"}, ErrEspecieNoEncontrada},
		{"cantidad_kg=0 → rechazado", models.Captura{EmbarcacionID: 1, EspecieID: 1, CantidadKG: 0, Fecha: "2026-06-28", EstadoFrescura: "fresco"}, ErrCantidadInvalida},
		{"fecha vacía → rechazado", models.Captura{EmbarcacionID: 1, EspecieID: 1, CantidadKG: 10, Fecha: "", EstadoFrescura: "fresco"}, ErrFechaVacia},
		{"frescura inválida → rechazado", models.Captura{EmbarcacionID: 1, EspecieID: 1, CantidadKG: 10, Fecha: "2026-06-28", EstadoFrescura: "podrido"}, ErrFrescuraInvalida},
		{"válida fresco → crea", models.Captura{EmbarcacionID: 1, EspecieID: 1, CantidadKG: 50, Fecha: "2026-06-28", EstadoFrescura: "fresco"}, nil},
		{"válida refrigerado → crea", models.Captura{EmbarcacionID: 1, EspecieID: 1, CantidadKG: 50, Fecha: "2026-06-28", EstadoFrescura: "refrigerado"}, nil},
		{"válida congelado → crea", models.Captura{EmbarcacionID: 1, EspecieID: 1, CantidadKG: 50, Fecha: "2026-06-28", EstadoFrescura: "congelado"}, nil},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			mock := nuevoMock()
			svc := NewPescaService(mock)
			res, err := svc.CrearCaptura(c.entrada)
			if c.esperaErr != nil {
				require.ErrorIs(t, err, c.esperaErr)
			} else {
				require.NoError(t, err)
				assert.NotZero(t, res.ID)
			}
		})
	}
}

func TestPescaService_ObtenerCaptura_Inexistente(t *testing.T) {
	svc := NewPescaService(nuevoMock())
	_, err := svc.ObtenerCaptura(9999)
	require.ErrorIs(t, err, ErrCapturaNoEncontrada)
}

// ═══════════════════════════ BODEGAS ═════════════════════════════════════════

func TestPescaService_CrearBodega(t *testing.T) {
	casos := []struct {
		nombre    string
		entrada   models.Bodega
		esperaErr error
	}{
		{"nombre vacío → rechazado", models.Bodega{Nombre: "", Ubicacion: "Manta", CapacidadKG: 100}, ErrNombreVacio},
		{"ubicacion vacía → rechazado", models.Bodega{Nombre: "B1", Ubicacion: "", CapacidadKG: 100}, ErrUbicacionVacia},
		{"capacidad=0 → rechazado", models.Bodega{Nombre: "B1", Ubicacion: "Manta", CapacidadKG: 0}, ErrCapacidadInvalida},
		{"válida → crea con Estado=true", models.Bodega{Nombre: "Bodega Central", Ubicacion: "Puerto Tarqui", CapacidadKG: 5000}, nil},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			mock := nuevoMock()
			svc := NewPescaService(mock)
			res, err := svc.CrearBodega(c.entrada)
			if c.esperaErr != nil {
				require.ErrorIs(t, err, c.esperaErr)
			} else {
				require.NoError(t, err)
				assert.True(t, res.Estado)
				assert.NotZero(t, res.ID)
			}
		})
	}
}

func TestPescaService_BorrarBodega(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creada := mock.CrearBodega(models.Bodega{Nombre: "B1", Ubicacion: "Manta", CapacidadKG: 100, Estado: true})

	require.ErrorIs(t, svc.BorrarBodega(9999), ErrBodegaNoEncontrada)
	require.NoError(t, svc.BorrarBodega(creada.ID))
}

// ═══════════════════════════ STOCKS ══════════════════════════════════════════

func TestPescaService_CrearStock(t *testing.T) {
	casos := []struct {
		nombre    string
		entrada   models.Stock
		esperaErr error
	}{
		{"bodega_id=0 → rechazado", models.Stock{BodegaID: 0, EspecieID: 1, CantidadKG: 10, FechaIngreso: "2026-06-28"}, ErrBodegaNoEncontrada},
		{"especie_id=0 → rechazado", models.Stock{BodegaID: 1, EspecieID: 0, CantidadKG: 10, FechaIngreso: "2026-06-28"}, ErrEspecieNoEncontrada},
		{"cantidad negativa → rechazado", models.Stock{BodegaID: 1, EspecieID: 1, CantidadKG: -1, FechaIngreso: "2026-06-28"}, ErrCantidadNegativa},
		{"fecha_ingreso vacía → rechazado", models.Stock{BodegaID: 1, EspecieID: 1, CantidadKG: 10, FechaIngreso: ""}, ErrFechaIngresoVacia},
		{"válido con cantidad > 0 → Estado=true", models.Stock{BodegaID: 1, EspecieID: 1, CantidadKG: 50, FechaIngreso: "2026-06-28"}, nil},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			mock := nuevoMock()
			svc := NewPescaService(mock)
			res, err := svc.CrearStock(c.entrada)
			if c.esperaErr != nil {
				require.ErrorIs(t, err, c.esperaErr)
			} else {
				require.NoError(t, err)
				assert.True(t, res.Estado, "stock con cantidad > 0 debe tener Estado=true")
				assert.NotZero(t, res.ID)
			}
		})
	}
}

func TestPescaService_ObtenerStock_Inexistente(t *testing.T) {
	svc := NewPescaService(nuevoMock())
	_, err := svc.ObtenerStock(9999)
	require.ErrorIs(t, err, ErrStockNoEncontrado)
}

func TestPescaService_ListarEspecies(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	// lista vacía al inicio
	assert.Empty(t, svc.ListarEspecies())
	// después de crear, aparece
	mock.CrearEspecie(models.Especie{NombreComun: "Atún", UnidadMedida: "kg", Estado: true})
	assert.Len(t, svc.ListarEspecies(), 1)
}

func TestPescaService_ListarPescadores(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	assert.Empty(t, svc.ListarPescadores())
	mock.CrearPescador(models.Pescador{UsuarioID: 1, Cedula: "123", PuertoBase: "Manta", Estado: true})
	assert.Len(t, svc.ListarPescadores(), 1)
}

func TestPescaService_ListarCapturas(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	assert.Empty(t, svc.ListarCapturas())
	mock.CrearCaptura(models.Captura{EmbarcacionID: 1, EspecieID: 1, CantidadKG: 10, Fecha: "2026-06-28", EstadoFrescura: "fresco"})
	assert.Len(t, svc.ListarCapturas(), 1)
}

func TestPescaService_ListarBodegas(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	assert.Empty(t, svc.ListarBodegas())
	mock.CrearBodega(models.Bodega{Nombre: "B1", Ubicacion: "Manta", CapacidadKG: 100, Estado: true})
	assert.Len(t, svc.ListarBodegas(), 1)
}

func TestPescaService_ListarStocks(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	assert.Empty(t, svc.ListarStocks())
	mock.CrearStock(models.Stock{BodegaID: 1, EspecieID: 1, CantidadKG: 50, FechaIngreso: "2026-06-28", Estado: true})
	assert.Len(t, svc.ListarStocks(), 1)
}

func TestPescaService_ActualizarPescador(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creado := mock.CrearPescador(models.Pescador{UsuarioID: 1, Cedula: "123", PuertoBase: "Manta", Estado: true})

	t.Run("datos inválidos → rechazado", func(t *testing.T) {
		_, err := svc.ActualizarPescador(creado.ID, models.Pescador{UsuarioID: 0, Cedula: "", PuertoBase: ""})
		require.Error(t, err)
	})
	t.Run("ID inexistente → error", func(t *testing.T) {
		_, err := svc.ActualizarPescador(9999, models.Pescador{UsuarioID: 1, Cedula: "123", PuertoBase: "Tarqui"})
		require.ErrorIs(t, err, ErrPescadorNoEncontrado)
	})
	t.Run("válido → actualiza", func(t *testing.T) {
		p, err := svc.ActualizarPescador(creado.ID, models.Pescador{UsuarioID: 1, Cedula: "123", PuertoBase: "Tarqui"})
		require.NoError(t, err)
		assert.Equal(t, "Tarqui", p.PuertoBase)
	})
}

func TestPescaService_ActualizarCaptura(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creada := mock.CrearCaptura(models.Captura{EmbarcacionID: 1, EspecieID: 1, CantidadKG: 10, Fecha: "2026-06-28", EstadoFrescura: "fresco"})

	t.Run("ID inexistente → error", func(t *testing.T) {
		_, err := svc.ActualizarCaptura(9999, models.Captura{EmbarcacionID: 1, EspecieID: 1, CantidadKG: 10, Fecha: "2026-06-28", EstadoFrescura: "fresco"})
		require.ErrorIs(t, err, ErrCapturaNoEncontrada)
	})
	t.Run("válida → actualiza", func(t *testing.T) {
		c, err := svc.ActualizarCaptura(creada.ID, models.Captura{EmbarcacionID: 1, EspecieID: 1, CantidadKG: 200, Fecha: "2026-06-28", EstadoFrescura: "congelado"})
		require.NoError(t, err)
		assert.Equal(t, 200.0, c.CantidadKG)
	})
}

func TestPescaService_BorrarCaptura(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creada := mock.CrearCaptura(models.Captura{EmbarcacionID: 1, EspecieID: 1, CantidadKG: 10, Fecha: "2026-06-28", EstadoFrescura: "fresco"})

	require.ErrorIs(t, svc.BorrarCaptura(9999), ErrCapturaNoEncontrada)
	require.NoError(t, svc.BorrarCaptura(creada.ID))
}

func TestPescaService_ActualizarBodega(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creada := mock.CrearBodega(models.Bodega{Nombre: "B1", Ubicacion: "Manta", CapacidadKG: 100, Estado: true})

	t.Run("ID inexistente → error", func(t *testing.T) {
		_, err := svc.ActualizarBodega(9999, models.Bodega{Nombre: "B2", Ubicacion: "Tarqui", CapacidadKG: 200})
		require.ErrorIs(t, err, ErrBodegaNoEncontrada)
	})
	t.Run("datos inválidos → rechazado", func(t *testing.T) {
		_, err := svc.ActualizarBodega(creada.ID, models.Bodega{Nombre: "", Ubicacion: "Manta", CapacidadKG: 100})
		require.ErrorIs(t, err, ErrNombreVacio)
	})
	t.Run("válida → actualiza", func(t *testing.T) {
		b, err := svc.ActualizarBodega(creada.ID, models.Bodega{Nombre: "Bodega Norte", Ubicacion: "Tarqui", CapacidadKG: 3000})
		require.NoError(t, err)
		assert.Equal(t, "Bodega Norte", b.Nombre)
	})
}

func TestPescaService_ObtenerBodega(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creada := mock.CrearBodega(models.Bodega{Nombre: "B1", Ubicacion: "Manta", CapacidadKG: 100, Estado: true})

	b, err := svc.ObtenerBodega(creada.ID)
	require.NoError(t, err)
	assert.Equal(t, "B1", b.Nombre)

	_, err = svc.ObtenerBodega(9999)
	require.ErrorIs(t, err, ErrBodegaNoEncontrada)
}

func TestPescaService_ActualizarStock(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creado := mock.CrearStock(models.Stock{BodegaID: 1, EspecieID: 1, CantidadKG: 50, FechaIngreso: "2026-06-28", Estado: true})

	t.Run("ID inexistente → error", func(t *testing.T) {
		_, err := svc.ActualizarStock(9999, models.Stock{BodegaID: 1, EspecieID: 1, CantidadKG: 10, FechaIngreso: "2026-06-28"})
		require.ErrorIs(t, err, ErrStockNoEncontrado)
	})
	t.Run("válido → actualiza", func(t *testing.T) {
		s, err := svc.ActualizarStock(creado.ID, models.Stock{BodegaID: 1, EspecieID: 1, CantidadKG: 300, FechaIngreso: "2026-06-28"})
		require.NoError(t, err)
		assert.Equal(t, 300.0, s.CantidadKG)
	})
}

func TestPescaService_BorrarStock(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creado := mock.CrearStock(models.Stock{BodegaID: 1, EspecieID: 1, CantidadKG: 50, FechaIngreso: "2026-06-28", Estado: true})

	require.ErrorIs(t, svc.BorrarStock(9999), ErrStockNoEncontrado)
	require.NoError(t, svc.BorrarStock(creado.ID))
}

// ═══════════════════════════ EMBARCACIONES ════════════════════════════════════

func TestPescaService_CrearEmbarcacion(t *testing.T) {
	casos := []struct {
		nombre    string
		entrada   models.Embarcacion
		esperaErr error
	}{
		{"pescador_id=0 → rechazado", models.Embarcacion{PescadorID: 0, Nombre: "X", Matricula: "Y", CapacidadKG: 100}, ErrPescadorNoEncontrado},
		{"nombre vacío → rechazado", models.Embarcacion{PescadorID: 1, Nombre: "", Matricula: "Y", CapacidadKG: 100}, ErrNombreVacio},
		{"matricula vacía → rechazado", models.Embarcacion{PescadorID: 1, Nombre: "X", Matricula: "", CapacidadKG: 100}, ErrMatriculaVacia},
		{"válida → crea con Estado=true", models.Embarcacion{PescadorID: 1, Nombre: "La Esperanza", Matricula: "MB-0451", CapacidadKG: 800}, nil},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			mock := nuevoMock()
			svc := NewPescaService(mock)
			res, err := svc.CrearEmbarcacion(c.entrada)
			if c.esperaErr != nil {
				require.ErrorIs(t, err, c.esperaErr)
			} else {
				require.NoError(t, err)
				assert.True(t, res.Estado)
				assert.NotZero(t, res.ID)
			}
		})
	}
}

func TestPescaService_ListarEmbarcaciones(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	assert.Empty(t, svc.ListarEmbarcaciones())
	mock.CrearEmbarcacion(models.Embarcacion{PescadorID: 1, Nombre: "X", Matricula: "Y", CapacidadKG: 100, Estado: true})
	assert.Len(t, svc.ListarEmbarcaciones(), 1)
}

func TestPescaService_ObtenerEmbarcacion(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creada := mock.CrearEmbarcacion(models.Embarcacion{PescadorID: 1, Nombre: "X", Matricula: "Y", CapacidadKG: 100, Estado: true})

	e, err := svc.ObtenerEmbarcacion(creada.ID)
	require.NoError(t, err)
	assert.Equal(t, "X", e.Nombre)

	_, err = svc.ObtenerEmbarcacion(9999)
	require.ErrorIs(t, err, ErrEmbarcacionNoEncontrada)
}

func TestPescaService_BorrarEmbarcacion(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creada := mock.CrearEmbarcacion(models.Embarcacion{PescadorID: 1, Nombre: "X", Matricula: "Y", CapacidadKG: 100, Estado: true})

	require.ErrorIs(t, svc.BorrarEmbarcacion(9999), ErrEmbarcacionNoEncontrada)
	require.NoError(t, svc.BorrarEmbarcacion(creada.ID))
}

func TestPescaService_ActualizarEmbarcacion(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creada := mock.CrearEmbarcacion(models.Embarcacion{PescadorID: 1, Nombre: "X", Matricula: "Y", CapacidadKG: 100, Estado: true})

	t.Run("datos inválidos → rechazado", func(t *testing.T) {
		_, err := svc.ActualizarEmbarcacion(creada.ID, models.Embarcacion{PescadorID: 0, Nombre: "", Matricula: ""})
		require.Error(t, err)
	})
	t.Run("ID inexistente → error", func(t *testing.T) {
		_, err := svc.ActualizarEmbarcacion(9999, models.Embarcacion{PescadorID: 1, Nombre: "X", Matricula: "Y"})
		require.ErrorIs(t, err, ErrEmbarcacionNoEncontrada)
	})
	t.Run("válida → actualiza", func(t *testing.T) {
		e, err := svc.ActualizarEmbarcacion(creada.ID, models.Embarcacion{PescadorID: 1, Nombre: "Nueva", Matricula: "MB-999"})
		require.NoError(t, err)
		assert.Equal(t, "Nueva", e.Nombre)
	})
}

func TestPescaService_ObtenerCaptura_Existente(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creada := mock.CrearCaptura(models.Captura{EmbarcacionID: 1, EspecieID: 1, CantidadKG: 50, Fecha: "2026-06-28", EstadoFrescura: "fresco"})

	c, err := svc.ObtenerCaptura(creada.ID)
	require.NoError(t, err)
	assert.Equal(t, 50.0, c.CantidadKG)
}

func TestPescaService_ObtenerStock_Existente(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creado := mock.CrearStock(models.Stock{BodegaID: 1, EspecieID: 1, CantidadKG: 30, FechaIngreso: "2026-06-28", Estado: true})

	s, err := svc.ObtenerStock(creado.ID)
	require.NoError(t, err)
	assert.Equal(t, 30.0, s.CantidadKG)
}

func TestPescaService_ActualizarCaptura_DatosInvalidos(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creada := mock.CrearCaptura(models.Captura{EmbarcacionID: 1, EspecieID: 1, CantidadKG: 10, Fecha: "2026-06-28", EstadoFrescura: "fresco"})

	// frescura inválida en actualización
	_, err := svc.ActualizarCaptura(creada.ID, models.Captura{EmbarcacionID: 1, EspecieID: 1, CantidadKG: 10, Fecha: "2026-06-28", EstadoFrescura: "mala"})
	require.ErrorIs(t, err, ErrFrescuraInvalida)
}

func TestPescaService_ActualizarStock_DatosInvalidos(t *testing.T) {
	mock := nuevoMock()
	svc := NewPescaService(mock)
	creado := mock.CrearStock(models.Stock{BodegaID: 1, EspecieID: 1, CantidadKG: 50, FechaIngreso: "2026-06-28", Estado: true})

	// fecha vacía en actualización
	_, err := svc.ActualizarStock(creado.ID, models.Stock{BodegaID: 1, EspecieID: 1, CantidadKG: 10, FechaIngreso: ""})
	require.ErrorIs(t, err, ErrFechaIngresoVacia)
}

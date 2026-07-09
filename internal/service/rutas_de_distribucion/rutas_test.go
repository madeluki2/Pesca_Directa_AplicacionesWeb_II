package rutas_de_distribucion

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	. "Pesca_Directa_AplicacionesWeb_II/internal/service"
)

// ─── Mock manual de AlmacenRutas ─────────────────────────────────────────────

type rutaRepoMock struct {
	llamaronRuta          bool
	llamaronTransportista bool
	llamaronPunto         bool
	llamaronEntrega       bool
	guardada              models.Ruta

	listarRutasFunc     func() []models.Ruta
	buscarRutaPorIDFunc func(id uint) (models.Ruta, bool)
	actualizarRutaFunc  func(id uint, d models.Ruta) (models.Ruta, bool)
	borrarRutaFunc      func(id uint) bool

	listarPuntosFunc     func() []models.Punto
	buscarPuntoPorIDFunc func(id uint) (models.Punto, bool)
	actualizarPuntoFunc  func(id uint, d models.Punto) (models.Punto, bool)
	borrarPuntoFunc      func(id uint) bool

	listarTransportistasFunc     func() []models.Transportista
	buscarTransportistaPorIDFunc func(id uint) (models.Transportista, bool)
	actualizarTransportistaFunc  func(id uint, d models.Transportista) (models.Transportista, bool)
	borrarTransportistaFunc      func(id uint) bool

	listarEntregasFunc     func() []models.EntregaPedido
	buscarEntregaPorIDFunc func(id uint) (models.EntregaPedido, bool)
	actualizarEntregaFunc  func(id uint, d models.EntregaPedido) (models.EntregaPedido, bool)
	borrarEntregaFunc      func(id uint) bool
}

func (m *rutaRepoMock) CrearRuta(r models.Ruta) models.Ruta {
	m.llamaronRuta = true
	r.ID = 1
	m.guardada = r
	return r
}

func (m *rutaRepoMock) CrearTransportista(t models.Transportista) models.Transportista {
	m.llamaronTransportista = true
	t.ID = 1
	return t
}

func (m *rutaRepoMock) CrearPunto(p models.Punto) models.Punto {
	m.llamaronPunto = true
	p.ID = 1
	return p
}

func (m *rutaRepoMock) CrearEntrega(e models.EntregaPedido) models.EntregaPedido {
	m.llamaronEntrega = true
	e.ID = 1
	return e
}

func (m *rutaRepoMock) ListarRutas() []models.Ruta {
	if m.listarRutasFunc != nil {
		return m.listarRutasFunc()
	}
	return nil
}

func (m *rutaRepoMock) BuscarRutaPorID(id uint) (models.Ruta, bool) {
	if m.buscarRutaPorIDFunc != nil {
		return m.buscarRutaPorIDFunc(id)
	}
	return models.Ruta{}, false
}

func (m *rutaRepoMock) ActualizarRuta(id uint, d models.Ruta) (models.Ruta, bool) {
	if m.actualizarRutaFunc != nil {
		return m.actualizarRutaFunc(id, d)
	}
	return d, true
}

func (m *rutaRepoMock) BorrarRuta(id uint) bool {
	if m.borrarRutaFunc != nil {
		return m.borrarRutaFunc(id)
	}
	return true
}

func (m *rutaRepoMock) ListarPuntos() []models.Punto {
	if m.listarPuntosFunc != nil {
		return m.listarPuntosFunc()
	}
	return nil
}

func (m *rutaRepoMock) BuscarPuntoPorID(id uint) (models.Punto, bool) {
	if m.buscarPuntoPorIDFunc != nil {
		return m.buscarPuntoPorIDFunc(id)
	}
	return models.Punto{}, false
}

func (m *rutaRepoMock) ActualizarPunto(id uint, d models.Punto) (models.Punto, bool) {
	if m.actualizarPuntoFunc != nil {
		return m.actualizarPuntoFunc(id, d)
	}
	return d, true
}

func (m *rutaRepoMock) BorrarPunto(id uint) bool {
	if m.borrarPuntoFunc != nil {
		return m.borrarPuntoFunc(id)
	}
	return true
}

func (m *rutaRepoMock) ListarTransportistas() []models.Transportista {
	if m.listarTransportistasFunc != nil {
		return m.listarTransportistasFunc()
	}
	return nil
}

func (m *rutaRepoMock) BuscarTransportistaPorID(id uint) (models.Transportista, bool) {
	if m.buscarTransportistaPorIDFunc != nil {
		return m.buscarTransportistaPorIDFunc(id)
	}
	return models.Transportista{}, false
}

func (m *rutaRepoMock) ActualizarTransportista(id uint, d models.Transportista) (models.Transportista, bool) {
	if m.actualizarTransportistaFunc != nil {
		return m.actualizarTransportistaFunc(id, d)
	}
	return d, true
}

func (m *rutaRepoMock) BorrarTransportista(id uint) bool {
	if m.borrarTransportistaFunc != nil {
		return m.borrarTransportistaFunc(id)
	}
	return true
}

func (m *rutaRepoMock) ListarEntregas() []models.EntregaPedido {
	if m.listarEntregasFunc != nil {
		return m.listarEntregasFunc()
	}
	return nil
}

func (m *rutaRepoMock) BuscarEntregaPorID(id uint) (models.EntregaPedido, bool) {
	if m.buscarEntregaPorIDFunc != nil {
		return m.buscarEntregaPorIDFunc(id)
	}
	return models.EntregaPedido{}, false
}

func (m *rutaRepoMock) ActualizarEntrega(id uint, d models.EntregaPedido) (models.EntregaPedido, bool) {
	if m.actualizarEntregaFunc != nil {
		return m.actualizarEntregaFunc(id, d)
	}
	return d, true
}

func (m *rutaRepoMock) BorrarEntrega(id uint) bool {
	if m.borrarEntregaFunc != nil {
		return m.borrarEntregaFunc(id)
	}
	return true
}

// ─── Test CrearRuta ─────────────────────────────────────────────────────────

func TestRutasService_CrearRuta(t *testing.T) {
	casos := []struct {
		nombre    string
		entrada   models.Ruta
		esperaErr error
		llamoRepo bool
	}{
		{
			nombre:    "nombre vacío → rechazado sin tocar el repo",
			entrada:   models.Ruta{Nombre: "", Origen: "Puerto", Destino: "Mercado"},
			esperaErr: ErrNombreVacio,
			llamoRepo: false,
		},
		{
			nombre:    "origen vacío → rechazado sin tocar el repo",
			entrada:   models.Ruta{Nombre: "Ruta 1", Origen: "", Destino: "Mercado"},
			esperaErr: ErrOrigenVacio,
			llamoRepo: false,
		},
		{
			nombre:    "destino vacío → rechazado sin tocar el repo",
			entrada:   models.Ruta{Nombre: "Ruta 1", Origen: "Puerto", Destino: ""},
			esperaErr: ErrDestinoVacio,
			llamoRepo: false,
		},
		{
			nombre:    "ruta válida → persiste con estado activo",
			entrada:   models.Ruta{Nombre: "Ruta Norte", Origen: "Puerto Tarqui", Destino: "Mercado Central"},
			esperaErr: nil,
			llamoRepo: true,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := &rutaRepoMock{}
			svc := NewRutasService(repo)

			resultado, err := svc.CrearRuta(c.entrada)

			if c.esperaErr != nil {
				require.ErrorIs(t, err, c.esperaErr)
				assert.False(t, repo.llamaronRuta, "el repositorio fue llamado aunque la validación debería haber fallado primero")
			} else {
				require.NoError(t, err)
				assert.True(t, repo.llamaronRuta, "el repositorio nunca fue invocado para guardar la ruta")
				assert.Equal(t, "activo", resultado.Estado, "esperaba Estado=activo en la ruta creada")
				assert.NotZero(t, resultado.ID, "esperaba que el repo asignara un ID > 0")
			}
		})
	}
}

// ─── Test CrearTransportista ────────────────────────────────────────────────

func TestRutasService_CrearTransportista(t *testing.T) {
	casos := []struct {
		nombre    string
		entrada   models.Transportista
		esperaErr error
		llamoRepo bool
	}{
		{
			nombre:    "nombre vacío → rechazado sin tocar el repo",
			entrada:   models.Transportista{Nombre: "", Telefono: "0991234567", PlacaVehiculo: "ABC-123"},
			esperaErr: ErrNombreVacio,
			llamoRepo: false,
		},
		{
			nombre:    "teléfono vacío → rechazado sin tocar el repo",
			entrada:   models.Transportista{Nombre: "Carlos", Telefono: "", PlacaVehiculo: "ABC-123"},
			esperaErr: ErrTelefonoVacio,
			llamoRepo: false,
		},
		{
			nombre:    "placa vacía → rechazado sin tocar el repo",
			entrada:   models.Transportista{Nombre: "Carlos", Telefono: "0991234567", PlacaVehiculo: ""},
			esperaErr: ErrPlacaVacia,
			llamoRepo: false,
		},
		{
			nombre:    "transportista válido → persiste con estado activo",
			entrada:   models.Transportista{Nombre: "Carlos Pérez", Telefono: "0991234567", PlacaVehiculo: "ABC-123"},
			esperaErr: nil,
			llamoRepo: true,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := &rutaRepoMock{}
			svc := NewRutasService(repo)

			resultado, err := svc.CrearTransportista(c.entrada)

			if c.esperaErr != nil {
				require.ErrorIs(t, err, c.esperaErr)
				assert.False(t, repo.llamaronTransportista, "el repositorio fue llamado aunque la validación debería haber fallado primero")
			} else {
				require.NoError(t, err)
				assert.True(t, repo.llamaronTransportista, "el repositorio nunca fue invocado para guardar el transportista")
				assert.Equal(t, "activo", resultado.Estado, "esperaba Estado=activo en el transportista creado")
			}
		})
	}
}

// ═══════════════════════════ RUTAS — tests ══════════════════════════════════

func TestRutasService_ListarRutas(t *testing.T) {
	esperadas := []models.Ruta{
		{ID: 1, Nombre: "Ruta Norte", Origen: "Puerto", Destino: "Mercado"},
		{ID: 2, Nombre: "Ruta Sur", Origen: "Puerto", Destino: "Mercado"},
	}
	repo := &rutaRepoMock{listarRutasFunc: func() []models.Ruta { return esperadas }}
	svc := NewRutasService(repo)

	resultado := svc.ListarRutas()

	assert.Equal(t, esperadas, resultado)
}

func TestRutasService_ObtenerRuta(t *testing.T) {
	t.Run("ruta existente → la devuelve sin error", func(t *testing.T) {
		esperada := models.Ruta{ID: 5, Nombre: "Ruta Norte", Origen: "Puerto", Destino: "Mercado"}
		repo := &rutaRepoMock{
			buscarRutaPorIDFunc: func(id uint) (models.Ruta, bool) {
				assert.Equal(t, uint(5), id, "se consultó un ID distinto al esperado")
				return esperada, true
			},
		}
		svc := NewRutasService(repo)

		resultado, err := svc.ObtenerRuta(5)

		require.NoError(t, err)
		assert.Equal(t, esperada, resultado)
	})

	t.Run("ruta inexistente → ErrRutaNoEncontrada", func(t *testing.T) {
		repo := &rutaRepoMock{
			buscarRutaPorIDFunc: func(id uint) (models.Ruta, bool) { return models.Ruta{}, false },
		}
		svc := NewRutasService(repo)

		_, err := svc.ObtenerRuta(999)

		require.ErrorIs(t, err, ErrRutaNoEncontrada)
	})
}

func TestRutasService_BorrarRuta(t *testing.T) {
	t.Run("ruta inexistente → ErrRutaNoEncontrada", func(t *testing.T) {
		repo := &rutaRepoMock{borrarRutaFunc: func(id uint) bool { return false }}
		svc := NewRutasService(repo)

		err := svc.BorrarRuta(1)

		require.ErrorIs(t, err, ErrRutaNoEncontrada)
	})

	t.Run("ruta existente → se borra sin error", func(t *testing.T) {
		repo := &rutaRepoMock{borrarRutaFunc: func(id uint) bool {
			assert.Equal(t, uint(3), id)
			return true
		}}
		svc := NewRutasService(repo)

		err := svc.BorrarRuta(3)

		require.NoError(t, err)
	})
}

// ═══════════════════════════ PUNTOS — tests ═════════════════════════════════

func TestRutasService_ListarPuntos(t *testing.T) {
	esperados := []models.Punto{{ID: 1, RutaID: 1, Nombre: "Muelle", Direccion: "Av. Principal"}}
	repo := &rutaRepoMock{listarPuntosFunc: func() []models.Punto { return esperados }}
	svc := NewRutasService(repo)

	resultado := svc.ListarPuntos()

	assert.Equal(t, esperados, resultado)
}

func TestRutasService_CrearPunto(t *testing.T) {
	casos := []struct {
		nombre    string
		entrada   models.Punto
		esperaErr error
	}{
		{
			nombre:    "RutaID vacío → rechazado",
			entrada:   models.Punto{RutaID: 0, Nombre: "Muelle", Direccion: "Av. Principal"},
			esperaErr: ErrRutaIDVacio,
		},
		{
			nombre:    "nombre vacío → rechazado",
			entrada:   models.Punto{RutaID: 1, Nombre: "", Direccion: "Av. Principal"},
			esperaErr: ErrNombreVacio,
		},
		{
			nombre:    "dirección vacía → rechazado",
			entrada:   models.Punto{RutaID: 1, Nombre: "Muelle", Direccion: ""},
			esperaErr: ErrDireccionVacia,
		},
		{
			nombre:    "punto válido → persiste con estado activo",
			entrada:   models.Punto{RutaID: 1, Nombre: "Muelle", Direccion: "Av. Principal"},
			esperaErr: nil,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := &rutaRepoMock{}
			svc := NewRutasService(repo)

			resultado, err := svc.CrearPunto(c.entrada)

			if c.esperaErr != nil {
				require.ErrorIs(t, err, c.esperaErr)
				assert.False(t, repo.llamaronPunto, "el repositorio fue llamado aunque la validación debería haber fallado primero")
			} else {
				require.NoError(t, err)
				assert.True(t, repo.llamaronPunto, "el repositorio nunca fue invocado para guardar el punto")
				assert.Equal(t, "activo", resultado.Estado)
			}
		})
	}
}

func TestRutasService_placaYaExiste(t *testing.T) {
	repo := &rutaRepoMock{
		listarTransportistasFunc: func() []models.Transportista {
			return []models.Transportista{
				{ID: 1, PlacaVehiculo: "ABC-123"},
				{ID: 2, PlacaVehiculo: "XYZ-999"},
			}
		},
	}
	svc := NewRutasService(repo)

	assert.True(t, svc.placaYaExiste("ABC-123", 0), "la placa existe y debería detectarse")
	assert.False(t, svc.placaYaExiste("ABC-123", 1), "la placa es del mismo ID que se está ignorando")
	assert.False(t, svc.placaYaExiste("NUEVA-000", 0), "la placa no existe en la lista")
}

// ═══════════════════════════ ENTREGAS — tests ═══════════════════════════════

func TestRutasService_ListarEntregas(t *testing.T) {
	esperadas := []models.EntregaPedido{{ID: 1, PedidoID: 1, PuntoID: 1, TransportistaID: 1}}
	repo := &rutaRepoMock{listarEntregasFunc: func() []models.EntregaPedido { return esperadas }}
	svc := NewRutasService(repo)

	resultado := svc.ListarEntregas()

	assert.Equal(t, esperadas, resultado)
}

func TestRutasService_CrearEntrega(t *testing.T) {
	casos := []struct {
		nombre    string
		entrada   models.EntregaPedido
		esperaErr error
	}{
		{
			nombre:    "PedidoID vacío → rechazado",
			entrada:   models.EntregaPedido{PedidoID: 0, PuntoID: 1, TransportistaID: 1},
			esperaErr: ErrPedidoIDVacio,
		},
		{
			nombre:    "PuntoID vacío → rechazado",
			entrada:   models.EntregaPedido{PedidoID: 1, PuntoID: 0, TransportistaID: 1},
			esperaErr: ErrPuntoIDVacio,
		},
		{
			nombre:    "TransportistaID vacío → rechazado",
			entrada:   models.EntregaPedido{PedidoID: 1, PuntoID: 1, TransportistaID: 0},
			esperaErr: ErrTransportistaIDVacio,
		},
		{
			nombre:    "entrega válida → persiste con estado pendiente",
			entrada:   models.EntregaPedido{PedidoID: 1, PuntoID: 1, TransportistaID: 1},
			esperaErr: nil,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := &rutaRepoMock{}
			svc := NewRutasService(repo)

			resultado, err := svc.CrearEntrega(c.entrada)

			if c.esperaErr != nil {
				require.ErrorIs(t, err, c.esperaErr)
				assert.False(t, repo.llamaronEntrega, "el repositorio fue llamado aunque la validación debería haber fallado primero")
			} else {
				require.NoError(t, err)
				assert.True(t, repo.llamaronEntrega, "el repositorio nunca fue invocado para guardar la entrega")
				assert.Equal(t, "pendiente", resultado.Estado, "esperaba Estado=pendiente en la entrega creada")
			}
		})
	}
}

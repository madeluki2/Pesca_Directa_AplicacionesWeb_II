package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// ─── Mock manual de AlmacenRutas ─────────────────────────────────────────────

type rutaRepoMock struct {
	llamaronRuta          bool
	llamaronTransportista bool
	guardada              models.Ruta
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

func (m *rutaRepoMock) ListarRutas() []models.Ruta {
	return nil
}

func (m *rutaRepoMock) BuscarRutaPorID(id uint) (models.Ruta, bool) {
	return models.Ruta{}, false
}

func (m *rutaRepoMock) ActualizarRuta(id uint, d models.Ruta) (models.Ruta, bool) {
	return d, true
}

func (m *rutaRepoMock) BorrarRuta(id uint) bool {
	return true
}

func (m *rutaRepoMock) ListarPuntos() []models.Punto                                 { return nil }
func (m *rutaRepoMock) BuscarPuntoPorID(id uint) (models.Punto, bool)                { return models.Punto{}, false }
func (m *rutaRepoMock) CrearPunto(p models.Punto) models.Punto                       { return p }
func (m *rutaRepoMock) ActualizarPunto(id uint, d models.Punto) (models.Punto, bool) { return d, true }
func (m *rutaRepoMock) BorrarPunto(id uint) bool                                     { return true }
func (m *rutaRepoMock) ListarTransportistas() []models.Transportista                 { return nil }
func (m *rutaRepoMock) BuscarTransportistaPorID(id uint) (models.Transportista, bool) {
	return models.Transportista{}, false
}
func (m *rutaRepoMock) ActualizarTransportista(id uint, d models.Transportista) (models.Transportista, bool) {
	return d, true
}
func (m *rutaRepoMock) BorrarTransportista(id uint) bool       { return true }
func (m *rutaRepoMock) ListarEntregas() []models.EntregaPedido { return nil }
func (m *rutaRepoMock) BuscarEntregaPorID(id uint) (models.EntregaPedido, bool) {
	return models.EntregaPedido{}, false
}
func (m *rutaRepoMock) CrearEntrega(e models.EntregaPedido) models.EntregaPedido { return e }
func (m *rutaRepoMock) ActualizarEntrega(id uint, d models.EntregaPedido) (models.EntregaPedido, bool) {
	return d, true
}
func (m *rutaRepoMock) BorrarEntrega(id uint) bool { return true }

// ─── Test CrearRuta ───────────────────────────────────────────────────────────

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

// ─── Test CrearTransportista ──────────────────────────────────────────────────

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

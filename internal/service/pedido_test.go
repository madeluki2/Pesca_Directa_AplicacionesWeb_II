package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

// clienteRepoMock es un doble de prueba que implementa storage.Almacen completo.
type clienteRepoMock struct {
	mock.Mock
}

// ── ClienteRepository ──────────────────────────────────────────────────────────
func (m *clienteRepoMock) ListarClientes() []models.Cliente {
	args := m.Called()
	return args.Get(0).([]models.Cliente)
}
func (m *clienteRepoMock) BuscarClientePorID(id int) (models.Cliente, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Cliente), args.Bool(1)
}
func (m *clienteRepoMock) CrearCliente(c models.Cliente) models.Cliente {
	args := m.Called(c)
	return args.Get(0).(models.Cliente)
}
func (m *clienteRepoMock) ActualizarCliente(id int, datos models.Cliente) (models.Cliente, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.Cliente), args.Bool(1)
}
func (m *clienteRepoMock) EliminarCliente(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}
func (m *clienteRepoMock) CambiarTipoCliente(id int, tipo string) (models.Cliente, bool) {
	args := m.Called(id, tipo)
	return args.Get(0).(models.Cliente), args.Bool(1)
}

// ── PedidoRepository ───────────────────────────────────────────────────────────
func (m *clienteRepoMock) ListarPedidos() []models.Pedido { return nil }
func (m *clienteRepoMock) BuscarPedidoPorID(id int) (models.Pedido, bool) {
	return models.Pedido{}, false
}
func (m *clienteRepoMock) CrearPedido(p models.Pedido) models.Pedido { return p }
func (m *clienteRepoMock) ActualizarPedido(id int, d models.Pedido) (models.Pedido, bool) {
	return d, true
}
func (m *clienteRepoMock) EliminarPedido(id int) bool { return true }

// ── DetalleRepository ──────────────────────────────────────────────────────────
func (m *clienteRepoMock) ListarDetalles() []models.DetallePedido { return nil }
func (m *clienteRepoMock) BuscarDetallePorID(id int) (models.DetallePedido, bool) {
	return models.DetallePedido{}, false
}
func (m *clienteRepoMock) CrearDetalle(d models.DetallePedido) models.DetallePedido { return d }
func (m *clienteRepoMock) ActualizarDetalle(id int, d models.DetallePedido) (models.DetallePedido, bool) {
	return d, true
}
func (m *clienteRepoMock) EliminarDetalle(id int) bool { return true }

// ── Compatibilidad con la Fábrica ──────────────────────────────────────────────
func (m *clienteRepoMock) Seed() {}

// Red de seguridad en tiempo de compilación.
var _ storage.Almacen = (*clienteRepoMock)(nil)

func TestPedidoService_CrearCliente(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Cliente
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre:        "nombre negocio vacío → ErrNombreNegocioVacio",
			entrada:       models.Cliente{TipoCliente: "restaurante", Direccion: "Av. 1", Telefono: "0991111111"},
			errEsperado:   service.ErrNombreNegocioVacio,
			debePersistir: false,
		},
		{
			nombre:        "tipo cliente inválido → ErrTipoClienteInvalido",
			entrada:       models.Cliente{NombreNegocio: "Test", TipoCliente: "invalido", Direccion: "Av. 1", Telefono: "0991111111"},
			errEsperado:   service.ErrTipoClienteInvalido,
			debePersistir: false,
		},
		{
			nombre:        "teléfono vacío → ErrTelefonoVacio",
			entrada:       models.Cliente{NombreNegocio: "Test", TipoCliente: "restaurante", Direccion: "Av. 1"},
			errEsperado:   service.ErrTelefonoVacio,
			debePersistir: false,
		},
		{
			nombre:        "cliente válido → sin error y se persiste",
			entrada:       models.Cliente{NombreNegocio: "Sushi Koi", TipoCliente: "restaurante", Direccion: "Av. Flavio Reyes", Telefono: "0991234567"},
			errEsperado:   nil,
			debePersistir: true,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(clienteRepoMock)
			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 1
				guardado.Estado = "activo"
				repo.On("CrearCliente", mock.Anything).Return(guardado)
			}
			svc := service.NewPedidoService(repo)

			creado, err := svc.CrearCliente(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearCliente")
			} else {
				require.NoError(t, err)
				assert.Equal(t, 1, creado.ID)
				assert.Equal(t, "activo", creado.Estado)
				repo.AssertCalled(t, "CrearCliente", mock.Anything)
			}
		})
	}
}

func TestPedidoService_ObtenerCliente_NoEncontrado(t *testing.T) {
	repo := new(clienteRepoMock)
	repo.On("BuscarClientePorID", 999).Return(models.Cliente{}, false)
	svc := service.NewPedidoService(repo)

	_, err := svc.ObtenerCliente(999)

	require.ErrorIs(t, err, service.ErrClienteNoEncontrado)
	repo.AssertExpectations(t)
}

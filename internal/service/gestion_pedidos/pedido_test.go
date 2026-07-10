package gestion_pedidos

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

type mockAlmacen struct {
	mock.Mock
}

// --- Métodos Reales que testeamos ---
func (m *mockAlmacen) CrearPedido(p models.Pedido) (models.Pedido, error) {
	args := m.Called(p)
	return args.Get(0).(models.Pedido), args.Error(1)
}
func (m *mockAlmacen) BuscarPedidoPorID(id int) (models.Pedido, error) {
	args := m.Called(id)
	return args.Get(0).(models.Pedido), args.Error(1)
}
func (m *mockAlmacen) CrearCliente(c models.Cliente) (models.Cliente, error) {
	args := m.Called(c)
	return args.Get(0).(models.Cliente), args.Error(1)
}
func (m *mockAlmacen) CrearDetalle(d models.DetallePedido) (models.DetallePedido, error) {
	args := m.Called(d)
	return args.Get(0).(models.DetallePedido), args.Error(1)
}
func (m *mockAlmacen) EliminarPedido(i int) error {
	args := m.Called(i)
	return args.Error(0)
}
func (m *mockAlmacen) ListarPedidos() ([]models.Pedido, error) {
	args := m.Called()
	return args.Get(0).([]models.Pedido), args.Error(1)
}
func (m *mockAlmacen) ActualizarPedido(i int, p models.Pedido) (models.Pedido, error) {
	args := m.Called(i, p)
	return args.Get(0).(models.Pedido), args.Error(1)
}

// --- Métodos Dummy ---
func (m *mockAlmacen) ListarClientes() ([]models.Cliente, error)        { return nil, nil }
func (m *mockAlmacen) BuscarClientePorID(i int) (models.Cliente, error) { return models.Cliente{}, nil }
func (m *mockAlmacen) ActualizarCliente(i int, c models.Cliente) (models.Cliente, error) {
	return models.Cliente{}, nil
}
func (m *mockAlmacen) EliminarCliente(i int) error { return nil }
func (m *mockAlmacen) CambiarTipoCliente(i int, t string) (models.Cliente, error) {
	return models.Cliente{}, nil
}
func (m *mockAlmacen) ListarDetalles() ([]models.DetallePedido, error) { return nil, nil }
func (m *mockAlmacen) BuscarDetallePorID(i int) (models.DetallePedido, error) {
	return models.DetallePedido{}, nil
}
func (m *mockAlmacen) ActualizarDetalle(i int, d models.DetallePedido) (models.DetallePedido, error) {
	return models.DetallePedido{}, nil
}
func (m *mockAlmacen) EliminarDetalle(i int) error { return nil }
func (m *mockAlmacen) Seed()                       {}

// ─── Tests de Pedido ─────────────────────────────────────────────────────────

func TestPedidoService_CrearPedido(t *testing.T) {
	mockRepo := new(mockAlmacen)
	svc := NewPedidoService(mockRepo)

	pedidoInput := models.Pedido{ClienteID: 1, Fecha: time.Now()}
	pedidoEsperado := pedidoInput
	pedidoEsperado.Estado = "pendiente"

	mockRepo.On("CrearPedido", mock.Anything).Return(pedidoEsperado, nil)

	result, err := svc.CrearPedido(pedidoInput)

	assert.NoError(t, err)
	assert.Equal(t, "pendiente", result.Estado)
	mockRepo.AssertExpectations(t)
}

func TestPedidoService_CrearPedido_SinClienteID(t *testing.T) {
	mockRepo := new(mockAlmacen)
	svc := NewPedidoService(mockRepo)

	_, err := svc.CrearPedido(models.Pedido{Fecha: time.Now()})

	assert.EqualError(t, err, ErrClienteIDInvalido.Error())
	mockRepo.AssertNotCalled(t, "CrearPedido")
}

func TestPedidoService_CrearPedido_SinFecha(t *testing.T) {
	mockRepo := new(mockAlmacen)
	svc := NewPedidoService(mockRepo)

	_, err := svc.CrearPedido(models.Pedido{ClienteID: 1})

	assert.EqualError(t, err, ErrFechaVacia.Error())
	mockRepo.AssertNotCalled(t, "CrearPedido")
}

func TestPedidoService_ListarPedidos(t *testing.T) {
	mockRepo := new(mockAlmacen)
	svc := NewPedidoService(mockRepo)

	pedidos := []models.Pedido{
		{ID: 1, ClienteID: 1, Estado: "pendiente"},
		{ID: 2, ClienteID: 2, Estado: "entregado"},
	}
	mockRepo.On("ListarPedidos").Return(pedidos, nil)

	result, err := svc.ListarPedidos()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestPedidoService_ObtenerPedido_Exitoso(t *testing.T) {
	mockRepo := new(mockAlmacen)
	pedidoEsperado := models.Pedido{ID: 1, ClienteID: 1, Estado: "pendiente"}
	mockRepo.On("BuscarPedidoPorID", 1).Return(pedidoEsperado, nil)
	svc := NewPedidoService(mockRepo)

	resultado, err := svc.ObtenerPedido(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, resultado.ID)
	assert.Equal(t, "pendiente", resultado.Estado)
	mockRepo.AssertExpectations(t)
}

func TestPedidoService_ObtenerPedido_NoEncontrado(t *testing.T) {
	mockRepo := new(mockAlmacen)
	mockRepo.On("BuscarPedidoPorID", 999).Return(models.Pedido{}, ErrPedidoNoEncontrado)
	svc := NewPedidoService(mockRepo)

	_, err := svc.ObtenerPedido(999)

	assert.EqualError(t, err, ErrPedidoNoEncontrado.Error())
	mockRepo.AssertExpectations(t)
}

func TestPedidoService_ActualizarPedido_Exitoso(t *testing.T) {
	mockRepo := new(mockAlmacen)
	actualizado := models.Pedido{ID: 1, ClienteID: 1, Estado: "entregado"}
	mockRepo.On("ActualizarPedido", 1, mock.Anything).Return(actualizado, nil)
	svc := NewPedidoService(mockRepo)

	resultado, err := svc.ActualizarPedido(1, models.Pedido{Estado: "entregado"})

	assert.NoError(t, err)
	assert.Equal(t, "entregado", resultado.Estado)
	mockRepo.AssertExpectations(t)
}

func TestPedidoService_EliminarPedido_Exitoso(t *testing.T) {
	mockRepo := new(mockAlmacen)
	mockRepo.On("EliminarPedido", 1).Return(nil)
	svc := NewPedidoService(mockRepo)

	err := svc.EliminarPedido(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPedidoService_EliminarPedido_NoEncontrado(t *testing.T) {
	mockRepo := new(mockAlmacen)
	mockRepo.On("EliminarPedido", 999).Return(ErrPedidoNoEncontrado)
	svc := NewPedidoService(mockRepo)

	err := svc.EliminarPedido(999)

	assert.EqualError(t, err, ErrPedidoNoEncontrado.Error())
	mockRepo.AssertExpectations(t)
}

func TestPedidoService_CrearCliente_Valido(t *testing.T) {
	mockRepo := new(mockAlmacen)
	svc := NewPedidoService(mockRepo)

	entrada := models.Cliente{
		NombreNegocio: "Sushi Koi",
		TipoCliente:   "restaurante",
		Telefono:      "0991234567",
		Direccion:     "Av. Flavio Reyes",
	}
	esperado := entrada
	esperado.ID = 1
	esperado.Estado = "activo"

	mockRepo.On("CrearCliente", mock.AnythingOfType("models.Cliente")).Return(esperado, nil)

	creado, err := svc.CrearCliente(entrada)

	assert.NoError(t, err)
	assert.Equal(t, "activo", creado.Estado)
	assert.Equal(t, 1, creado.ID)
	mockRepo.AssertExpectations(t)
}

func TestPedidoService_CrearCliente_NombreVacio(t *testing.T) {
	mockRepo := new(mockAlmacen)
	svc := NewPedidoService(mockRepo)

	_, err := svc.CrearCliente(models.Cliente{
		TipoCliente: "restaurante",
		Telefono:    "0991234567",
		Direccion:   "Av. 1",
	})

	assert.EqualError(t, err, ErrNombreNegocioVacio.Error())
	mockRepo.AssertNotCalled(t, "CrearCliente")
}

func TestPedidoService_CrearCliente_TipoInvalido(t *testing.T) {
	mockRepo := new(mockAlmacen)
	svc := NewPedidoService(mockRepo)

	_, err := svc.CrearCliente(models.Cliente{
		NombreNegocio: "Test",
		TipoCliente:   "invalido",
		Telefono:      "0991234567",
		Direccion:     "Av. 1",
	})

	assert.EqualError(t, err, ErrTipoClienteInvalido.Error())
	mockRepo.AssertNotCalled(t, "CrearCliente")
}

func TestPedidoService_CrearCliente_TelefonoVacio(t *testing.T) {
	mockRepo := new(mockAlmacen)
	svc := NewPedidoService(mockRepo)

	_, err := svc.CrearCliente(models.Cliente{
		NombreNegocio: "Test",
		TipoCliente:   "restaurante",
		Direccion:     "Av. 1",
	})

	assert.EqualError(t, err, ErrTelefonoVacio.Error())
	mockRepo.AssertNotCalled(t, "CrearCliente")
}

func TestPedidoService_CrearDetalle_Valido(t *testing.T) {
	mockRepo := new(mockAlmacen)
	svc := NewPedidoService(mockRepo)

	entrada := models.DetallePedido{
		PedidoID:       1,
		EspecieID:      1,
		CantidadKg:     10,
		PrecioUnitario: 8.5,
	}
	esperado := entrada
	esperado.ID = 1
	esperado.Subtotal = 85.0

	mockRepo.On("CrearDetalle", mock.AnythingOfType("models.DetallePedido")).Return(esperado, nil)

	creado, err := svc.CrearDetalle(entrada)

	assert.NoError(t, err)
	assert.Equal(t, 85.0, creado.Subtotal)
	mockRepo.AssertExpectations(t)
}

func TestPedidoService_CrearDetalle_SinPedidoID(t *testing.T) {
	mockRepo := new(mockAlmacen)
	svc := NewPedidoService(mockRepo)

	_, err := svc.CrearDetalle(models.DetallePedido{
		EspecieID:      1,
		CantidadKg:     10,
		PrecioUnitario: 8.5,
	})

	assert.EqualError(t, err, ErrPedidoIDInvalido.Error())
	mockRepo.AssertNotCalled(t, "CrearDetalle")
}

func TestPedidoService_CrearDetalle_SinEspecieID(t *testing.T) {
	mockRepo := new(mockAlmacen)
	svc := NewPedidoService(mockRepo)

	_, err := svc.CrearDetalle(models.DetallePedido{
		PedidoID:       1,
		CantidadKg:     10,
		PrecioUnitario: 8.5,
	})

	assert.EqualError(t, err, ErrEspecieIDInvalido.Error())
	mockRepo.AssertNotCalled(t, "CrearDetalle")
}

func TestPedidoService_CrearDetalle_CantidadInvalida(t *testing.T) {
	mockRepo := new(mockAlmacen)
	svc := NewPedidoService(mockRepo)

	_, err := svc.CrearDetalle(models.DetallePedido{
		PedidoID:       1,
		EspecieID:      1,
		CantidadKg:     0,
		PrecioUnitario: 8.5,
	})

	assert.EqualError(t, err, ErrCantidadInvalida.Error())
	mockRepo.AssertNotCalled(t, "CrearDetalle")
}

func TestPedidoService_CrearDetalle_PrecioInvalido(t *testing.T) {
	mockRepo := new(mockAlmacen)
	svc := NewPedidoService(mockRepo)

	_, err := svc.CrearDetalle(models.DetallePedido{
		PedidoID:   1,
		EspecieID:  1,
		CantidadKg: 10,
	})

	assert.EqualError(t, err, ErrPrecioInvalido.Error())
	mockRepo.AssertNotCalled(t, "CrearDetalle")
}

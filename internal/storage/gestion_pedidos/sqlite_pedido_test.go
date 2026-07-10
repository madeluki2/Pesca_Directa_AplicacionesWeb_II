package gestion_pedidos

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

func abrirDBPedidosEnMemoria(t *testing.T) *gorm.DB {
	t.Helper()
	gdb, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = gdb.AutoMigrate(&models.Cliente{}, &models.Pedido{}, &models.DetallePedido{})
	require.NoError(t, err)
	return gdb
}

func TestRepositorioPedidos_CrearYListarPedido(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBPedidosEnMemoria(t))

	// Preparamos cliente necesario para el pedido
	cliente, _ := repo.CrearCliente(models.Cliente{NombreNegocio: "Restaurante El Faro"})

	pedido := models.Pedido{
		ClienteID: cliente.ID,
		Estado:    "pendiente",
		Total:     250.50,
	}

	creado, err := repo.CrearPedido(pedido)
	require.NoError(t, err)
	assert.NotZero(t, creado.ID)

	lista, err := repo.ListarPedidos()
	require.NoError(t, err)
	assert.Len(t, lista, 1)
	assert.Equal(t, "pendiente", lista[0].Estado)
}

func TestRepositorioPedidos_BuscarPedidoPorID(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBPedidosEnMemoria(t))

	cliente, _ := repo.CrearCliente(models.Cliente{NombreNegocio: "Cliente Test"})
	pedidoCreado, _ := repo.CrearPedido(models.Pedido{
		ClienteID: cliente.ID,
		Estado:    "en_proceso",
	})

	encontrado, err := repo.BuscarPedidoPorID(pedidoCreado.ID)

	require.NoError(t, err)
	assert.Equal(t, "en_proceso", encontrado.Estado)
	assert.NotNil(t, encontrado.Cliente)
	assert.Equal(t, "Cliente Test", encontrado.Cliente.NombreNegocio)
}

func TestRepositorioPedidos_CrearYBuscarCliente(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBPedidosEnMemoria(t))

	creado, err := repo.CrearCliente(models.Cliente{
		NombreNegocio: "Distribuidora El Puerto",
		TipoCliente:   "mayorista",
		Telefono:      "0997654321",
		Direccion:     "Calle 10 de Agosto",
		Estado:        "activo",
	})
	require.NoError(t, err)
	assert.NotZero(t, creado.ID)

	encontrado, err := repo.BuscarClientePorID(creado.ID)
	require.NoError(t, err)
	assert.Equal(t, "Distribuidora El Puerto", encontrado.NombreNegocio)
}

func TestRepositorioPedidos_ActualizarPedido(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBPedidosEnMemoria(t))

	cliente, _ := repo.CrearCliente(models.Cliente{NombreNegocio: "Test"})
	creado, _ := repo.CrearPedido(models.Pedido{
		ClienteID: cliente.ID,
		Estado:    "pendiente",
	})

	actualizado, err := repo.ActualizarPedido(creado.ID, models.Pedido{
		ClienteID: cliente.ID,
		Estado:    "entregado",
	})
	require.NoError(t, err)
	assert.Equal(t, "entregado", actualizado.Estado)
}

func TestRepositorioPedidos_EliminarPedido(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBPedidosEnMemoria(t))

	cliente, _ := repo.CrearCliente(models.Cliente{NombreNegocio: "Test"})
	creado, _ := repo.CrearPedido(models.Pedido{
		ClienteID: cliente.ID,
		Estado:    "pendiente",
	})

	err := repo.EliminarPedido(creado.ID)
	require.NoError(t, err)

	_, err = repo.BuscarPedidoPorID(creado.ID)
	assert.Error(t, err)
}

func TestRepositorioPedidos_CrearDetalle(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBPedidosEnMemoria(t))

	cliente, _ := repo.CrearCliente(models.Cliente{NombreNegocio: "Test"})
	pedido, _ := repo.CrearPedido(models.Pedido{ClienteID: cliente.ID, Estado: "pendiente"})

	detalle, err := repo.CrearDetalle(models.DetallePedido{
		PedidoID:       pedido.ID,
		EspecieID:      1,
		CantidadKg:     5,
		PrecioUnitario: 10.0,
		Subtotal:       50.0,
	})
	require.NoError(t, err)
	assert.NotZero(t, detalle.ID)
	assert.Equal(t, 50.0, detalle.Subtotal)
}

func TestRepositorioPedidos_ListarClientes(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBPedidosEnMemoria(t))

	repo.CrearCliente(models.Cliente{NombreNegocio: "Cliente 1", TipoCliente: "restaurante"})
	repo.CrearCliente(models.Cliente{NombreNegocio: "Cliente 2", TipoCliente: "mayorista"})

	lista, err := repo.ListarClientes()
	require.NoError(t, err)
	assert.Len(t, lista, 2)
}

func TestRepositorioPedidos_OperacionesPedidoPuro(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBPedidosEnMemoria(t))

	// 1. Preparación: Crear un cliente para asociar al pedido
	cliente, err := repo.CrearCliente(models.Cliente{
		NombreNegocio: "Restaurante Central",
		TipoCliente:   "restaurante",
	})
	require.NoError(t, err)

	// 2. Test: Crear Pedido
	pedido := models.Pedido{
		ClienteID: cliente.ID,
		Estado:    "pendiente",
		Total:     500.0,
	}
	creado, err := repo.CrearPedido(pedido)
	require.NoError(t, err)
	assert.NotZero(t, creado.ID)

	// 3. Test: Buscar Pedido
	encontrado, err := repo.BuscarPedidoPorID(creado.ID)
	require.NoError(t, err)
	assert.Equal(t, "pendiente", encontrado.Estado)
	assert.Equal(t, cliente.ID, encontrado.ClienteID)

	// 4. Test: Actualizar Pedido (cambio de estado y total)
	encontrado.Estado = "completado"
	encontrado.Total = 550.0
	actualizado, err := repo.ActualizarPedido(encontrado.ID, encontrado)
	require.NoError(t, err)
	assert.Equal(t, "completado", actualizado.Estado)
	assert.Equal(t, 550.0, actualizado.Total)

	// 5. Test: Listar Pedidos
	lista, err := repo.ListarPedidos()
	require.NoError(t, err)
	assert.Len(t, lista, 1)
}

func TestRepositorioPedidos_EliminarPedidoSolo(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBPedidosEnMemoria(t))

	// Preparación
	cliente, _ := repo.CrearCliente(models.Cliente{NombreNegocio: "Test Borrado"})
	pedido, _ := repo.CrearPedido(models.Pedido{ClienteID: cliente.ID, Estado: "pendiente"})

	// Ejecutar eliminación
	err := repo.EliminarPedido(pedido.ID)
	require.NoError(t, err)

	// Verificar
	_, err = repo.BuscarPedidoPorID(pedido.ID)
	assert.Error(t, err, "El pedido debería haber sido eliminado de la base de datos")
}

package storage

import "Pesca_Directa_AplicacionesWeb_II/internal/models"

type ClienteRepository interface {
	ListarClientes() []models.Cliente
	BuscarClientePorID(id int) (models.Cliente, bool)
	CrearCliente(c models.Cliente) models.Cliente
	ActualizarCliente(id int, datos models.Cliente) (models.Cliente, bool)
	EliminarCliente(id int) bool
	CambiarTipoCliente(id int, tipo string) (models.Cliente, bool)
}

type PedidoRepository interface {
	ListarPedidos() []models.Pedido
	BuscarPedidoPorID(id int) (models.Pedido, bool)
	CrearPedido(p models.Pedido) models.Pedido
	ActualizarPedido(id int, datos models.Pedido) (models.Pedido, bool)
	EliminarPedido(id int) bool
}

type DetallePedidoRepository interface {
	ListarDetalles() []models.DetallePedido
	BuscarDetallePorID(id int) (models.DetallePedido, bool)
	CrearDetalle(d models.DetallePedido) models.DetallePedido
	ActualizarDetalle(id int, datos models.DetallePedido) (models.DetallePedido, bool)
	EliminarDetalle(id int) bool
}

// Almacen une todos los repositorios del módulo de Gestión de Pedidos.
type Almacen interface {
	ClienteRepository
	PedidoRepository
	DetallePedidoRepository
	Seed() // Requerido para la compatibilidad con el Factory
}

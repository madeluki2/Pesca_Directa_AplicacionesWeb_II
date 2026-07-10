package gestion_pedidos

import "Pesca_Directa_AplicacionesWeb_II/internal/models"

type ClienteRepository interface {
	ListarClientes() ([]models.Cliente, error)
	BuscarClientePorID(id int) (models.Cliente, error)
	CrearCliente(c models.Cliente) (models.Cliente, error)
	ActualizarCliente(id int, datos models.Cliente) (models.Cliente, error)
	EliminarCliente(id int) error
	CambiarTipoCliente(id int, tipo string) (models.Cliente, error)
}

type PedidoRepository interface {
	ListarPedidos() ([]models.Pedido, error)
	BuscarPedidoPorID(id int) (models.Pedido, error)
	CrearPedido(p models.Pedido) (models.Pedido, error)
	ActualizarPedido(id int, datos models.Pedido) (models.Pedido, error)
	EliminarPedido(id int) error
}

type DetallePedidoRepository interface {
	ListarDetalles() ([]models.DetallePedido, error)
	BuscarDetallePorID(id int) (models.DetallePedido, error)
	CrearDetalle(d models.DetallePedido) (models.DetallePedido, error)
	ActualizarDetalle(id int, datos models.DetallePedido) (models.DetallePedido, error)
	EliminarDetalle(id int) error
}

// Almacen une todos los repositorios del módulo de Gestión de Pedidos.
type Almacen interface {
	ClienteRepository
	PedidoRepository
	DetallePedidoRepository
	Seed() // Requerido para la compatibilidad con el Factory y carga automática de seeders (Gate G2)
}

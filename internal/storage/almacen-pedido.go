package storage

import "Pesca_Directa_AplicacionesWeb_II/internal/models"

// UsuarioRepository define las operaciones de autenticación.
// Va separado de AlmacenPedido porque el usuario es transversal a todos los módulos.
type UsuarioRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

// ClienteRepository define las operaciones del repositorio de clientes.
type ClienteRepository interface {
	ListarClientes() []models.Cliente
	BuscarClientePorID(id int) (models.Cliente, bool)
	CrearCliente(c models.Cliente) models.Cliente
	ActualizarCliente(id int, datos models.Cliente) (models.Cliente, bool)
	EliminarCliente(id int) bool
	CambiarTipoCliente(id int, nuevoTipo string) (models.Cliente, bool)
}

// PedidoRepository define las operaciones del repositorio de pedidos.
type PedidoRepository interface {
	ListarPedidos() []models.Pedido
	BuscarPedidoPorID(id int) (models.Pedido, bool)
	CrearPedido(p models.Pedido) models.Pedido
	ActualizarPedido(id int, datos models.Pedido) (models.Pedido, bool)
	EliminarPedido(id int) bool
}

// DetalleRepository define las operaciones del repositorio de detalles de pedido.
type DetalleRepository interface {
	ListarDetalles() []models.DetallePedido
	BuscarDetallePorID(id int) (models.DetallePedido, bool)
	CrearDetalle(d models.DetallePedido) models.DetallePedido
	ActualizarDetalle(id int, datos models.DetallePedido) (models.DetallePedido, bool)
	EliminarDetalle(id int) bool
}

// Almacen une todos los repositorios del módulo de Gestión de Pedidos.
// La implementación puede ser en memoria o SQLite sin tocar ningún handler.
type Almacen interface {
	ClienteRepository
	PedidoRepository
	DetalleRepository
}

// Chequeo en tiempo de compilación: si Memoria dejara de cumplir Almacen,
// el proyecto NO compila.
var _ Almacen = (*Memoria)(nil)

// Chequeo en tiempo de compilación: si AlmacenSQLite dejara de cumplir Almacen,
// el proyecto NO compila.
var _ Almacen = (*AlmacenSQLite)(nil)

package storage

import "Pesca_Directa_AplicacionesWeb_II/internal/models"

// Almacen define QUÉ sabe hacer el almacén del módulo de Gestión de Pedidos,
// sin decir CÓMO. Memoria (slices) ya cumple esta interfaz por duck typing.
// Si en el futuro se usa SQLite, solo hay que cumplir esta misma interfaz
// sin tocar ningún handler.
type Almacen interface {
	// Clientes
	ListarClientes() []models.Cliente
	BuscarClientePorID(id int) (models.Cliente, bool)
	CrearCliente(c models.Cliente) models.Cliente
	ActualizarCliente(id int, datos models.Cliente) (models.Cliente, bool)
	EliminarCliente(id int) bool
	CambiarTipoCliente(id int, nuevoTipo string) (models.Cliente, bool)

	// Pedidos
	ListarPedidos() []models.Pedido
	BuscarPedidoPorID(id int) (models.Pedido, bool)
	CrearPedido(p models.Pedido) models.Pedido
	ActualizarPedido(id int, datos models.Pedido) (models.Pedido, bool)
	EliminarPedido(id int) bool

	// Detalles de Pedido
	ListarDetalles() []models.DetallePedido
	BuscarDetallePorID(id int) (models.DetallePedido, bool)
	CrearDetalle(d models.DetallePedido) models.DetallePedido
	ActualizarDetalle(id int, datos models.DetallePedido) (models.DetallePedido, bool)
	EliminarDetalle(id int) bool
}

// Chequeo en tiempo de compilación: si Memoria dejara de cumplir Almacen,
// el proyecto NO compila. Red de seguridad.
var _ Almacen = (*Memoria)(nil)

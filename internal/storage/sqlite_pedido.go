package storage

import (
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// AlmacenSQLite implementa Almacen usando GORM + SQLite.
type AlmacenSQLite struct {
	db *gorm.DB
}

// NuevoAlmacenSQLite crea un almacén SQLite con GORM inyectado.
func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

// =========================================================
// CLIENTES
// =========================================================

// ListarClientes devuelve todos los clientes de la base de datos.
func (a *AlmacenSQLite) ListarClientes() []models.Cliente {
	var clientes []models.Cliente
	a.db.Find(&clientes)
	return clientes
}

// BuscarClientePorID devuelve el cliente con el ID dado (patrón comma-ok).
func (a *AlmacenSQLite) BuscarClientePorID(id int) (models.Cliente, bool) {
	var c models.Cliente
	resultado := a.db.First(&c, id)
	if resultado.Error != nil {
		return models.Cliente{}, false
	}
	return c, true
}

// CrearCliente inserta un nuevo cliente en la base de datos.
func (a *AlmacenSQLite) CrearCliente(c models.Cliente) models.Cliente {
	a.db.Create(&c)
	return c
}

// ActualizarCliente reemplaza los datos del cliente con el ID dado.
func (a *AlmacenSQLite) ActualizarCliente(id int, datos models.Cliente) (models.Cliente, bool) {
	var c models.Cliente
	if a.db.First(&c, id).Error != nil {
		return models.Cliente{}, false
	}
	if datos.TipoCliente != "" {
		c.TipoCliente = datos.TipoCliente
	}
	if datos.NombreNegocio != "" {
		c.NombreNegocio = datos.NombreNegocio
	}
	if datos.Direccion != "" {
		c.Direccion = datos.Direccion
	}
	if datos.Telefono != "" {
		c.Telefono = datos.Telefono
	}
	if datos.Estado != "" {
		c.Estado = datos.Estado
	}
	if datos.UsuarioID != 0 {
		c.UsuarioID = datos.UsuarioID
	}
	a.db.Save(&c)
	return c, true
}

// EliminarCliente elimina el cliente con el ID dado.
func (a *AlmacenSQLite) EliminarCliente(id int) bool {
	resultado := a.db.Delete(&models.Cliente{}, id)
	return resultado.RowsAffected > 0
}

// CambiarTipoCliente actualiza únicamente el tipo de un cliente existente.
func (a *AlmacenSQLite) CambiarTipoCliente(id int, nuevoTipo string) (models.Cliente, bool) {
	var c models.Cliente
	if a.db.First(&c, id).Error != nil {
		return models.Cliente{}, false
	}
	c.TipoCliente = nuevoTipo
	a.db.Save(&c)
	return c, true
}

// =========================================================
// PEDIDOS
// =========================================================

// ListarPedidos devuelve todos los pedidos de la base de datos.
func (a *AlmacenSQLite) ListarPedidos() []models.Pedido {
	var pedidos []models.Pedido
	a.db.Find(&pedidos)
	return pedidos
}

// BuscarPedidoPorID devuelve el pedido con el ID dado (patrón comma-ok).
func (a *AlmacenSQLite) BuscarPedidoPorID(id int) (models.Pedido, bool) {
	var p models.Pedido
	resultado := a.db.First(&p, id)
	if resultado.Error != nil {
		return models.Pedido{}, false
	}
	return p, true
}

// CrearPedido inserta un nuevo pedido en la base de datos.
func (a *AlmacenSQLite) CrearPedido(p models.Pedido) models.Pedido {
	a.db.Create(&p)
	return p
}

// ActualizarPedido reemplaza los datos del pedido con el ID dado.
func (a *AlmacenSQLite) ActualizarPedido(id int, datos models.Pedido) (models.Pedido, bool) {
	var p models.Pedido
	if a.db.First(&p, id).Error != nil {
		return models.Pedido{}, false
	}
	if datos.Estado != "" {
		p.Estado = datos.Estado
	}
	if datos.Fecha != "" {
		p.Fecha = datos.Fecha
	}
	if datos.Total != 0 {
		p.Total = datos.Total
	}
	if datos.ClienteID != 0 {
		p.ClienteID = datos.ClienteID
	}
	a.db.Save(&p)
	return p, true
}

// EliminarPedido elimina el pedido con el ID dado.
func (a *AlmacenSQLite) EliminarPedido(id int) bool {
	resultado := a.db.Delete(&models.Pedido{}, id)
	return resultado.RowsAffected > 0
}

// =========================================================
// DETALLES DE PEDIDO
// =========================================================

// ListarDetalles devuelve todos los detalles de la base de datos.
func (a *AlmacenSQLite) ListarDetalles() []models.DetallePedido {
	var detalles []models.DetallePedido
	a.db.Find(&detalles)
	return detalles
}

// BuscarDetallePorID devuelve el detalle con el ID dado (patrón comma-ok).
func (a *AlmacenSQLite) BuscarDetallePorID(id int) (models.DetallePedido, bool) {
	var d models.DetallePedido
	resultado := a.db.First(&d, id)
	if resultado.Error != nil {
		return models.DetallePedido{}, false
	}
	return d, true
}

// CrearDetalle inserta un nuevo detalle en la base de datos.
func (a *AlmacenSQLite) CrearDetalle(d models.DetallePedido) models.DetallePedido {
	a.db.Create(&d)
	return d
}

// ActualizarDetalle reemplaza los datos del detalle con el ID dado.
func (a *AlmacenSQLite) ActualizarDetalle(id int, datos models.DetallePedido) (models.DetallePedido, bool) {
	var d models.DetallePedido
	if a.db.First(&d, id).Error != nil {
		return models.DetallePedido{}, false
	}
	if datos.CantidadKg != 0 {
		d.CantidadKg = datos.CantidadKg
	}
	if datos.PrecioUnitario != 0 {
		d.PrecioUnitario = datos.PrecioUnitario
	}
	if datos.Subtotal != 0 {
		d.Subtotal = datos.Subtotal
	}
	if datos.EspecieID != 0 {
		d.EspecieID = datos.EspecieID
	}
	if datos.PedidoID != 0 {
		d.PedidoID = datos.PedidoID
	}
	a.db.Save(&d)
	return d, true
}

// EliminarDetalle elimina el detalle con el ID dado.
func (a *AlmacenSQLite) EliminarDetalle(id int) bool {
	resultado := a.db.Delete(&models.DetallePedido{}, id)
	return resultado.RowsAffected > 0
}

// Chequeo en tiempo de compilación: si AlmacenSQLite dejara de cumplir Almacen,
// el proyecto NO compila.
var _ Almacen = (*AlmacenSQLite)(nil)

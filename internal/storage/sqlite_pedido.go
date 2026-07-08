package storage

import (
	"Pesca_Directa_AplicacionesWeb_II/internal/models"

	"gorm.io/gorm"
)

type AlmacenSQLite struct {
	db *gorm.DB
}

func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

func (a *AlmacenSQLite) Seed() {} // No realiza operaciones en BD real

// ═══════════════════════════ CLIENTES ══════════════════════════════════════

func (a *AlmacenSQLite) ListarClientes() []models.Cliente {
	var lista []models.Cliente
	a.db.Find(&lista)
	return lista
}

func (a *AlmacenSQLite) BuscarClientePorID(id int) (models.Cliente, bool) {
	var c models.Cliente
	if err := a.db.First(&c, id).Error; err != nil {
		return models.Cliente{}, false
	}
	return c, true
}

func (a *AlmacenSQLite) CrearCliente(c models.Cliente) models.Cliente {
	a.db.Create(&c)
	return c
}

func (a *AlmacenSQLite) ActualizarCliente(id int, datos models.Cliente) (models.Cliente, bool) {
	var existente models.Cliente
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Cliente{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) EliminarCliente(id int) bool {
	res := a.db.Delete(&models.Cliente{}, id)
	return res.RowsAffected > 0
}

func (a *AlmacenSQLite) CambiarTipoCliente(id int, tipo string) (models.Cliente, bool) {
	var c models.Cliente
	if err := a.db.First(&c, id).Error; err != nil {
		return models.Cliente{}, false
	}
	a.db.Model(&c).Update("tipo_cliente", tipo)
	return c, true
}

// ═══════════════════════════ PEDIDOS ═══════════════════════════════════════

func (a *AlmacenSQLite) ListarPedidos() []models.Pedido {
	var lista []models.Pedido
	a.db.Find(&lista)
	return lista
}

func (a *AlmacenSQLite) BuscarPedidoPorID(id int) (models.Pedido, bool) {
	var p models.Pedido
	if err := a.db.First(&p, id).Error; err != nil {
		return models.Pedido{}, false
	}
	return p, true
}

func (a *AlmacenSQLite) CrearPedido(p models.Pedido) models.Pedido {
	a.db.Create(&p)
	return p
}

func (a *AlmacenSQLite) ActualizarPedido(id int, datos models.Pedido) (models.Pedido, bool) {
	var existente models.Pedido
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Pedido{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) EliminarPedido(id int) bool {
	res := a.db.Delete(&models.Pedido{}, id)
	return res.RowsAffected > 0
}

// ═══════════════════════════ DETALLES PEDIDO ══════════════════════════════

func (a *AlmacenSQLite) ListarDetalles() []models.DetallePedido {
	var lista []models.DetallePedido
	a.db.Find(&lista)
	return lista
}

func (a *AlmacenSQLite) BuscarDetallePorID(id int) (models.DetallePedido, bool) {
	var d models.DetallePedido
	if err := a.db.First(&d, id).Error; err != nil {
		return models.DetallePedido{}, false
	}
	return d, true
}

func (a *AlmacenSQLite) CrearDetalle(d models.DetallePedido) models.DetallePedido {
	a.db.Create(&d)
	return d
}

func (a *AlmacenSQLite) ActualizarDetalle(id int, datos models.DetallePedido) (models.DetallePedido, bool) {
	var existente models.DetallePedido
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.DetallePedido{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) EliminarDetalle(id int) bool {
	res := a.db.Delete(&models.DetallePedido{}, id)
	return res.RowsAffected > 0
}

var _ Almacen = (*AlmacenSQLite)(nil)

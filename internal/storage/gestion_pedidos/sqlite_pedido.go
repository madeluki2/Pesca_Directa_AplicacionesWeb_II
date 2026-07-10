package gestion_pedidos

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

func (a *AlmacenSQLite) Seed() {}

// ═══════════════════════════ CLIENTES ══════════════════════════════════════

func (a *AlmacenSQLite) ListarClientes() ([]models.Cliente, error) {
	var lista []models.Cliente
	if err := a.db.Find(&lista).Error; err != nil {
		return nil, err
	}
	return lista, nil
}

func (a *AlmacenSQLite) BuscarClientePorID(id int) (models.Cliente, error) {
	var c models.Cliente
	if err := a.db.First(&c, id).Error; err != nil {
		return models.Cliente{}, err
	}
	return c, nil
}

func (a *AlmacenSQLite) CrearCliente(c models.Cliente) (models.Cliente, error) {
	if err := a.db.Create(&c).Error; err != nil {
		return models.Cliente{}, err
	}
	return c, nil
}

func (a *AlmacenSQLite) ActualizarCliente(id int, datos models.Cliente) (models.Cliente, error) {
	var existente models.Cliente
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Cliente{}, err
	}
	datos.ID = id
	if err := a.db.Save(&datos).Error; err != nil {
		return models.Cliente{}, err
	}
	return datos, nil
}

func (a *AlmacenSQLite) EliminarCliente(id int) error {
	res := a.db.Delete(&models.Cliente{}, id)
	return res.Error
}

func (a *AlmacenSQLite) CambiarTipoCliente(id int, tipo string) (models.Cliente, error) {
	var c models.Cliente
	if err := a.db.First(&c, id).Error; err != nil {
		return models.Cliente{}, err
	}
	if err := a.db.Model(&c).Update("tipo_cliente", tipo).Error; err != nil {
		return models.Cliente{}, err
	}
	return c, nil
}

// ═══════════════════════════ PEDIDOS ═══════════════════════════════════════

func (a *AlmacenSQLite) ListarPedidos() ([]models.Pedido, error) {
	var lista []models.Pedido
	if err := a.db.Preload("Detalles").Preload("Cliente").Find(&lista).Error; err != nil {
		return nil, err
	}
	return lista, nil
}

func (a *AlmacenSQLite) BuscarPedidoPorID(id int) (models.Pedido, error) {
	var p models.Pedido
	if err := a.db.Preload("Detalles").Preload("Cliente").First(&p, id).Error; err != nil {
		return models.Pedido{}, err
	}
	return p, nil
}

func (a *AlmacenSQLite) CrearPedido(p models.Pedido) (models.Pedido, error) {
	if err := a.db.Create(&p).Error; err != nil {
		return models.Pedido{}, err
	}
	return p, nil
}

func (a *AlmacenSQLite) ActualizarPedido(id int, datos models.Pedido) (models.Pedido, error) {
	var existente models.Pedido
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Pedido{}, err
	}
	datos.ID = id
	if err := a.db.Save(&datos).Error; err != nil {
		return models.Pedido{}, err
	}
	return datos, nil
}

func (a *AlmacenSQLite) EliminarPedido(id int) error {
	res := a.db.Delete(&models.Pedido{}, id)
	return res.Error
}

// ═══════════════════════════ DETALLES PEDIDO ══════════════════════════════

func (a *AlmacenSQLite) ListarDetalles() ([]models.DetallePedido, error) {
	var lista []models.DetallePedido
	if err := a.db.Find(&lista).Error; err != nil {
		return nil, err
	}
	return lista, nil
}

func (a *AlmacenSQLite) BuscarDetallePorID(id int) (models.DetallePedido, error) {
	var d models.DetallePedido
	if err := a.db.First(&d, id).Error; err != nil {
		return models.DetallePedido{}, err
	}
	return d, nil
}

func (a *AlmacenSQLite) CrearDetalle(d models.DetallePedido) (models.DetallePedido, error) {
	if err := a.db.Create(&d).Error; err != nil {
		return models.DetallePedido{}, err
	}
	return d, nil
}

func (a *AlmacenSQLite) ActualizarDetalle(id int, datos models.DetallePedido) (models.DetallePedido, error) {
	var existente models.DetallePedido
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.DetallePedido{}, err
	}
	datos.ID = id
	if err := a.db.Save(&datos).Error; err != nil {
		return models.DetallePedido{}, err
	}
	return datos, nil
}

func (a *AlmacenSQLite) EliminarDetalle(id int) error {
	res := a.db.Delete(&models.DetallePedido{}, id)
	return res.Error
}

// --- SOLUCIÓN: Eliminamos el prefijo "storage." porque todo está en el mismo paquete ---
var _ Almacen = (*AlmacenSQLite)(nil)

package models

import "time"

// Cliente representa a un restaurante, intermediario o mayorista que realiza pedidos.
type Cliente struct {
	ID            int      `json:"id"             gorm:"primaryKey;autoIncrement"`
	UsuarioID     int      `json:"usuario_id"     gorm:"not null"`
	TipoCliente   string   `json:"tipo_cliente"   gorm:"not null"` // restaurante / intermediario / mayorista
	NombreNegocio string   `json:"nombre_negocio" gorm:"not null"`
	Direccion     string   `json:"direccion"      gorm:"not null"`
	Telefono      string   `json:"telefono"       gorm:"not null"`
	Estado        string   `json:"estado"         gorm:"default:activo"`
	Pedidos       []Pedido `json:"pedidos,omitempty" gorm:"foreignKey:ClienteID;constraint:OnDelete:CASCADE;"` // Relación Has-Many (Un cliente tiene muchos pedidos)
}

// Pedido representa una solicitud de compra realizada por un cliente.
type Pedido struct {
	ID        int             `json:"id"         gorm:"primaryKey;autoIncrement"`
	ClienteID int             `json:"cliente_id" gorm:"not null"`
	Cliente   *Cliente        `json:"cliente,omitempty" gorm:"foreignKey:ClienteID"` // Relación Belongs-To (Un pedido pertenece a un Cliente)
	Fecha     time.Time       `json:"fecha"      gorm:"not null;type:datetime"`      // Mejorado a time.Time para que GORM maneje fechas reales
	Estado    string          `json:"estado"     gorm:"default:pendiente"`           // pendiente / en_proceso / entregado / cancelado
	Total     float64         `json:"total"      gorm:"default:0"`
	Detalles  []DetallePedido `json:"detalles"   gorm:"foreignKey:PedidoID;constraint:OnDelete:CASCADE;"` // Relación Has-Many (Un pedido tiene muchos detalles)
}

// DetallePedido representa cada producto (especie) dentro de un pedido.
type DetallePedido struct {
	ID             int     `json:"id"             gorm:"primaryKey;autoIncrement"`
	PedidoID       int     `json:"pedido_id"      gorm:"not null"`
	EspecieID      int     `json:"especie_id"      gorm:"not null"` // referencia externa al módulo de Pesca
	CantidadKg     float64 `json:"cantidad_kg"     gorm:"not null"`
	PrecioUnitario float64 `json:"precio_unitario" gorm:"not null"`
	Subtotal       float64 `json:"subtotal"        gorm:"not null"`
}

package models

// Cliente representa a un restaurante, intermediario o mayorista que realiza pedidos
type Cliente struct {
	ID            int    `json:"id"`
	UsuarioID     int    `json:"usuario_id"`
	TipoCliente   string `json:"tipo_cliente"` // restaurante / intermediario / mayorista
	NombreNegocio string `json:"nombre_negocio"`
	Direccion     string `json:"direccion"`
	Telefono      string `json:"telefono"`
	Estado        string `json:"estado"`
}

// Pedido representa una solicitud de compra realizada por un cliente
type Pedido struct {
	ID        int     `json:"id"`
	ClienteID int     `json:"cliente_id"`
	Fecha     string  `json:"fecha"`
	Estado    string  `json:"estado"` // pendiente / en_proceso / entregado / cancelado
	Total     float64 `json:"total"`
}

// DetallePedido representa cada producto (especie) dentro de un pedido
type DetallePedido struct {
	ID             int     `json:"id"`
	PedidoID       int     `json:"pedido_id"`
	EspecieID      int     `json:"especie_id"` // referencia externa al módulo de Pesca
	CantidadKg     float64 `json:"cantidad_kg"`
	PrecioUnitario float64 `json:"precio_unitario"`
	Subtotal       float64 `json:"subtotal"`
}

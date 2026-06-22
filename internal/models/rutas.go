package models

import "time"

// ═══════════════════════════════════════════════════════════════
// MÓDULO: RUTA DE DISTRIBUCIÓN
// Entidades: Ruta, Punto, Transportista, EntregaPedido
// ═══════════════════════════════════════════════════════════════

// Ruta define un trayecto de distribución con origen y destino.
type Ruta struct {
	ID          uint      `json:"id"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Origen      string    `json:"origen"`
	Destino     string    `json:"destino"`
	Estado      string    `json:"estado"` // activo / inactivo
	CreadoEn    time.Time `json:"creado_en"`
}

// Punto es una parada geolocalizada dentro de una Ruta.
type Punto struct {
	ID        uint      `json:"id"`
	RutaID    uint      `json:"ruta_id"` // FK → Ruta
	Nombre    string    `json:"nombre"`
	Direccion string    `json:"direccion"`
	Latitud   float64   `json:"latitud"`
	Longitud  float64   `json:"longitud"`
	OrdenRuta int       `json:"orden_ruta"`
	Estado    string    `json:"estado"` // activo / inactivo
	CreadoEn  time.Time `json:"creado_en"`
}

// Transportista representa al conductor asignado a una entrega.
type Transportista struct {
	ID            uint      `json:"id"`
	Nombre        string    `json:"nombre"`
	Telefono      string    `json:"telefono"`
	PlacaVehiculo string    `json:"placa_vehiculo"` // único
	Estado        string    `json:"estado"`         // activo / inactivo
	CreadoEn      time.Time `json:"creado_en"`
}

// EntregaPedido vincula un pedido con un punto de entrega y un transportista.
type EntregaPedido struct {
	ID                   uint      `json:"id"`
	PedidoID             uint      `json:"pedido_id"`        // FK → Pedido
	PuntoID              uint      `json:"punto_id"`         // FK → Punto
	TransportistaID      uint      `json:"transportista_id"` // FK → Transportista
	FechaSalida          time.Time `json:"fecha_salida"`
	FechaEntregaEstimada time.Time `json:"fecha_entrega_estimada"`
	FechaEntregaReal     time.Time `json:"fecha_entrega_real"`
	Estado               string    `json:"estado"` // pendiente / en_ruta / entregado / retrasado / cancelado
	Observacion          string    `json:"observacion"`
	CreadoEn             time.Time `json:"creado_en"`
}

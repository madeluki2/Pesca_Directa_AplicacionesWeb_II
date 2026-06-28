package models

import "time"

// ═══════════════════════════════════════════════════════════════
// MÓDULO: RUTA DE DISTRIBUCIÓN
// Entidades: Ruta, Punto, Transportista, EntregaPedido
// ═══════════════════════════════════════════════════════════════

// Ruta define un trayecto de distribución con origen y destino.
type Ruta struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Nombre      string    `json:"nombre" gorm:"not null"`
	Descripcion string    `json:"descripcion"`
	Origen      string    `json:"origen" gorm:"not null"`
	Destino     string    `json:"destino" gorm:"not null"`
	Estado      string    `json:"estado" gorm:"default:activo"` // activo / inactivo
	CreadoEn    time.Time `json:"creado_en" gorm:"autoCreateTime"`
}

// Punto es una parada geolocalizada dentro de una Ruta.
type Punto struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	RutaID    uint      `json:"ruta_id" gorm:"not null;index"` // FK → Ruta
	Nombre    string    `json:"nombre" gorm:"not null"`
	Direccion string    `json:"direccion" gorm:"not null"`
	Latitud   float64   `json:"latitud"`
	Longitud  float64   `json:"longitud"`
	OrdenRuta int       `json:"orden_ruta"`
	Estado    string    `json:"estado" gorm:"default:activo"` // activo / inactivo
	CreadoEn  time.Time `json:"creado_en" gorm:"autoCreateTime"`
}

// Transportista representa al conductor asignado a una entrega.
type Transportista struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Nombre        string    `json:"nombre" gorm:"not null"`
	Telefono      string    `json:"telefono" gorm:"not null"`
	PlacaVehiculo string    `json:"placa_vehiculo" gorm:"uniqueIndex;not null"` // único
	Estado        string    `json:"estado" gorm:"default:activo"`               // activo / inactivo
	CreadoEn      time.Time `json:"creado_en" gorm:"autoCreateTime"`
}

// EntregaPedido vincula un pedido con un punto de entrega y un transportista.
type EntregaPedido struct {
	ID                   uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	PedidoID             uint      `json:"pedido_id" gorm:"not null;index"`        // FK → Pedido
	PuntoID              uint      `json:"punto_id" gorm:"not null;index"`         // FK → Punto
	TransportistaID      uint      `json:"transportista_id" gorm:"not null;index"` // FK → Transportista
	FechaSalida          time.Time `json:"fecha_salida"`
	FechaEntregaEstimada time.Time `json:"fecha_entrega_estimada"`
	FechaEntregaReal     time.Time `json:"fecha_entrega_real"`
	Estado               string    `json:"estado" gorm:"default:pendiente"` // pendiente / en_ruta / entregado / retrasado / cancelado
	Observacion          string    `json:"observacion"`
	CreadoEn             time.Time `json:"creado_en" gorm:"autoCreateTime"`
}

package models

type Pescador struct {
	ID              int    `json:"id"               gorm:"primaryKey"`
	UsuarioID       int    `json:"usuario_id"       gorm:"not null"`
	Cedula          string `json:"cedula"           gorm:"not null"`
	ExperienciaAños int    `json:"experiencia_años"`
	PuertoBase      string `json:"puerto_base"      gorm:"not null"`
	Estado          bool   `json:"estado"           gorm:"default:true"`
}

type Embarcacion struct {
	ID          int     `json:"id"           gorm:"primaryKey"`
	PescadorID  int     `json:"pescador_id"  gorm:"not null"`
	Nombre      string  `json:"nombre"       gorm:"not null"`
	Matricula   string  `json:"matricula"    gorm:"not null"`
	CapacidadKG float64 `json:"capacidad_kg"`
	Estado      bool    `json:"estado"       gorm:"default:true"`
}

type Especie struct {
	ID               int    `json:"id"                gorm:"primaryKey"`
	NombreComun      string `json:"nombre_comun"      gorm:"not null"`
	NombreCientifico string `json:"nombre_cientifico"`
	UnidadMedida     string `json:"unidad_medida"     gorm:"not null"`
	Temporada        string `json:"temporada"`
	Estado           bool   `json:"estado"            gorm:"default:true"`
}

type Captura struct {
	ID             int     `json:"id"              gorm:"primaryKey"`
	EmbarcacionID  int     `json:"embarcacion_id"  gorm:"not null"`
	EspecieID      int     `json:"especie_id"      gorm:"not null"`
	Fecha          string  `json:"fecha"           gorm:"not null"`
	CantidadKG     float64 `json:"cantidad_kg"     gorm:"not null"`
	PrecioSugerido float64 `json:"precio_sugerido"`
	EstadoFrescura string  `json:"estado_frescura" gorm:"not null"`
}

type Bodega struct {
	ID          int     `json:"id"           gorm:"primaryKey"`
	Nombre      string  `json:"nombre"       gorm:"not null"`
	Ubicacion   string  `json:"ubicacion"    gorm:"not null"`
	CapacidadKG float64 `json:"capacidad_kg"`
	Estado      bool    `json:"estado"       gorm:"default:true"`
}

type Stock struct {
	ID           int     `json:"id"            gorm:"primaryKey"`
	BodegaID     int     `json:"bodega_id"     gorm:"not null"`
	EspecieID    int     `json:"especie_id"    gorm:"not null"`
	CantidadKG   float64 `json:"cantidad_kg"   gorm:"not null"`
	FechaIngreso string  `json:"fecha_ingreso" gorm:"not null"`
	Estado       bool    `json:"estado"        gorm:"default:true"`
}

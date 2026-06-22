package models

type Usuario struct {
	Id           int    `json:"id"`
	Nombre       string `json:"nombre"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Telefono     string `json:"telefono"`
	TipoUsuario  string `json:"tipo_usuario"`
	Estado       bool   `json:"estado"`
}

type Pescador struct {
	ID              int    `json:"id"`
	UsuarioID       int    `json:"usuario_id"`
	Cedula          string `json:"cedula"`
	ExperienciaAños int    `json:"experiencia_años"`
	PuertoBase      string `json:"puerto_base"`
	Estado          bool   `json:"estado"`
}

type Embarcacion struct {
	ID          int     `json:"id"`
	PescadorID  int     `json:"pescador_id"`
	Nombre      string  `json:"nombre"`
	Matricula   string  `json:"matricula"`
	CapacidadKG float64 `json:"capacidad_kg"`
	Estado      bool    `json:"estado"`
}

type Especie struct {
	ID               int    `json:"id"`
	NombreComun      string `json:"nombre_comun"`
	NombreCientifico string `json:"nombre_cientifico"`
	UnidadMedida     string `json:"unidad_medida"`
	Temporada        string `json:"temporada"`
	Estado           bool   `json:"estado"`
}

type Captura struct {
	ID             int     `json:"id"`
	EmbarcacionID  int     `json:"embarcacion_id"`
	EspecieID      int     `json:"especie_id"`
	Fecha          string  `json:"fecha"`
	CantidadKG     float64 `json:"cantidad_kg"`
	PrecioSugerido float64 `json:"precio_sugerido"`
	EstadoFrescura string  `json:"estado_frescura"`
}

type Bodega struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Ubicacion   string  `json:"ubicacion"`
	CapacidadKG float64 `json:"capacidad_kg"`
	Estado      bool    `json:"estado"`
}

type Stock struct {
	ID           int     `json:"id"`
	BodegaID     int     `json:"bodega_id"`
	EspecieID    int     `json:"especie_id"`
	CantidadKG   float64 `json:"cantidad_kg"`
	FechaIngreso string  `json:"fecha_ingreso"`
	Estado       bool    `json:"estado"`
}

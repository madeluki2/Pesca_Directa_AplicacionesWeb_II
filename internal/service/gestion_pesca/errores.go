package gestion_pesca

import "errors"

var (
	ErrCredencialesInvalidas   = errors.New("email o contraseña incorrectos")
	ErrPescadorNoEncontrado    = errors.New("pescador no encontrado")
	ErrEmbarcacionNoEncontrada = errors.New("embarcacion no encontrada")
	ErrEspecieNoEncontrada     = errors.New("especie no encontrada")
	ErrCapturaNoEncontrada     = errors.New("captura no encontrada")
	ErrBodegaNoEncontrada      = errors.New("bodega no encontrada")
	ErrStockNoEncontrado       = errors.New("stock no encontrado")
	ErrCedulaVacia             = errors.New("cedula es requerida")
	ErrPuertoVacio             = errors.New("puerto_base es requerido")
	ErrNombreVacio             = errors.New("nombre es requerido")
	ErrMatriculaVacia          = errors.New("matricula es requerida")
	ErrCantidadInvalida        = errors.New("cantidad_kg debe ser mayor a 0")
	ErrFechaVacia              = errors.New("fecha es requerida")
	ErrNombreComunVacio        = errors.New("nombre_comun es requerido")
	ErrUnidadMedidaVacia       = errors.New("unidad_medida es requerida")
	ErrFrescuraInvalida        = errors.New("estado_frescura debe ser 'fresco', 'refrigerado' o 'congelado'")
	ErrUbicacionVacia          = errors.New("ubicacion es requerida")
	ErrCapacidadInvalida       = errors.New("capacidad_kg debe ser mayor a 0")
	ErrCantidadNegativa        = errors.New("cantidad_kg debe ser mayor a 0")
	ErrFechaIngresoVacia       = errors.New("fecha_ingreso es requerida")
)

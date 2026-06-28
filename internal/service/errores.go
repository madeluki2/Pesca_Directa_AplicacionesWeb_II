package service

import "errors"

// Errores posibles del dominio Pesca-Directa Tarqui :v
var (
	// Auth
	ErrCredencialesInvalidas = errors.New("email o contraseña incorrectos")
	ErrEmailEnUso            = errors.New("email ya registrado")

	// Genérico
	ErrNoEncontrado = errors.New("recurso no encontrado")

	// Pescadores
	ErrPescadorNoEncontrado = errors.New("pescador no encontrado")
	ErrCedulaVacia          = errors.New("cedula es requerida")
	ErrPuertoVacio          = errors.New("puerto_base es requerido")

	// Embarcaciones
	ErrEmbarcacionNoEncontrada = errors.New("embarcacion no encontrada")
	ErrNombreVacio             = errors.New("nombre es requerido")
	ErrMatriculaVacia          = errors.New("matricula es requerida")

	// Especies
	ErrEspecieNoEncontrada = errors.New("especie no encontrada")
	ErrNombreComunVacio    = errors.New("nombre_comun es requerido")
	ErrUnidadMedidaVacia   = errors.New("unidad_medida es requerida")

	// Capturas
	ErrCapturaNoEncontrada = errors.New("captura no encontrada")
	ErrCantidadInvalida    = errors.New("cantidad_kg debe ser mayor a 0")
	ErrFechaVacia          = errors.New("fecha es requerida")
	ErrFrescuraInvalida    = errors.New("estado_frescura debe ser 'fresco', 'refrigerado' o 'congelado'")

	// Bodegas
	ErrBodegaNoEncontrada = errors.New("bodega no encontrada")
	ErrUbicacionVacia     = errors.New("ubicacion es requerida")
	ErrCapacidadInvalida  = errors.New("capacidad_kg debe ser mayor a 0")

	// Stocks
	ErrStockNoEncontrado = errors.New("stock no encontrado")
	ErrFechaIngresoVacia = errors.New("fecha_ingreso es requerida")
	ErrCantidadNegativa  = errors.New("cantidad_kg no puede ser negativa")
)

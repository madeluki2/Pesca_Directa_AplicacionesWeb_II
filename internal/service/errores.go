package service

import "errors"

// Errores posibles del dominio Rutas de Distribución.
var (
	// Auth
	ErrCredencialesInvalidas = errors.New("email o contraseña incorrectos")
	ErrEmailEnUso            = errors.New("email ya registrado")

	// Genérico
	ErrNoEncontrado = errors.New("recurso no encontrado")

	// Rutas
	ErrRutaNoEncontrada = errors.New("ruta no encontrada")
	ErrNombreVacio      = errors.New("nombre es requerido")
	ErrOrigenVacio      = errors.New("origen es requerido")
	ErrDestinoVacio     = errors.New("destino es requerido")

	// Puntos
	ErrPuntoNoEncontrado = errors.New("punto no encontrado")
	ErrDireccionVacia    = errors.New("direccion es requerida")
	ErrRutaIDVacio       = errors.New("ruta_id es requerido")

	// Transportistas
	ErrTransportistaNoEncontrado = errors.New("transportista no encontrado")
	ErrTelefonoVacio             = errors.New("telefono es requerido")
	ErrPlacaVacia                = errors.New("placa_vehiculo es requerida")
	ErrPlacaDuplicada            = errors.New("la placa de vehículo ya existe")

	// Entregas
	ErrEntregaNoEncontrada  = errors.New("entrega no encontrada")
	ErrPedidoIDVacio        = errors.New("pedido_id es requerido")
	ErrPuntoIDVacio         = errors.New("punto_id es requerido")
	ErrTransportistaIDVacio = errors.New("transportista_id es requerido")
)

package service

import "errors"

// Errores del dominio Gestión de Pedidos
var (
	// Genérico
	ErrNoEncontrado = errors.New("recurso no encontrado")

	// Auth
	ErrCredencialesInvalidas = errors.New("email o contraseña incorrectos")
	ErrEmailEnUso            = errors.New("email ya registrado")

	// Clientes
	ErrClienteNoEncontrado = errors.New("cliente no encontrado")
	ErrNombreNegocioVacio  = errors.New("nombre_negocio es requerido")
	ErrTipoClienteInvalido = errors.New("tipo_cliente debe ser: restaurante, intermediario o mayorista")
	ErrTelefonoVacio       = errors.New("telefono es requerido")
	ErrDireccionVacia      = errors.New("direccion es requerida")

	// Pedidos
	ErrPedidoNoEncontrado = errors.New("pedido no encontrado")
	ErrClienteIDInvalido  = errors.New("cliente_id es requerido")
	ErrFechaVacia         = errors.New("fecha es requerida")

	// Detalles
	ErrDetalleNoEncontrado = errors.New("detalle de pedido no encontrado")
	ErrPedidoIDInvalido    = errors.New("pedido_id es requerido")
	ErrEspecieIDInvalido   = errors.New("especie_id es requerido")
	ErrCantidadInvalida    = errors.New("cantidad_kg debe ser mayor a 0")
	ErrPrecioInvalido      = errors.New("precio_unitario debe ser mayor a 0")
)

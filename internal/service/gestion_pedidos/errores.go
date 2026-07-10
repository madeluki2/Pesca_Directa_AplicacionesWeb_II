package gestion_pedidos

import "errors"

// Errores genéricos
var (
	ErrNoEncontrado          = errors.New("recurso no encontrado")
	ErrCredencialesInvalidas = errors.New("email o contraseña incorrectos")
	ErrEmailEnUso            = errors.New("email ya registrado")
)

// Errores del módulo Gestión de Pedidos
var (
	ErrClienteNoEncontrado = errors.New("cliente no encontrado")
	ErrNombreNegocioVacio  = errors.New("nombre_negocio es requerido")
	ErrTipoClienteInvalido = errors.New("tipo_cliente debe ser: restaurante, intermediario o mayorista")
	ErrTelefonoVacio       = errors.New("telefono es requerido")
	ErrDireccionVacia      = errors.New("direccion es requerida")
	ErrClienteIDInvalido   = errors.New("cliente_id es requerido")
	ErrFechaVacia          = errors.New("fecha es requerida")
	ErrPedidoNoEncontrado  = errors.New("pedido no encontrado")
	ErrPedidoIDInvalido    = errors.New("pedido_id es requerido")
	ErrEspecieIDInvalido   = errors.New("especie_id es requerido")
	ErrCantidadInvalida    = errors.New("cantidad_kg debe ser mayor a 0")
	ErrPrecioInvalido      = errors.New("precio_unitario debe ser mayor a 0")
	ErrDetalleNoEncontrado = errors.New("detalle de pedido no encontrado")
)

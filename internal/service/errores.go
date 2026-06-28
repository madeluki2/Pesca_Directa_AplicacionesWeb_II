package service

import "errors"

// Errores genéricos
var (
	ErrNoEncontrado          = errors.New("recurso no encontrado")
	ErrCredencialesInvalidas = errors.New("email o contraseña incorrectos")
	ErrEmailEnUso            = errors.New("email ya registrado")
)

// Errores del módulo Gestión de Pesca
var (
	ErrPescadorNoEncontrado    = errors.New("pescador no encontrado")
	ErrEmbarcacionNoEncontrada = errors.New("embarcacion no encontrada")
	ErrEspecieNoEncontrada     = errors.New("especie no encontrada")
	ErrCapturaNoEncontrada     = errors.New("captura no encontrada")
	ErrBodegaNoEncontrada      = errors.New("bodega no encontrada")
	ErrStockNoEncontrado       = errors.New("stock no encontrado")
	ErrCedulaVacia             = errors.New("cedula es requerida")
	ErrPuertoVacio             = errors.New("puerto_base es requerido")
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

// Errores del módulo Gestión de Pedidos
var (
	ErrClienteNoEncontrado = errors.New("cliente no encontrado")
	ErrNombreNegocioVacio  = errors.New("nombre_negocio es requerido")
	ErrTipoClienteInvalido = errors.New("tipo_cliente debe ser: restaurante, intermediario o mayorista")
	ErrTelefonoVacio       = errors.New("telefono es requerido")
	ErrDireccionVacia      = errors.New("direccion es requerida")
	ErrClienteIDInvalido   = errors.New("cliente_id es requerido")
	ErrPedidoNoEncontrado  = errors.New("pedido no encontrado")
	ErrPedidoIDInvalido    = errors.New("pedido_id es requerido")
	ErrEspecieIDInvalido   = errors.New("especie_id es requerido")
	ErrPrecioInvalido      = errors.New("precio_unitario debe ser mayor a 0")
	ErrDetalleNoEncontrado = errors.New("detalle de pedido no encontrado")
)

// Errores del módulo Rutas de Distribución
var (
	ErrRutaNoEncontrada          = errors.New("ruta no encontrada")
	ErrPuntoNoEncontrado         = errors.New("punto no encontrado")
	ErrTransportistaNoEncontrado = errors.New("transportista no encontrado")
	ErrEntregaNoEncontrada       = errors.New("entrega no encontrada")
	ErrOrigenVacio               = errors.New("origen es requerido")
	ErrDestinoVacio              = errors.New("destino es requerido")
	ErrNombreVacio               = errors.New("nombre es requerido")
	ErrRutaIDVacio               = errors.New("ruta_id es requerido")
	ErrPlacaDuplicada            = errors.New("placa_vehiculo ya registrada")
	ErrPlacaVacia                = errors.New("placa_vehiculo es requerida")
	ErrPedidoIDVacio             = errors.New("pedido_id es requerido")
	ErrPuntoIDVacio              = errors.New("punto_id es requerido")
	ErrTransportistaIDVacio      = errors.New("transportista_id es requerido")
)

package service

import (
	"strings"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

// PedidoService concentra toda la lógica de negocio del módulo de Gestión de Pedidos.
type PedidoService struct {
	repo storage.Almacen
}

// NewPedidoService crea un service con el almacén inyectado.
func NewPedidoService(repo storage.Almacen) *PedidoService {
	return &PedidoService{repo: repo}
}

// =========================================================
// CLIENTES
// =========================================================

func (s *PedidoService) ListarClientes() []models.Cliente {
	return s.repo.ListarClientes()
}

func (s *PedidoService) ObtenerCliente(id int) (models.Cliente, error) {
	c, ok := s.repo.BuscarClientePorID(id)
	if !ok {
		return models.Cliente{}, ErrClienteNoEncontrado
	}
	return c, nil
}

func (s *PedidoService) CrearCliente(c models.Cliente) (models.Cliente, error) {
	if err := validarCliente(c); err != nil {
		return models.Cliente{}, err
	}
	// El estado siempre inicia como activo
	c.Estado = "activo"
	return s.repo.CrearCliente(c), nil
}

func (s *PedidoService) ActualizarCliente(id int, datos models.Cliente) (models.Cliente, error) {
	c, ok := s.repo.ActualizarCliente(id, datos)
	if !ok {
		return models.Cliente{}, ErrClienteNoEncontrado
	}
	return c, nil
}

func (s *PedidoService) EliminarCliente(id int) error {
	if !s.repo.EliminarCliente(id) {
		return ErrClienteNoEncontrado
	}
	return nil
}

func (s *PedidoService) CambiarTipoCliente(id int, tipo string) (models.Cliente, error) {
	tipo = strings.TrimSpace(tipo)
	if tipo != "restaurante" && tipo != "intermediario" && tipo != "mayorista" {
		return models.Cliente{}, ErrTipoClienteInvalido
	}
	c, ok := s.repo.CambiarTipoCliente(id, tipo)
	if !ok {
		return models.Cliente{}, ErrClienteNoEncontrado
	}
	return c, nil
}

func validarCliente(c models.Cliente) error {
	if strings.TrimSpace(c.NombreNegocio) == "" {
		return ErrNombreNegocioVacio
	}
	if c.TipoCliente != "restaurante" && c.TipoCliente != "intermediario" && c.TipoCliente != "mayorista" {
		return ErrTipoClienteInvalido
	}
	if strings.TrimSpace(c.Telefono) == "" {
		return ErrTelefonoVacio
	}
	if strings.TrimSpace(c.Direccion) == "" {
		return ErrDireccionVacia
	}
	return nil
}

// =========================================================
// PEDIDOS
// =========================================================

func (s *PedidoService) ListarPedidos() []models.Pedido {
	return s.repo.ListarPedidos()
}

func (s *PedidoService) ObtenerPedido(id int) (models.Pedido, error) {
	p, ok := s.repo.BuscarPedidoPorID(id)
	if !ok {
		return models.Pedido{}, ErrPedidoNoEncontrado
	}
	return p, nil
}

func (s *PedidoService) CrearPedido(p models.Pedido) (models.Pedido, error) {
	if err := validarPedido(p); err != nil {
		return models.Pedido{}, err
	}
	// El estado siempre inicia como pendiente
	p.Estado = "pendiente"
	p.Total = 0
	return s.repo.CrearPedido(p), nil
}

func (s *PedidoService) ActualizarPedido(id int, datos models.Pedido) (models.Pedido, error) {
	p, ok := s.repo.ActualizarPedido(id, datos)
	if !ok {
		return models.Pedido{}, ErrPedidoNoEncontrado
	}
	return p, nil
}

func (s *PedidoService) EliminarPedido(id int) error {
	if !s.repo.EliminarPedido(id) {
		return ErrPedidoNoEncontrado
	}
	return nil
}

func validarPedido(p models.Pedido) error {
	if p.ClienteID == 0 {
		return ErrClienteIDInvalido
	}
	if strings.TrimSpace(p.Fecha) == "" {
		return ErrFechaVacia
	}
	return nil
}

// =========================================================
// DETALLES DE PEDIDO
// =========================================================

func (s *PedidoService) ListarDetalles() []models.DetallePedido {
	return s.repo.ListarDetalles()
}

func (s *PedidoService) ObtenerDetalle(id int) (models.DetallePedido, error) {
	d, ok := s.repo.BuscarDetallePorID(id)
	if !ok {
		return models.DetallePedido{}, ErrDetalleNoEncontrado
	}
	return d, nil
}

func (s *PedidoService) CrearDetalle(d models.DetallePedido) (models.DetallePedido, error) {
	if err := validarDetalle(d); err != nil {
		return models.DetallePedido{}, err
	}
	// Calculamos el subtotal automáticamente
	d.Subtotal = d.CantidadKg * d.PrecioUnitario
	return s.repo.CrearDetalle(d), nil
}

func (s *PedidoService) ActualizarDetalle(id int, datos models.DetallePedido) (models.DetallePedido, error) {
	// Recalculamos subtotal si vienen los dos campos
	if datos.CantidadKg != 0 && datos.PrecioUnitario != 0 {
		datos.Subtotal = datos.CantidadKg * datos.PrecioUnitario
	}
	d, ok := s.repo.ActualizarDetalle(id, datos)
	if !ok {
		return models.DetallePedido{}, ErrDetalleNoEncontrado
	}
	return d, nil
}

func (s *PedidoService) EliminarDetalle(id int) error {
	if !s.repo.EliminarDetalle(id) {
		return ErrDetalleNoEncontrado
	}
	return nil
}

func validarDetalle(d models.DetallePedido) error {
	if d.PedidoID == 0 {
		return ErrPedidoIDInvalido
	}
	if d.EspecieID == 0 {
		return ErrEspecieIDInvalido
	}
	if d.CantidadKg <= 0 {
		return ErrCantidadInvalida
	}
	if d.PrecioUnitario <= 0 {
		return ErrPrecioInvalido
	}
	return nil
}

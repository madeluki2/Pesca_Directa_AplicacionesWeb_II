package gestion_pedidos

import (
	"strings"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

// Almacen es la interfaz que este service necesita del repositorio.
// Se define aquí para que el service no dependa de ningún package de storage.
type Almacen interface {
	// Clientes
	ListarClientes() ([]models.Cliente, error)
	BuscarClientePorID(id int) (models.Cliente, error)
	CrearCliente(c models.Cliente) (models.Cliente, error)
	ActualizarCliente(id int, datos models.Cliente) (models.Cliente, error)
	EliminarCliente(id int) error
	CambiarTipoCliente(id int, tipo string) (models.Cliente, error)
	// Pedidos
	ListarPedidos() ([]models.Pedido, error)
	BuscarPedidoPorID(id int) (models.Pedido, error)
	CrearPedido(p models.Pedido) (models.Pedido, error)
	ActualizarPedido(id int, datos models.Pedido) (models.Pedido, error)
	EliminarPedido(id int) error
	// Detalles
	ListarDetalles() ([]models.DetallePedido, error)
	BuscarDetallePorID(id int) (models.DetallePedido, error)
	CrearDetalle(d models.DetallePedido) (models.DetallePedido, error)
	ActualizarDetalle(id int, datos models.DetallePedido) (models.DetallePedido, error)
	EliminarDetalle(id int) error
}

// PedidoService concentra toda la lógica de negocio del módulo de Gestión de Pedidos.
type PedidoService struct {
	repo Almacen
}

// NewPedidoService crea el service inyectando el repositorio.
func NewPedidoService(repo Almacen) *PedidoService {
	return &PedidoService{repo: repo}
}

// ═══════════════════════════ CLIENTES ════════════════════════════════════════

func (s *PedidoService) ListarClientes() ([]models.Cliente, error) {
	return s.repo.ListarClientes()
}

func (s *PedidoService) ObtenerCliente(id int) (models.Cliente, error) {
	c, err := s.repo.BuscarClientePorID(id)
	if err != nil {
		return models.Cliente{}, service.ErrClienteNoEncontrado
	}
	return c, nil
}

func (s *PedidoService) CrearCliente(c models.Cliente) (models.Cliente, error) {
	if err := validarCliente(c); err != nil {
		return models.Cliente{}, err
	}
	c.Estado = "activo"
	return s.repo.CrearCliente(c)
}

func (s *PedidoService) ActualizarCliente(id int, datos models.Cliente) (models.Cliente, error) {
	c, err := s.repo.ActualizarCliente(id, datos)
	if err != nil {
		return models.Cliente{}, service.ErrClienteNoEncontrado
	}
	return c, nil
}

func (s *PedidoService) EliminarCliente(id int) error {
	if err := s.repo.EliminarCliente(id); err != nil {
		return service.ErrClienteNoEncontrado
	}
	return nil
}

func (s *PedidoService) CambiarTipoCliente(id int, tipo string) (models.Cliente, error) {
	tipo = strings.TrimSpace(tipo)
	if tipo != "restaurante" && tipo != "intermediario" && tipo != "mayorista" {
		return models.Cliente{}, service.ErrTipoClienteInvalido
	}
	c, err := s.repo.CambiarTipoCliente(id, tipo)
	if err != nil {
		return models.Cliente{}, service.ErrClienteNoEncontrado
	}
	return c, nil
}

func validarCliente(c models.Cliente) error {
	if strings.TrimSpace(c.NombreNegocio) == "" {
		return service.ErrNombreNegocioVacio
	}
	if c.TipoCliente != "restaurante" && c.TipoCliente != "intermediario" && c.TipoCliente != "mayorista" {
		return service.ErrTipoClienteInvalido
	}
	if strings.TrimSpace(c.Telefono) == "" {
		return service.ErrTelefonoVacio
	}
	if strings.TrimSpace(c.Direccion) == "" {
		return service.ErrDireccionVacia
	}
	return nil
}

// ═══════════════════════════ PEDIDOS ═════════════════════════════════════════

func (s *PedidoService) ListarPedidos() ([]models.Pedido, error) {
	return s.repo.ListarPedidos()
}

func (s *PedidoService) ObtenerPedido(id int) (models.Pedido, error) {
	p, err := s.repo.BuscarPedidoPorID(id)
	if err != nil {
		return models.Pedido{}, service.ErrPedidoNoEncontrado
	}
	return p, nil
}

func (s *PedidoService) CrearPedido(p models.Pedido) (models.Pedido, error) {
	if p.ClienteID == 0 {
		return models.Pedido{}, service.ErrClienteIDInvalido
	}
	if p.Fecha.IsZero() {
		return models.Pedido{}, service.ErrFechaVacia
	}
	p.Estado = "pendiente"
	return s.repo.CrearPedido(p)
}

func (s *PedidoService) ActualizarPedido(id int, datos models.Pedido) (models.Pedido, error) {
	p, err := s.repo.ActualizarPedido(id, datos)
	if err != nil {
		return models.Pedido{}, service.ErrPedidoNoEncontrado
	}
	return p, nil
}

func (s *PedidoService) EliminarPedido(id int) error {
	if err := s.repo.EliminarPedido(id); err != nil {
		return service.ErrPedidoNoEncontrado
	}
	return nil
}

// ═══════════════════════════ DETALLES ════════════════════════════════════════

func (s *PedidoService) ListarDetalles() ([]models.DetallePedido, error) {
	return s.repo.ListarDetalles()
}

func (s *PedidoService) ObtenerDetalle(id int) (models.DetallePedido, error) {
	d, err := s.repo.BuscarDetallePorID(id)
	if err != nil {
		return models.DetallePedido{}, service.ErrDetalleNoEncontrado
	}
	return d, nil
}

func (s *PedidoService) CrearDetalle(d models.DetallePedido) (models.DetallePedido, error) {
	if d.PedidoID == 0 {
		return models.DetallePedido{}, service.ErrPedidoIDInvalido
	}
	if d.EspecieID == 0 {
		return models.DetallePedido{}, service.ErrEspecieIDInvalido
	}
	if d.CantidadKg <= 0 {
		return models.DetallePedido{}, service.ErrCantidadInvalida
	}
	if d.PrecioUnitario <= 0 {
		return models.DetallePedido{}, service.ErrPrecioInvalido
	}
	d.Subtotal = d.CantidadKg * d.PrecioUnitario
	return s.repo.CrearDetalle(d)
}

func (s *PedidoService) ActualizarDetalle(id int, datos models.DetallePedido) (models.DetallePedido, error) {
	if datos.CantidadKg != 0 && datos.PrecioUnitario != 0 {
		datos.Subtotal = datos.CantidadKg * datos.PrecioUnitario
	}
	d, err := s.repo.ActualizarDetalle(id, datos)
	if err != nil {
		return models.DetallePedido{}, service.ErrDetalleNoEncontrado
	}
	return d, nil
}

func (s *PedidoService) EliminarDetalle(id int) error {
	if err := s.repo.EliminarDetalle(id); err != nil {
		return service.ErrDetalleNoEncontrado
	}
	return nil
}

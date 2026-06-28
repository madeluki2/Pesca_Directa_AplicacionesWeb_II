// Package storage gestiona el almacenamiento en memoria del módulo
// de Gestión de Pedidos de Pesca-Directa Tarqui.
//
// El tipo Memoria mantiene en un solo lugar todos los datos del dominio:
// Clientes, Pedidos y Detalles de Pedido.
package storage

import (
	"sync"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// Memoria es un almacén unificado del módulo de Gestión de Pedidos.
type Memoria struct {
	clientes      []models.Cliente
	nextClienteID int

	pedidos      []models.Pedido
	nextPedidoID int

	detalles      []models.DetallePedido
	nextDetalleID int

	mu sync.Mutex
}

// NewMemoria crea un almacén vacío y listo para usar.
func NewMemoria() *Memoria {
	return &Memoria{
		clientes:      []models.Cliente{},
		nextClienteID: 1,
		pedidos:       []models.Pedido{},
		nextPedidoID:  1,
		detalles:      []models.DetallePedido{},
		nextDetalleID: 1,
	}
}

// =========================================================
// SEED
// =========================================================

// Seed carga datos iniciales de prueba en memoria.
func (m *Memoria) Seed() {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Clientes de prueba
	m.clientes = []models.Cliente{
		{ID: 1, UsuarioID: 1, TipoCliente: "restaurante", NombreNegocio: "Sushi Koi", Direccion: "Av. Flavio Reyes", Telefono: "0991234567", Estado: "activo"},
		{ID: 2, UsuarioID: 2, TipoCliente: "mayorista", NombreNegocio: "Distribuidora El Puerto", Direccion: "Calle 10 de Agosto", Telefono: "0997654321", Estado: "activo"},
	}
	m.nextClienteID = 3

	// Pedidos de prueba
	m.pedidos = []models.Pedido{
		{ID: 1, ClienteID: 1, Fecha: "2026-06-10", Estado: "pendiente", Total: 150.00},
		{ID: 2, ClienteID: 2, Fecha: "2026-06-11", Estado: "en_proceso", Total: 320.50},
	}
	m.nextPedidoID = 3

	// Detalles de prueba
	m.detalles = []models.DetallePedido{
		{ID: 1, PedidoID: 1, EspecieID: 1, CantidadKg: 10.0, PrecioUnitario: 8.50, Subtotal: 85.00},
		{ID: 2, PedidoID: 1, EspecieID: 2, CantidadKg: 5.0, PrecioUnitario: 13.00, Subtotal: 65.00},
	}
	m.nextDetalleID = 3
}

// =========================================================
// CLIENTES
// =========================================================

// ListarClientes devuelve todos los clientes en memoria.
func (m *Memoria) ListarClientes() []models.Cliente {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Cliente, len(m.clientes))
	copy(copia, m.clientes)
	return copia
}

// BuscarClientePorID devuelve el cliente con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarClientePorID(id int) (models.Cliente, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, c := range m.clientes {
		if c.ID == id {
			return c, true
		}
	}
	return models.Cliente{}, false
}

// CrearCliente agrega un cliente nuevo y devuelve el cliente con ID asignado.
func (m *Memoria) CrearCliente(c models.Cliente) models.Cliente {
	m.mu.Lock()
	defer m.mu.Unlock()

	c.ID = m.nextClienteID
	m.nextClienteID++
	m.clientes = append(m.clientes, c)
	return c
}

// ActualizarCliente reemplaza el cliente con el ID dado.
func (m *Memoria) ActualizarCliente(id int, datos models.Cliente) (models.Cliente, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.clientes {
		if c.ID == id {
			datos.ID = id
			if datos.UsuarioID == 0 {
				datos.UsuarioID = c.UsuarioID
			}
			if datos.TipoCliente == "" {
				datos.TipoCliente = c.TipoCliente
			}
			if datos.NombreNegocio == "" {
				datos.NombreNegocio = c.NombreNegocio
			}
			if datos.Direccion == "" {
				datos.Direccion = c.Direccion
			}
			if datos.Telefono == "" {
				datos.Telefono = c.Telefono
			}
			if datos.Estado == "" {
				datos.Estado = c.Estado
			}
			m.clientes[i] = datos
			return datos, true
		}
	}
	return models.Cliente{}, false
}

// EliminarCliente elimina el cliente con el ID dado.
func (m *Memoria) EliminarCliente(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.clientes {
		if c.ID == id {
			m.clientes = append(m.clientes[:i], m.clientes[i+1:]...)
			return true
		}
	}
	return false
}

// CambiarTipoCliente actualiza únicamente el tipo de un cliente existente.
func (m *Memoria) CambiarTipoCliente(id int, nuevoTipo string) (models.Cliente, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.clientes {
		if c.ID == id {
			m.clientes[i].TipoCliente = nuevoTipo
			return m.clientes[i], true
		}
	}
	return models.Cliente{}, false
}

// =========================================================
// PEDIDOS
// =========================================================

// ListarPedidos devuelve todos los pedidos en memoria.
func (m *Memoria) ListarPedidos() []models.Pedido {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Pedido, len(m.pedidos))
	copy(copia, m.pedidos)
	return copia
}

// BuscarPedidoPorID devuelve el pedido con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarPedidoPorID(id int) (models.Pedido, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, p := range m.pedidos {
		if p.ID == id {
			return p, true
		}
	}
	return models.Pedido{}, false
}

// CrearPedido agrega un pedido nuevo y devuelve el pedido con ID asignado.
func (m *Memoria) CrearPedido(p models.Pedido) models.Pedido {
	m.mu.Lock()
	defer m.mu.Unlock()

	p.ID = m.nextPedidoID
	m.nextPedidoID++
	m.pedidos = append(m.pedidos, p)
	return p
}

// ActualizarPedido reemplaza el pedido con el ID dado.
func (m *Memoria) ActualizarPedido(id int, datos models.Pedido) (models.Pedido, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.pedidos {
		if p.ID == id {
			datos.ID = id
			if datos.Estado == "" {
				datos.Estado = p.Estado
			}
			if datos.Fecha == "" {
				datos.Fecha = p.Fecha
			}
			if datos.Total == 0 {
				datos.Total = p.Total
			}
			if datos.ClienteID == 0 {
				datos.ClienteID = p.ClienteID
			}
			m.pedidos[i] = datos
			return datos, true
		}
	}
	return models.Pedido{}, false
}

// EliminarPedido elimina el pedido con el ID dado.
func (m *Memoria) EliminarPedido(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.pedidos {
		if p.ID == id {
			m.pedidos = append(m.pedidos[:i], m.pedidos[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// DETALLES DE PEDIDO
// =========================================================

// ListarDetalles devuelve todos los detalles en memoria.
func (m *Memoria) ListarDetalles() []models.DetallePedido {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.DetallePedido, len(m.detalles))
	copy(copia, m.detalles)
	return copia
}

// BuscarDetallePorID devuelve el detalle con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarDetallePorID(id int) (models.DetallePedido, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, d := range m.detalles {
		if d.ID == id {
			return d, true
		}
	}
	return models.DetallePedido{}, false
}

// CrearDetalle agrega un detalle nuevo y devuelve el detalle con ID asignado.
func (m *Memoria) CrearDetalle(d models.DetallePedido) models.DetallePedido {
	m.mu.Lock()
	defer m.mu.Unlock()

	d.ID = m.nextDetalleID
	m.nextDetalleID++
	m.detalles = append(m.detalles, d)
	return d
}

// ActualizarDetalle reemplaza el detalle con el ID dado.
func (m *Memoria) ActualizarDetalle(id int, datos models.DetallePedido) (models.DetallePedido, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, d := range m.detalles {
		if d.ID == id {
			datos.ID = id
			if datos.CantidadKg == 0 {
				datos.CantidadKg = d.CantidadKg
			}
			if datos.PrecioUnitario == 0 {
				datos.PrecioUnitario = d.PrecioUnitario
			}
			if datos.Subtotal == 0 {
				datos.Subtotal = d.Subtotal
			}
			if datos.EspecieID == 0 {
				datos.EspecieID = d.EspecieID
			}
			if datos.PedidoID == 0 {
				datos.PedidoID = d.PedidoID
			}
			m.detalles[i] = datos
			return datos, true
		}
	}
	return models.DetallePedido{}, false
}

// EliminarDetalle elimina el detalle con el ID dado.
func (m *Memoria) EliminarDetalle(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, d := range m.detalles {
		if d.ID == id {
			m.detalles = append(m.detalles[:i], m.detalles[i+1:]...)
			return true
		}
	}
	return false
}

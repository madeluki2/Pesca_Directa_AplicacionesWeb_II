package storage

import (
	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"errors"
)

// ─────────────────────────────────────────────
// Errores de dominio reutilizables
// ─────────────────────────────────────────────

var (
	ErrNotFound       = errors.New("registro no encontrado")
	ErrPlacaDuplicada = errors.New("la placa de vehículo ya existe")
)

// ══════════════════════════════════════════════
// INTERFACES
// ══════════════════════════════════════════════

// RutaStore define las operaciones CRUD para Ruta.
type AlmacenRuta interface {
	CrearRuta(r models.Ruta) (models.Ruta, error)
	ObtenerRutas() ([]models.Ruta, error)
	ObtenerRutaPorID(id uint) (models.Ruta, error)
	ActualizarRuta(id uint, r models.Ruta) (models.Ruta, error)
	EliminarRuta(id uint) error

	// PuntoStore define las operaciones CRUD para Punto.
	CrearPunto(p models.Punto) (models.Punto, error)
	ObtenerPuntos() ([]models.Punto, error)
	ObtenerPuntoPorID(id uint) (models.Punto, error)
	ActualizarPunto(id uint, p models.Punto) (models.Punto, error)
	EliminarPunto(id uint) error

	// TransportistaStore define las operaciones CRUD para Transportista.
	CrearTransportista(t models.Transportista) (models.Transportista, error)
	ObtenerTransportistas() ([]models.Transportista, error)
	ObtenerTransportistaPorID(id uint) (models.Transportista, error)
	ActualizarTransportista(id uint, t models.Transportista) (models.Transportista, error)
	EliminarTransportista(id uint) error

	// EntregaStore define las operaciones CRUD para EntregaPedido.
	CrearEntrega(e models.EntregaPedido) (models.EntregaPedido, error)
	ObtenerEntregas() ([]models.EntregaPedido, error)
	ObtenerEntregaPorID(id uint) (models.EntregaPedido, error)
	ActualizarEntrega(id uint, e models.EntregaPedido) (models.EntregaPedido, error)
	EliminarEntrega(id uint) error
}

var _ AlmacenRuta = (*MemoriaRutas)(nil)

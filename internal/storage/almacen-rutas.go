package storage

import (
	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// ══════════════════════════════════════════════
// AlmacenRutas: contrato único para los backends
// de persistencia del módulo Rutas de Distribución
// (MemoriaRutas y AlmacenSQLiteRutas lo implementan).
//
// Patrón: los métodos de búsqueda devuelven (T, bool) en vez de
// (T, error). El bool indica si el registro existe; es el SERVICE
// (internal/service/rutas.go) quien traduce ese bool a un error
// de dominio (ErrRutaNoEncontrada, etc). El storage queda como
// CRUD puro, sin lógica de negocio.
// ══════════════════════════════════════════════
type AlmacenRutas interface {
	// Rutas
	ListarRutas() []models.Ruta
	BuscarRutaPorID(id uint) (models.Ruta, bool)
	CrearRuta(r models.Ruta) models.Ruta
	ActualizarRuta(id uint, r models.Ruta) (models.Ruta, bool)
	BorrarRuta(id uint) bool

	// Puntos
	ListarPuntos() []models.Punto
	BuscarPuntoPorID(id uint) (models.Punto, bool)
	CrearPunto(p models.Punto) models.Punto
	ActualizarPunto(id uint, p models.Punto) (models.Punto, bool)
	BorrarPunto(id uint) bool

	// Transportistas
	ListarTransportistas() []models.Transportista
	BuscarTransportistaPorID(id uint) (models.Transportista, bool)
	CrearTransportista(t models.Transportista) models.Transportista
	ActualizarTransportista(id uint, t models.Transportista) (models.Transportista, bool)
	BorrarTransportista(id uint) bool

	// Entregas
	ListarEntregas() []models.EntregaPedido
	BuscarEntregaPorID(id uint) (models.EntregaPedido, bool)
	CrearEntrega(e models.EntregaPedido) models.EntregaPedido
	ActualizarEntrega(id uint, e models.EntregaPedido) (models.EntregaPedido, bool)
	BorrarEntrega(id uint) bool
}

// UserRepository: contrato para la persistencia de Usuario (auth).
// Siempre se sirve desde GORM (UsuarioGORM), nunca desde memoria,
// porque register/login necesitan persistencia real.
type UserRepository interface {
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
	CrearUsuario(u models.Usuario) (models.Usuario, error)
}

var _ AlmacenRutas = (*AlmacenSQLiteRutas)(nil)

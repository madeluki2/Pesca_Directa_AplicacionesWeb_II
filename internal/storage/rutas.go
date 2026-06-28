package storage

import (
	"Pesca_Directa_AplicacionesWeb_II/internal/models"

	"gorm.io/gorm"
)

// AlmacenSQLiteRutas implementa AlmacenRutas usando GORM + SQLite.
type AlmacenSQLiteRutas struct {
	db *gorm.DB
}

// NuevoAlmacenSQLiteRutas es el constructor que main.go llama.
func NuevoAlmacenSQLiteRutas(db *gorm.DB) *AlmacenSQLiteRutas {
	return &AlmacenSQLiteRutas{db: db}
}

// ── Rutas ─────────────────────────────────────

func (s *AlmacenSQLiteRutas) ListarRutas() []models.Ruta {
	var rutas []models.Ruta
	s.db.Find(&rutas)
	return rutas
}

func (s *AlmacenSQLiteRutas) BuscarRutaPorID(id uint) (models.Ruta, bool) {
	var r models.Ruta
	if err := s.db.First(&r, id).Error; err != nil {
		return models.Ruta{}, false
	}
	return r, true
}

func (s *AlmacenSQLiteRutas) CrearRuta(r models.Ruta) models.Ruta {
	s.db.Create(&r)
	return r
}

func (s *AlmacenSQLiteRutas) ActualizarRuta(id uint, nuevo models.Ruta) (models.Ruta, bool) {
	var r models.Ruta
	if err := s.db.First(&r, id).Error; err != nil {
		return models.Ruta{}, false
	}
	nuevo.ID = id
	nuevo.CreadoEn = r.CreadoEn
	s.db.Save(&nuevo)
	return nuevo, true
}

func (s *AlmacenSQLiteRutas) BorrarRuta(id uint) bool {
	result := s.db.Delete(&models.Ruta{}, id)
	return result.RowsAffected > 0
}

// ── Puntos ────────────────────────────────────

func (s *AlmacenSQLiteRutas) ListarPuntos() []models.Punto {
	var puntos []models.Punto
	s.db.Find(&puntos)
	return puntos
}

func (s *AlmacenSQLiteRutas) BuscarPuntoPorID(id uint) (models.Punto, bool) {
	var p models.Punto
	if err := s.db.First(&p, id).Error; err != nil {
		return models.Punto{}, false
	}
	return p, true
}

func (s *AlmacenSQLiteRutas) CrearPunto(p models.Punto) models.Punto {
	s.db.Create(&p)
	return p
}

func (s *AlmacenSQLiteRutas) ActualizarPunto(id uint, nuevo models.Punto) (models.Punto, bool) {
	var p models.Punto
	if err := s.db.First(&p, id).Error; err != nil {
		return models.Punto{}, false
	}
	nuevo.ID = id
	nuevo.CreadoEn = p.CreadoEn
	s.db.Save(&nuevo)
	return nuevo, true
}

func (s *AlmacenSQLiteRutas) BorrarPunto(id uint) bool {
	result := s.db.Delete(&models.Punto{}, id)
	return result.RowsAffected > 0
}

// ── Transportistas ────────────────────────────

func (s *AlmacenSQLiteRutas) ListarTransportistas() []models.Transportista {
	var lista []models.Transportista
	s.db.Find(&lista)
	return lista
}

func (s *AlmacenSQLiteRutas) BuscarTransportistaPorID(id uint) (models.Transportista, bool) {
	var t models.Transportista
	if err := s.db.First(&t, id).Error; err != nil {
		return models.Transportista{}, false
	}
	return t, true
}

func (s *AlmacenSQLiteRutas) CrearTransportista(t models.Transportista) models.Transportista {
	s.db.Create(&t)
	return t
}

func (s *AlmacenSQLiteRutas) ActualizarTransportista(id uint, nuevo models.Transportista) (models.Transportista, bool) {
	var t models.Transportista
	if err := s.db.First(&t, id).Error; err != nil {
		return models.Transportista{}, false
	}
	nuevo.ID = id
	nuevo.CreadoEn = t.CreadoEn
	s.db.Save(&nuevo)
	return nuevo, true
}

func (s *AlmacenSQLiteRutas) BorrarTransportista(id uint) bool {
	result := s.db.Delete(&models.Transportista{}, id)
	return result.RowsAffected > 0
}

// ── Entregas ──────────────────────────────────

func (s *AlmacenSQLiteRutas) ListarEntregas() []models.EntregaPedido {
	var lista []models.EntregaPedido
	s.db.Find(&lista)
	return lista
}

func (s *AlmacenSQLiteRutas) BuscarEntregaPorID(id uint) (models.EntregaPedido, bool) {
	var e models.EntregaPedido
	if err := s.db.First(&e, id).Error; err != nil {
		return models.EntregaPedido{}, false
	}
	return e, true
}


func (s *AlmacenSQLiteRutas) CrearEntrega(e models.EntregaPedido) models.EntregaPedido {
	s.db.Create(&e)
	return e
}

func (s *AlmacenSQLiteRutas) ActualizarEntrega(id uint, nuevo models.EntregaPedido) (models.EntregaPedido, bool) {
	var e models.EntregaPedido
	if err := s.db.First(&e, id).Error; err != nil {
		return models.EntregaPedido{}, false
	}
	nuevo.ID = id
	nuevo.CreadoEn = e.CreadoEn
	s.db.Save(&nuevo)
	return nuevo, true
}

func (s *AlmacenSQLiteRutas) BorrarEntrega(id uint) bool {
	result := s.db.Delete(&models.EntregaPedido{}, id)
	return result.RowsAffected > 0
}

// Verificación en tiempo de compilación: AlmacenSQLiteRutas implementa AlmacenRutas.
var _ AlmacenRutas = (*AlmacenSQLiteRutas)(nil)
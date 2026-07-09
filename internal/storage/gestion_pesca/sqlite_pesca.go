package storage

import (
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// AlmacenSQLitePesca implementa AlmacenPesca usando GORM sobre SQLite.
type AlmacenSQLitePesca struct {
	db *gorm.DB
}

// NuevoAlmacenSQLitePesca envuelve una conexión *gorm.DB ya abierta y migrada.
func NuevoAlmacenSQLitePesca(db *gorm.DB) *AlmacenSQLitePesca {
	return &AlmacenSQLitePesca{db: db}
}

// ═══════════════════════════ PESCADORES ══════════════════════════════════════

func (a *AlmacenSQLitePesca) ListarPescadores() []models.Pescador {
	var lista []models.Pescador
	a.db.Find(&lista)
	return lista
}

func (a *AlmacenSQLitePesca) BuscarPescadorPorID(id int) (models.Pescador, bool) {
	var p models.Pescador
	if err := a.db.First(&p, id).Error; err != nil {
		return models.Pescador{}, false
	}
	return p, true
}

func (a *AlmacenSQLitePesca) CrearPescador(p models.Pescador) models.Pescador {
	a.db.Create(&p)
	return p
}

func (a *AlmacenSQLitePesca) ActualizarPescador(id int, datos models.Pescador) (models.Pescador, bool) {
	var existente models.Pescador
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Pescador{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLitePesca) BorrarPescador(id int) bool {
	res := a.db.Delete(&models.Pescador{}, id)
	return res.RowsAffected > 0
}

// ═══════════════════════════ EMBARCACIONES ════════════════════════════════════

func (a *AlmacenSQLitePesca) ListarEmbarcaciones() []models.Embarcacion {
	var lista []models.Embarcacion
	a.db.Find(&lista)
	return lista
}

func (a *AlmacenSQLitePesca) BuscarEmbarcacionPorID(id int) (models.Embarcacion, bool) {
	var e models.Embarcacion
	if err := a.db.First(&e, id).Error; err != nil {
		return models.Embarcacion{}, false
	}
	return e, true
}

func (a *AlmacenSQLitePesca) CrearEmbarcacion(e models.Embarcacion) models.Embarcacion {
	a.db.Create(&e)
	return e
}

func (a *AlmacenSQLitePesca) ActualizarEmbarcacion(id int, datos models.Embarcacion) (models.Embarcacion, bool) {
	var existente models.Embarcacion
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Embarcacion{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLitePesca) BorrarEmbarcacion(id int) bool {
	res := a.db.Delete(&models.Embarcacion{}, id)
	return res.RowsAffected > 0
}

// ═══════════════════════════ ESPECIES ════════════════════════════════════════

func (a *AlmacenSQLitePesca) ListarEspecies() []models.Especie {
	var lista []models.Especie
	a.db.Find(&lista)
	return lista
}

func (a *AlmacenSQLitePesca) BuscarEspeciePorID(id int) (models.Especie, bool) {
	var e models.Especie
	if err := a.db.First(&e, id).Error; err != nil {
		return models.Especie{}, false
	}
	return e, true
}

func (a *AlmacenSQLitePesca) CrearEspecie(e models.Especie) models.Especie {
	a.db.Create(&e)
	return e
}

func (a *AlmacenSQLitePesca) ActualizarEspecie(id int, datos models.Especie) (models.Especie, bool) {
	var existente models.Especie
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Especie{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLitePesca) BorrarEspecie(id int) bool {
	res := a.db.Delete(&models.Especie{}, id)
	return res.RowsAffected > 0
}

// ═══════════════════════════ CAPTURAS ════════════════════════════════════════

func (a *AlmacenSQLitePesca) ListarCapturas() []models.Captura {
	var lista []models.Captura
	a.db.Find(&lista)
	return lista
}

func (a *AlmacenSQLitePesca) BuscarCapturaPorID(id int) (models.Captura, bool) {
	var c models.Captura
	if err := a.db.First(&c, id).Error; err != nil {
		return models.Captura{}, false
	}
	return c, true
}

func (a *AlmacenSQLitePesca) CrearCaptura(c models.Captura) models.Captura {
	a.db.Create(&c)
	return c
}

func (a *AlmacenSQLitePesca) ActualizarCaptura(id int, datos models.Captura) (models.Captura, bool) {
	var existente models.Captura
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Captura{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLitePesca) BorrarCaptura(id int) bool {
	res := a.db.Delete(&models.Captura{}, id)
	return res.RowsAffected > 0
}

// ═══════════════════════════ BODEGAS ═════════════════════════════════════════

func (a *AlmacenSQLitePesca) ListarBodegas() []models.Bodega {
	var lista []models.Bodega
	a.db.Find(&lista)
	return lista
}

func (a *AlmacenSQLitePesca) BuscarBodegaPorID(id int) (models.Bodega, bool) {
	var b models.Bodega
	if err := a.db.First(&b, id).Error; err != nil {
		return models.Bodega{}, false
	}
	return b, true
}

func (a *AlmacenSQLitePesca) CrearBodega(b models.Bodega) models.Bodega {
	a.db.Create(&b)
	return b
}

func (a *AlmacenSQLitePesca) ActualizarBodega(id int, datos models.Bodega) (models.Bodega, bool) {
	var existente models.Bodega
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Bodega{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLitePesca) BorrarBodega(id int) bool {
	res := a.db.Delete(&models.Bodega{}, id)
	return res.RowsAffected > 0
}

// ═══════════════════════════ STOCKS ══════════════════════════════════════════

func (a *AlmacenSQLitePesca) ListarStocks() []models.Stock {
	var lista []models.Stock
	a.db.Find(&lista)
	return lista
}

func (a *AlmacenSQLitePesca) BuscarStockPorID(id int) (models.Stock, bool) {
	var s models.Stock
	if err := a.db.First(&s, id).Error; err != nil {
		return models.Stock{}, false
	}
	return s, true
}

func (a *AlmacenSQLitePesca) CrearStock(s models.Stock) models.Stock {
	a.db.Create(&s)
	return s
}

func (a *AlmacenSQLitePesca) ActualizarStock(id int, datos models.Stock) (models.Stock, bool) {
	var existente models.Stock
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Stock{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLitePesca) BorrarStock(id int) bool {
	res := a.db.Delete(&models.Stock{}, id)
	return res.RowsAffected > 0
}

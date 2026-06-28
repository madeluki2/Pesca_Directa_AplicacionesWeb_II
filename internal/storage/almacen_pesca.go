package storage

import "Pesca_Directa_AplicacionesWeb_II/internal/models"

type PescadorRepository interface {
	ListarPescadores() []models.Pescador
	BuscarPescadorPorID(id int) (models.Pescador, bool)
	CrearPescador(p models.Pescador) models.Pescador
	ActualizarPescador(id int, datos models.Pescador) (models.Pescador, bool)
	BorrarPescador(id int) bool
}

type EmbarcacionRepository interface {
	ListarEmbarcaciones() []models.Embarcacion
	BuscarEmbarcacionPorID(id int) (models.Embarcacion, bool)
	CrearEmbarcacion(e models.Embarcacion) models.Embarcacion
	ActualizarEmbarcacion(id int, datos models.Embarcacion) (models.Embarcacion, bool)
	BorrarEmbarcacion(id int) bool
}

type EspecieRepository interface {
	ListarEspecies() []models.Especie
	BuscarEspeciePorID(id int) (models.Especie, bool)
	CrearEspecie(e models.Especie) models.Especie
	ActualizarEspecie(id int, datos models.Especie) (models.Especie, bool)
	BorrarEspecie(id int) bool
}

type CapturaRepository interface {
	ListarCapturas() []models.Captura
	BuscarCapturaPorID(id int) (models.Captura, bool)
	CrearCaptura(c models.Captura) models.Captura
	ActualizarCaptura(id int, datos models.Captura) (models.Captura, bool)
	BorrarCaptura(id int) bool
}

type BodegaRepository interface {
	ListarBodegas() []models.Bodega
	BuscarBodegaPorID(id int) (models.Bodega, bool)
	CrearBodega(b models.Bodega) models.Bodega
	ActualizarBodega(id int, datos models.Bodega) (models.Bodega, bool)
	BorrarBodega(id int) bool
}

type StockRepository interface {
	ListarStocks() []models.Stock
	BuscarStockPorID(id int) (models.Stock, bool)
	CrearStock(s models.Stock) models.Stock
	ActualizarStock(id int, datos models.Stock) (models.Stock, bool)
	BorrarStock(id int) bool
}

// AlmacenPesca une todos los repositorios del módulo de pesca.
type AlmacenPesca interface {
	PescadorRepository
	EmbarcacionRepository
	EspecieRepository
	CapturaRepository
	BodegaRepository
	StockRepository
}

// UserRepository une todos los repositorios del módulo de usuarios.
type UserRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

// Verificación
var _ AlmacenPesca = (*MemoriaPesca)(nil)
var _ AlmacenPesca = (*AlmacenSQLitePesca)(nil)
var _ UserRepository = (*UsuarioGORM)(nil)

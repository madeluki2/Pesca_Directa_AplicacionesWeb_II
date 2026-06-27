package storage

import "Pesca_Directa_AplicacionesWeb_II/internal/models"

// AlmacenPesca define todas las operaciones de persistencia del módulo.
// Al usar una interfaz, si en el futuro cambias de memoria a SQLite
// solo cambias la implementación — los handlers no se tocan.

type AlmacenPesca interface {

	// Usuarios
	ListarUsuarios() []models.Usuario
	BuscarUsuarioPorID(id int) (models.Usuario, bool)
	CrearUsuario(u models.Usuario) models.Usuario
	ActualizarUsuario(id int, datos models.Usuario) (models.Usuario, bool)
	BorrarUsuario(id int) bool

	// Pescadores
	ListarPescadores() []models.Pescador
	BuscarPescadorPorID(id int) (models.Pescador, bool)
	CrearPescador(p models.Pescador) models.Pescador
	ActualizarPescador(id int, datos models.Pescador) (models.Pescador, bool)
	BorrarPescador(id int) bool

	// Embarcaciones
	ListarEmbarcaciones() []models.Embarcacion
	BuscarEmbarcacionPorID(id int) (models.Embarcacion, bool)
	CrearEmbarcacion(e models.Embarcacion) models.Embarcacion
	ActualizarEmbarcacion(id int, datos models.Embarcacion) (models.Embarcacion, bool)
	BorrarEmbarcacion(id int) bool

	// Especies
	ListarEspecies() []models.Especie
	BuscarEspeciePorID(id int) (models.Especie, bool)
	CrearEspecie(e models.Especie) models.Especie
	ActualizarEspecie(id int, datos models.Especie) (models.Especie, bool)
	BorrarEspecie(id int) bool

	// Capturas
	ListarCapturas() []models.Captura
	BuscarCapturaPorID(id int) (models.Captura, bool)
	CrearCaptura(c models.Captura) models.Captura
	ActualizarCaptura(id int, datos models.Captura) (models.Captura, bool)
	BorrarCaptura(id int) bool

	// Bodegas
	ListarBodegas() []models.Bodega
	BuscarBodegaPorID(id int) (models.Bodega, bool)
	CrearBodega(b models.Bodega) models.Bodega
	ActualizarBodega(id int, datos models.Bodega) (models.Bodega, bool)
	BorrarBodega(id int) bool

	// Stocks
	ListarStocks() []models.Stock
	BuscarStockPorID(id int) (models.Stock, bool)
	CrearStock(s models.Stock) models.Stock
	ActualizarStock(id int, datos models.Stock) (models.Stock, bool)
	BorrarStock(id int) bool
}

// Verificación en tiemmpo de compilación de que MemoriaPesca no cumple la interfaz
var _ AlmacenPesca = (*MemoriaPesca)(nil)

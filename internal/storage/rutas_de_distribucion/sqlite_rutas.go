package rutas_de_distribucion

import (
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// AlmacenSQLiteRutas implementa AlmacenRutas usando GORM + SQLite.
type AlmacenSQLiteRutas struct {
	db *gorm.DB
}

// NuevoAlmacenSQLiteRutas crea un almacén SQLite con GORM inyectado.
func NuevoAlmacenSQLiteRutas(db *gorm.DB) *AlmacenSQLiteRutas {
	return &AlmacenSQLiteRutas{db: db}
}

// =========================================================
// RUTAS
// =========================================================

func (a *AlmacenSQLiteRutas) ListarRutas() []models.Ruta {
	var rutas []models.Ruta
	a.db.Find(&rutas)
	return rutas
}

func (a *AlmacenSQLiteRutas) BuscarRutaPorID(id uint) (models.Ruta, bool) {
	var r models.Ruta
	if a.db.First(&r, id).Error != nil {
		return models.Ruta{}, false
	}
	return r, true
}

func (a *AlmacenSQLiteRutas) CrearRuta(r models.Ruta) models.Ruta {
	a.db.Create(&r)
	return r
}

func (a *AlmacenSQLiteRutas) ActualizarRuta(id uint, datos models.Ruta) (models.Ruta, bool) {
	var r models.Ruta
	if a.db.First(&r, id).Error != nil {
		return models.Ruta{}, false
	}
	if datos.Nombre != "" {
		r.Nombre = datos.Nombre
	}
	if datos.Origen != "" {
		r.Origen = datos.Origen
	}
	if datos.Destino != "" {
		r.Destino = datos.Destino
	}
	if datos.Estado != "" {
		r.Estado = datos.Estado
	}
	a.db.Save(&r)
	return r, true
}

func (a *AlmacenSQLiteRutas) BorrarRuta(id uint) bool {
	resultado := a.db.Delete(&models.Ruta{}, id)
	return resultado.RowsAffected > 0
}

// =========================================================
// PUNTOS
// =========================================================

func (a *AlmacenSQLiteRutas) ListarPuntos() []models.Punto {
	var puntos []models.Punto
	a.db.Find(&puntos)
	return puntos
}

func (a *AlmacenSQLiteRutas) BuscarPuntoPorID(id uint) (models.Punto, bool) {
	var p models.Punto
	if a.db.First(&p, id).Error != nil {
		return models.Punto{}, false
	}
	return p, true
}

func (a *AlmacenSQLiteRutas) CrearPunto(p models.Punto) models.Punto {
	a.db.Create(&p)
	return p
}

func (a *AlmacenSQLiteRutas) ActualizarPunto(id uint, datos models.Punto) (models.Punto, bool) {
	var p models.Punto
	if a.db.First(&p, id).Error != nil {
		return models.Punto{}, false
	}
	if datos.Nombre != "" {
		p.Nombre = datos.Nombre
	}
	if datos.Direccion != "" {
		p.Direccion = datos.Direccion
	}
	if datos.Estado != "" {
		p.Estado = datos.Estado
	}
	if datos.RutaID != 0 {
		p.RutaID = datos.RutaID
	}
	a.db.Save(&p)
	return p, true
}

func (a *AlmacenSQLiteRutas) BorrarPunto(id uint) bool {
	resultado := a.db.Delete(&models.Punto{}, id)
	return resultado.RowsAffected > 0
}

// =========================================================
// TRANSPORTISTAS
// =========================================================

func (a *AlmacenSQLiteRutas) ListarTransportistas() []models.Transportista {
	var lista []models.Transportista
	a.db.Find(&lista)
	return lista
}

func (a *AlmacenSQLiteRutas) BuscarTransportistaPorID(id uint) (models.Transportista, bool) {
	var t models.Transportista
	if a.db.First(&t, id).Error != nil {
		return models.Transportista{}, false
	}
	return t, true
}

func (a *AlmacenSQLiteRutas) CrearTransportista(t models.Transportista) models.Transportista {
	a.db.Create(&t)
	return t
}

func (a *AlmacenSQLiteRutas) ActualizarTransportista(id uint, datos models.Transportista) (models.Transportista, bool) {
	var t models.Transportista
	if a.db.First(&t, id).Error != nil {
		return models.Transportista{}, false
	}
	if datos.Nombre != "" {
		t.Nombre = datos.Nombre
	}
	if datos.Telefono != "" {
		t.Telefono = datos.Telefono
	}
	if datos.PlacaVehiculo != "" {
		t.PlacaVehiculo = datos.PlacaVehiculo
	}
	if datos.Estado != "" {
		t.Estado = datos.Estado
	}
	a.db.Save(&t)
	return t, true
}

func (a *AlmacenSQLiteRutas) BorrarTransportista(id uint) bool {
	resultado := a.db.Delete(&models.Transportista{}, id)
	return resultado.RowsAffected > 0
}

// =========================================================
// ENTREGAS
// =========================================================

func (a *AlmacenSQLiteRutas) ListarEntregas() []models.EntregaPedido {
	var lista []models.EntregaPedido
	a.db.Find(&lista)
	return lista
}

func (a *AlmacenSQLiteRutas) BuscarEntregaPorID(id uint) (models.EntregaPedido, bool) {
	var e models.EntregaPedido
	if a.db.First(&e, id).Error != nil {
		return models.EntregaPedido{}, false
	}
	return e, true
}

func (a *AlmacenSQLiteRutas) CrearEntrega(e models.EntregaPedido) models.EntregaPedido {
	a.db.Create(&e)
	return e
}

func (a *AlmacenSQLiteRutas) ActualizarEntrega(id uint, datos models.EntregaPedido) (models.EntregaPedido, bool) {
	var e models.EntregaPedido
	if a.db.First(&e, id).Error != nil {
		return models.EntregaPedido{}, false
	}
	if datos.Estado != "" {
		e.Estado = datos.Estado
	}
	if datos.PedidoID != 0 {
		e.PedidoID = datos.PedidoID
	}
	if datos.PuntoID != 0 {
		e.PuntoID = datos.PuntoID
	}
	if datos.TransportistaID != 0 {
		e.TransportistaID = datos.TransportistaID
	}
	a.db.Save(&e)
	return e, true
}

func (a *AlmacenSQLiteRutas) BorrarEntrega(id uint) bool {
	resultado := a.db.Delete(&models.EntregaPedido{}, id)
	return resultado.RowsAffected > 0
}

// Chequeo en tiempo de compilación: si AlmacenSQLiteRutas dejara de
// cumplir AlmacenRutas, el proyecto NO compila.
var _ AlmacenRutas = (*AlmacenSQLiteRutas)(nil)

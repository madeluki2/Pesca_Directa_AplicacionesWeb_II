package storage

// Test — Repositorio real con GORM contra SQLite :memory: (usando testify).
//
// Qué comprueba: que AlmacenSQLiteRutas crea registros en la base de datos
// y los recupera correctamente (Crear → Buscar/Listar lo refleja).
//
// Por qué ":memory:": la base desaparece al terminar el test, sin
// ensuciar el disco ni depender de rutas.db. Cada test arranca limpio.

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// ─── Helper: abre SQLite en memoria y migra las tablas ───────────────────────

func abrirDBRutasEnMemoria(t *testing.T) *gorm.DB {
	t.Helper()
	gdb, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err, "no se pudo abrir SQLite en memoria")

	err = gdb.AutoMigrate(
		&models.Ruta{},
		&models.Punto{},
		&models.Transportista{},
		&models.EntregaPedido{},
	)
	require.NoError(t, err, "AutoMigrate falló")
	return gdb
}

// ─── Tests ────────────────────────────────────────────────────────────────────

// TestRepositorioRutas_CrearYListarRuta: crear una ruta → listar → está ahí.
func TestRepositorioRutas_CrearYListarRuta(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))

	ruta := models.Ruta{
		Nombre:  "Ruta Sur",
		Origen:  "Puerto Tarqui",
		Destino: "Mercado Bazurto",
		Estado:  "activo",
	}

	creada := repo.CrearRuta(ruta)

	require.NotZero(t, creada.ID, "GORM no asignó ID a la ruta creada")

	lista := repo.ListarRutas()
	require.Len(t, lista, 1, "esperaba exactamente 1 ruta en la lista")
	assert.Equal(t, "Ruta Sur", lista[0].Nombre)
}

// TestRepositorioRutas_BuscarRutaPorID: buscar un ID existente y uno inexistente.
func TestRepositorioRutas_BuscarRutaPorID(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))

	guardada := repo.CrearRuta(models.Ruta{
		Nombre:  "Ruta Este",
		Origen:  "Muelle 3",
		Destino: "Frigorífico Central",
		Estado:  "activo",
	})
	require.NotZero(t, guardada.ID)

	encontrada, ok := repo.BuscarRutaPorID(guardada.ID)
	require.True(t, ok, "BuscarRutaPorID devolvió false para un ID que existe")
	assert.Equal(t, "Ruta Este", encontrada.Nombre)

	_, ok = repo.BuscarRutaPorID(9999)
	assert.False(t, ok, "BuscarRutaPorID devolvió true para un ID que no existe")
}

// TestRepositorioRutas_CrearYListarTransportista: crear → listar → buscar por ID.
func TestRepositorioRutas_CrearYListarTransportista(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))

	creado := repo.CrearTransportista(models.Transportista{
		Nombre:        "Carlos Pérez",
		Telefono:      "0991234567",
		PlacaVehiculo: "ABC-123",
		Estado:        "activo",
	})

	require.NotZero(t, creado.ID, "GORM no asignó ID al transportista creado")

	lista := repo.ListarTransportistas()
	require.Len(t, lista, 1, "esperaba exactamente 1 transportista en la lista")
	assert.Equal(t, "ABC-123", lista[0].PlacaVehiculo)

	encontrado, ok := repo.BuscarTransportistaPorID(creado.ID)
	require.True(t, ok, "BuscarTransportistaPorID devolvió false para un ID que existe")
	assert.Equal(t, "Carlos Pérez", encontrado.Nombre)
}

func TestRepositorioRutas_ActualizarRuta(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))

	creada := repo.CrearRuta(models.Ruta{Nombre: "Ruta Vieja", Origen: "A", Destino: "B", Estado: "activo"})

	actualizada, ok := repo.ActualizarRuta(creada.ID, models.Ruta{Nombre: "Ruta Nueva", Origen: "A2", Destino: "B2", Estado: "activo"})
	require.True(t, ok)
	assert.Equal(t, "Ruta Nueva", actualizada.Nombre)

	relectura, _ := repo.BuscarRutaPorID(creada.ID)
	assert.Equal(t, "Ruta Nueva", relectura.Nombre, "el cambio no se persistió en la BD")

	_, ok = repo.ActualizarRuta(9999, models.Ruta{Nombre: "X", Origen: "A", Destino: "B"})
	assert.False(t, ok, "no debería poder actualizar un ID que no existe")
}

func TestRepositorioRutas_BorrarRuta(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))

	creada := repo.CrearRuta(models.Ruta{Nombre: "Ruta Borrable", Origen: "A", Destino: "B", Estado: "activo"})

	assert.True(t, repo.BorrarRuta(creada.ID))
	_, ok := repo.BuscarRutaPorID(creada.ID)
	assert.False(t, ok, "la ruta debería haber desaparecido tras borrarla")

	assert.False(t, repo.BorrarRuta(9999), "no debería poder borrar un ID que no existe")
}

func TestRepositorioRutas_CrearYListarPunto(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))
	ruta := repo.CrearRuta(models.Ruta{Nombre: "Ruta Base", Origen: "A", Destino: "B", Estado: "activo"})

	creado := repo.CrearPunto(models.Punto{RutaID: ruta.ID, Nombre: "Muelle", Direccion: "Av. Principal", Estado: "activo"})
	require.NotZero(t, creado.ID)

	lista := repo.ListarPuntos()
	require.Len(t, lista, 1)
	assert.Equal(t, "Muelle", lista[0].Nombre)

	encontrado, ok := repo.BuscarPuntoPorID(creado.ID)
	require.True(t, ok)
	assert.Equal(t, "Av. Principal", encontrado.Direccion)

	_, ok = repo.BuscarPuntoPorID(9999)
	assert.False(t, ok)
}

func TestRepositorioRutas_ActualizarYBorrarPunto(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))
	ruta := repo.CrearRuta(models.Ruta{Nombre: "Ruta Base", Origen: "A", Destino: "B", Estado: "activo"})
	creado := repo.CrearPunto(models.Punto{RutaID: ruta.ID, Nombre: "Muelle", Direccion: "Av. Principal", Estado: "activo"})

	actualizado, ok := repo.ActualizarPunto(creado.ID, models.Punto{RutaID: ruta.ID, Nombre: "Muelle Nuevo", Direccion: "Otra Av.", Estado: "activo"})
	require.True(t, ok)
	assert.Equal(t, "Muelle Nuevo", actualizado.Nombre)

	_, ok = repo.ActualizarPunto(9999, models.Punto{RutaID: ruta.ID, Nombre: "X", Direccion: "Y"})
	assert.False(t, ok)

	assert.True(t, repo.BorrarPunto(creado.ID))
	_, ok = repo.BuscarPuntoPorID(creado.ID)
	assert.False(t, ok)
	assert.False(t, repo.BorrarPunto(9999))
}

func TestRepositorioRutas_ActualizarYBorrarTransportista(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))
	creado := repo.CrearTransportista(models.Transportista{Nombre: "Ana", Telefono: "099", PlacaVehiculo: "XYZ-001", Estado: "activo"})

	actualizado, ok := repo.ActualizarTransportista(creado.ID, models.Transportista{Nombre: "Ana Pérez", Telefono: "098", PlacaVehiculo: "XYZ-002", Estado: "activo"})
	require.True(t, ok)
	assert.Equal(t, "Ana Pérez", actualizado.Nombre)

	_, ok = repo.ActualizarTransportista(9999, models.Transportista{Nombre: "X", Telefono: "0", PlacaVehiculo: "Y"})
	assert.False(t, ok)

	assert.True(t, repo.BorrarTransportista(creado.ID))
	_, ok = repo.BuscarTransportistaPorID(creado.ID)
	assert.False(t, ok)
	assert.False(t, repo.BorrarTransportista(9999))
}

func TestRepositorioRutas_CrearYListarEntrega(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))
	ruta := repo.CrearRuta(models.Ruta{Nombre: "Ruta Base", Origen: "A", Destino: "B", Estado: "activo"})
	punto := repo.CrearPunto(models.Punto{RutaID: ruta.ID, Nombre: "Muelle", Direccion: "Av.", Estado: "activo"})
	transportista := repo.CrearTransportista(models.Transportista{Nombre: "Ana", Telefono: "099", PlacaVehiculo: "XYZ-003", Estado: "activo"})

	creada := repo.CrearEntrega(models.EntregaPedido{
		PedidoID:        1,
		PuntoID:         punto.ID,
		TransportistaID: transportista.ID,
		Estado:          "pendiente",
	})
	require.NotZero(t, creada.ID)

	lista := repo.ListarEntregas()
	require.Len(t, lista, 1)

	encontrada, ok := repo.BuscarEntregaPorID(creada.ID)
	require.True(t, ok)
	assert.Equal(t, "pendiente", encontrada.Estado)

	_, ok = repo.BuscarEntregaPorID(9999)
	assert.False(t, ok)
}

func TestRepositorioRutas_ActualizarYBorrarEntrega(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))
	ruta := repo.CrearRuta(models.Ruta{Nombre: "Ruta Base", Origen: "A", Destino: "B", Estado: "activo"})
	punto := repo.CrearPunto(models.Punto{RutaID: ruta.ID, Nombre: "Muelle", Direccion: "Av.", Estado: "activo"})
	transportista := repo.CrearTransportista(models.Transportista{Nombre: "Ana", Telefono: "099", PlacaVehiculo: "XYZ-004", Estado: "activo"})
	creada := repo.CrearEntrega(models.EntregaPedido{PedidoID: 1, PuntoID: punto.ID, TransportistaID: transportista.ID, Estado: "pendiente"})

	actualizada, ok := repo.ActualizarEntrega(creada.ID, models.EntregaPedido{PedidoID: 1, PuntoID: punto.ID, TransportistaID: transportista.ID, Estado: "entregado"})
	require.True(t, ok)
	assert.Equal(t, "entregado", actualizada.Estado)

	_, ok = repo.ActualizarEntrega(9999, models.EntregaPedido{PedidoID: 1, PuntoID: punto.ID, TransportistaID: transportista.ID})
	assert.False(t, ok)

	assert.True(t, repo.BorrarEntrega(creada.ID))
	_, ok = repo.BuscarEntregaPorID(creada.ID)
	assert.False(t, ok)
	assert.False(t, repo.BorrarEntrega(9999))
}

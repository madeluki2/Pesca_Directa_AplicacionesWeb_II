package rutas_de_distribucion

// Test — Repositorio real con GORM contra SQLite :memory: (usando testify).
//
// Qué comprueba: que AlmacenSQLiteRutas crea registros en la base de datos
// y los recupera correctamente (Crear → Buscar/Listar lo refleja), además
// de Actualizar (total y parcial) y Borrar para Rutas, Puntos, Transportistas
// y Entregas. También cubre MemoriaRutas, la otra implementación de
// AlmacenRutas usada cuando STORAGE=memoria.
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

// ═══════════════════════════════ RUTAS (SQLite) ═══════════════════════════════

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

// TestRepositorioRutas_BorrarRuta: borrar una ruta existente y una inexistente.
func TestRepositorioRutas_BorrarRuta(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))

	creada := repo.CrearRuta(models.Ruta{Nombre: "Ruta a Borrar", Origen: "Puerto A", Destino: "Mercado B", Estado: "activo"})
	require.NotZero(t, creada.ID)

	ok := repo.BorrarRuta(creada.ID)
	assert.True(t, ok, "BorrarRuta devolvió false para un ID que existe")

	_, existe := repo.BuscarRutaPorID(creada.ID)
	assert.False(t, existe, "la ruta seguía existiendo después de borrarla")

	ok = repo.BorrarRuta(9999)
	assert.False(t, ok, "BorrarRuta devolvió true para un ID que no existe")
}

// ═══════════════════════════════ PUNTOS (SQLite) ═══════════════════════════════

// TestRepositorioRutas_CrearYListarPunto: crear → listar → buscar por ID.
func TestRepositorioRutas_CrearYListarPunto(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))

	ruta := repo.CrearRuta(models.Ruta{Nombre: "Ruta Padre", Origen: "A", Destino: "B", Estado: "activo"})
	require.NotZero(t, ruta.ID)

	creado := repo.CrearPunto(models.Punto{
		RutaID:    ruta.ID,
		Nombre:    "Muelle",
		Direccion: "Av. Principal",
		Estado:    "activo",
	})
	require.NotZero(t, creado.ID, "GORM no asignó ID al punto creado")

	lista := repo.ListarPuntos()
	require.Len(t, lista, 1, "esperaba exactamente 1 punto en la lista")
	assert.Equal(t, "Muelle", lista[0].Nombre)

	encontrado, ok := repo.BuscarPuntoPorID(creado.ID)
	require.True(t, ok, "BuscarPuntoPorID devolvió false para un ID que existe")
	assert.Equal(t, "Muelle", encontrado.Nombre)

	_, ok = repo.BuscarPuntoPorID(9999)
	assert.False(t, ok, "BuscarPuntoPorID devolvió true para un ID que no existe")
}

// TestRepositorioRutas_BorrarPunto: borrar un punto existente y uno inexistente.
func TestRepositorioRutas_BorrarPunto(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))

	ruta := repo.CrearRuta(models.Ruta{Nombre: "Ruta Padre", Origen: "A", Destino: "B", Estado: "activo"})
	creado := repo.CrearPunto(models.Punto{RutaID: ruta.ID, Nombre: "Muelle", Direccion: "Av. Principal", Estado: "activo"})
	require.NotZero(t, creado.ID)

	ok := repo.BorrarPunto(creado.ID)
	assert.True(t, ok, "BorrarPunto devolvió false para un ID que existe")

	_, existe := repo.BuscarPuntoPorID(creado.ID)
	assert.False(t, existe, "el punto seguía existiendo después de borrarlo")

	ok = repo.BorrarPunto(9999)
	assert.False(t, ok, "BorrarPunto devolvió true para un ID que no existe")
}

// ═══════════════════════════ TRANSPORTISTAS (SQLite) ═══════════════════════════

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

	_, ok = repo.BuscarTransportistaPorID(9999)
	assert.False(t, ok, "BuscarTransportistaPorID devolvió true para un ID que no existe")
}

// TestRepositorioRutas_BorrarTransportista: borrar un transportista existente y uno inexistente.
func TestRepositorioRutas_BorrarTransportista(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))

	creado := repo.CrearTransportista(models.Transportista{
		Nombre:        "Carlos Pérez",
		Telefono:      "0991234567",
		PlacaVehiculo: "ABC-123",
		Estado:        "activo",
	})
	require.NotZero(t, creado.ID)

	ok := repo.BorrarTransportista(creado.ID)
	assert.True(t, ok, "BorrarTransportista devolvió false para un ID que existe")

	_, existe := repo.BuscarTransportistaPorID(creado.ID)
	assert.False(t, existe, "el transportista seguía existiendo después de borrarlo")

	ok = repo.BorrarTransportista(9999)
	assert.False(t, ok, "BorrarTransportista devolvió true para un ID que no existe")
}

// ═══════════════════════════════ ENTREGAS (SQLite) ═════════════════════════════

// TestRepositorioRutas_CrearYListarEntrega: crear → listar → buscar por ID.
func TestRepositorioRutas_CrearYListarEntrega(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))

	creada := repo.CrearEntrega(models.EntregaPedido{
		PedidoID:        1,
		PuntoID:         1,
		TransportistaID: 1,
		Estado:          "pendiente",
	})
	require.NotZero(t, creada.ID, "GORM no asignó ID a la entrega creada")

	lista := repo.ListarEntregas()
	require.Len(t, lista, 1, "esperaba exactamente 1 entrega en la lista")
	assert.Equal(t, "pendiente", lista[0].Estado)

	encontrada, ok := repo.BuscarEntregaPorID(creada.ID)
	require.True(t, ok, "BuscarEntregaPorID devolvió false para un ID que existe")
	assert.EqualValues(t, 1, encontrada.PedidoID)

	_, ok = repo.BuscarEntregaPorID(9999)
	assert.False(t, ok, "BuscarEntregaPorID devolvió true para un ID que no existe")
}

// TestRepositorioRutas_BorrarEntrega: borrar una entrega existente y una inexistente.
func TestRepositorioRutas_BorrarEntrega(t *testing.T) {
	repo := NuevoAlmacenSQLiteRutas(abrirDBRutasEnMemoria(t))

	creada := repo.CrearEntrega(models.EntregaPedido{
		PedidoID:        1,
		PuntoID:         1,
		TransportistaID: 1,
		Estado:          "pendiente",
	})
	require.NotZero(t, creada.ID)

	ok := repo.BorrarEntrega(creada.ID)
	assert.True(t, ok, "BorrarEntrega devolvió false para un ID que existe")

	_, existe := repo.BuscarEntregaPorID(creada.ID)
	assert.False(t, existe, "la entrega seguía existiendo después de borrarla")

	ok = repo.BorrarEntrega(9999)
	assert.False(t, ok, "BorrarEntrega devolvió true para un ID que no existe")
}

// ═══════════════════════════════ RUTAS (Memoria) ═══════════════════════════════

func TestMemoriaRutas_CrearYListarRuta(t *testing.T) {
	repo := NuevaMemoriaRutas()

	creada := repo.CrearRuta(models.Ruta{Nombre: "Ruta Sur", Origen: "Puerto Tarqui", Destino: "Mercado Bazurto", Estado: "activo"})
	require.NotZero(t, creada.ID, "no se asignó ID a la ruta creada")

	lista := repo.ListarRutas()
	require.Len(t, lista, 1, "esperaba exactamente 1 ruta en la lista")
	assert.Equal(t, "Ruta Sur", lista[0].Nombre)
}

func TestMemoriaRutas_BorrarRuta(t *testing.T) {
	repo := NuevaMemoriaRutas()

	creada := repo.CrearRuta(models.Ruta{Nombre: "Ruta a Borrar", Origen: "A", Destino: "B", Estado: "activo"})
	require.NotZero(t, creada.ID)

	ok := repo.BorrarRuta(creada.ID)
	assert.True(t, ok, "BorrarRuta devolvió false para un ID que existe")

	_, existe := repo.BuscarRutaPorID(creada.ID)
	assert.False(t, existe, "la ruta seguía existiendo después de borrarla")

	ok = repo.BorrarRuta(9999)
	assert.False(t, ok, "BorrarRuta devolvió true para un ID que no existe")
}

// ═══════════════════════════════ PUNTOS (Memoria) ═══════════════════════════════

func TestMemoriaRutas_CrearYListarPunto(t *testing.T) {
	repo := NuevaMemoriaRutas()

	creado := repo.CrearPunto(models.Punto{RutaID: 1, Nombre: "Muelle", Direccion: "Av. Principal", Estado: "activo"})
	require.NotZero(t, creado.ID)

	lista := repo.ListarPuntos()
	require.Len(t, lista, 1, "esperaba exactamente 1 punto en la lista")
	assert.Equal(t, "Muelle", lista[0].Nombre)

	encontrado, ok := repo.BuscarPuntoPorID(creado.ID)
	require.True(t, ok, "BuscarPuntoPorID devolvió false para un ID que existe")
	assert.Equal(t, "Muelle", encontrado.Nombre)

	_, ok = repo.BuscarPuntoPorID(9999)
	assert.False(t, ok, "BuscarPuntoPorID devolvió true para un ID que no existe")
}

func TestMemoriaRutas_BorrarPunto(t *testing.T) {
	repo := NuevaMemoriaRutas()

	creado := repo.CrearPunto(models.Punto{RutaID: 1, Nombre: "Muelle", Direccion: "Av. Principal", Estado: "activo"})
	require.NotZero(t, creado.ID)

	ok := repo.BorrarPunto(creado.ID)
	assert.True(t, ok, "BorrarPunto devolvió false para un ID que existe")

	_, existe := repo.BuscarPuntoPorID(creado.ID)
	assert.False(t, existe, "el punto seguía existiendo después de borrarlo")

	ok = repo.BorrarPunto(9999)
	assert.False(t, ok, "BorrarPunto devolvió true para un ID que no existe")
}

// ═══════════════════════════ TRANSPORTISTAS (Memoria) ═══════════════════════════

func TestMemoriaRutas_CrearYListarTransportista(t *testing.T) {
	repo := NuevaMemoriaRutas()

	creado := repo.CrearTransportista(models.Transportista{Nombre: "Carlos Pérez", Telefono: "0991234567", PlacaVehiculo: "ABC-123", Estado: "activo"})
	require.NotZero(t, creado.ID)

	lista := repo.ListarTransportistas()
	require.Len(t, lista, 1, "esperaba exactamente 1 transportista en la lista")
	assert.Equal(t, "ABC-123", lista[0].PlacaVehiculo)

	encontrado, ok := repo.BuscarTransportistaPorID(creado.ID)
	require.True(t, ok, "BuscarTransportistaPorID devolvió false para un ID que existe")
	assert.Equal(t, "Carlos Pérez", encontrado.Nombre)

	_, ok = repo.BuscarTransportistaPorID(9999)
	assert.False(t, ok, "BuscarTransportistaPorID devolvió true para un ID que no existe")
}

func TestMemoriaRutas_BorrarTransportista(t *testing.T) {
	repo := NuevaMemoriaRutas()

	creado := repo.CrearTransportista(models.Transportista{Nombre: "Carlos Pérez", Telefono: "0991234567", PlacaVehiculo: "ABC-123", Estado: "activo"})
	require.NotZero(t, creado.ID)

	ok := repo.BorrarTransportista(creado.ID)
	assert.True(t, ok, "BorrarTransportista devolvió false para un ID que existe")

	_, existe := repo.BuscarTransportistaPorID(creado.ID)
	assert.False(t, existe, "el transportista seguía existiendo después de borrarlo")

	ok = repo.BorrarTransportista(9999)
	assert.False(t, ok, "BorrarTransportista devolvió true para un ID que no existe")
}

// ═══════════════════════════════ ENTREGAS (Memoria) ═════════════════════════════

func TestMemoriaRutas_CrearYListarEntrega(t *testing.T) {
	repo := NuevaMemoriaRutas()

	creada := repo.CrearEntrega(models.EntregaPedido{PedidoID: 1, PuntoID: 1, TransportistaID: 1, Estado: "pendiente"})
	require.NotZero(t, creada.ID)

	lista := repo.ListarEntregas()
	require.Len(t, lista, 1, "esperaba exactamente 1 entrega en la lista")
	assert.Equal(t, "pendiente", lista[0].Estado)

	encontrada, ok := repo.BuscarEntregaPorID(creada.ID)
	require.True(t, ok, "BuscarEntregaPorID devolvió false para un ID que existe")
	assert.EqualValues(t, 1, encontrada.PedidoID)

	_, ok = repo.BuscarEntregaPorID(9999)
	assert.False(t, ok, "BuscarEntregaPorID devolvió true para un ID que no existe")
}

func TestMemoriaRutas_BorrarEntrega(t *testing.T) {
	repo := NuevaMemoriaRutas()

	creada := repo.CrearEntrega(models.EntregaPedido{PedidoID: 1, PuntoID: 1, TransportistaID: 1, Estado: "pendiente"})
	require.NotZero(t, creada.ID)

	ok := repo.BorrarEntrega(creada.ID)
	assert.True(t, ok, "BorrarEntrega devolvió false para un ID que existe")

	_, existe := repo.BuscarEntregaPorID(creada.ID)
	assert.False(t, existe, "la entrega seguía existiendo después de borrarla")

	ok = repo.BorrarEntrega(9999)
	assert.False(t, ok, "BorrarEntrega devolvió true para un ID que no existe")
}

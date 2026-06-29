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

package storage

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

func abrirDBEnMemoria(t *testing.T) *gorm.DB {
	t.Helper()
	gdb, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err, "no se pudo abrir SQLite en memoria")
	err = gdb.AutoMigrate(
		&models.Pescador{},
		&models.Embarcacion{},
		&models.Especie{},
		&models.Captura{},
		&models.Bodega{},
		&models.Stock{},
	)
	require.NoError(t, err, "AutoMigrate falló")
	return gdb
}

// ═══════════════════════════ ESPECIES ════════════════════════════════════════

// TEST 1: Crear y listar
func TestRepositorio_CrearYListarEspecie(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creada := repo.CrearEspecie(models.Especie{
		NombreComun: "Atún", NombreCientifico: "Thunnus albacares",
		UnidadMedida: "kg", Temporada: "todo el año", Estado: true,
	})
	require.NotZero(t, creada.ID, "GORM no asignó ID")
	lista := repo.ListarEspecies()
	require.Len(t, lista, 1)
	assert.Equal(t, "Atún", lista[0].NombreComun)
}

// TEST 2: Buscar por ID existente
func TestRepositorio_BuscarEspeciePorID(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	guardada := repo.CrearEspecie(models.Especie{NombreComun: "Corvina", UnidadMedida: "kg", Estado: true})
	require.NotZero(t, guardada.ID)

	encontrada, ok := repo.BuscarEspeciePorID(guardada.ID)
	require.True(t, ok)
	assert.Equal(t, "Corvina", encontrada.NombreComun)
}

// TEST 3: Buscar ID que no existe (Caso Edge / Error Path)
func TestRepositorio_BuscarEspecieNoExistente(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))

	_, ok := repo.BuscarEspeciePorID(9999)
	assert.False(t, ok, "No debería encontrar una especie con ID inexistente")
}

// TEST 4: Actualizar
func TestRepositorio_ActualizarEspecie(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creada := repo.CrearEspecie(models.Especie{NombreComun: "Atún", UnidadMedida: "kg", Estado: true})

	actualizada, ok := repo.ActualizarEspecie(creada.ID, models.Especie{NombreComun: "Atún Aleta Amarilla", UnidadMedida: "kg", Estado: true})
	require.True(t, ok)
	assert.Equal(t, "Atún Aleta Amarilla", actualizada.NombreComun)

	_, ok = repo.ActualizarEspecie(9999, models.Especie{NombreComun: "X", UnidadMedida: "kg"})
	assert.False(t, ok)
}

// TEST 5: Borrar
func TestRepositorio_BorrarEspecie(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creada := repo.CrearEspecie(models.Especie{NombreComun: "Atún", UnidadMedida: "kg", Estado: true})

	assert.False(t, repo.BorrarEspecie(9999))
	assert.True(t, repo.BorrarEspecie(creada.ID))
	assert.Len(t, repo.ListarEspecies(), 0)
}

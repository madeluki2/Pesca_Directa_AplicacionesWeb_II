package gestion_pesca

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

	// Migramos únicamente tu entidad individual
	err = gdb.AutoMigrate(&models.Especie{})
	require.NoError(t, err, "AutoMigrate de Especie falló")
	return gdb
}

// ═══════════════════════════ TESTS REPOSITORIO (ESPECIE) ═══════════════════════════

func TestRepositorio_CrearYListarEspecie(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))

	creada := repo.CrearEspecie(models.Especie{
		NombreComun:  "Corvina",
		UnidadMedida: "kg",
	})
	assert.NotZero(t, creada.ID)

	lista := repo.ListarEspecies()
	assert.Len(t, lista, 1)
	assert.Equal(t, "Corvina", lista[0].NombreComun)
}

func TestRepositorio_BuscarEspecieNoExistente(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))

	_, ok := repo.BuscarEspeciePorID(9999)
	assert.False(t, ok, "No debería encontrar una especie con ID inexistente")
}

func TestRepositorio_ActualizarEspecie(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creada := repo.CrearEspecie(models.Especie{
		NombreComun:  "Atún",
		UnidadMedida: "kg",
	})

	actualizada, ok := repo.ActualizarEspecie(creada.ID, models.Especie{
		NombreComun:  "Atún Aleta Amarilla",
		UnidadMedida: "kg",
	})
	require.True(t, ok)
	assert.Equal(t, "Atún Aleta Amarilla", actualizada.NombreComun)
}

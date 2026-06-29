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

func TestRepositorio_BuscarEspeciePorID(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	guardada := repo.CrearEspecie(models.Especie{NombreComun: "Corvina", UnidadMedida: "kg", Estado: true})
	require.NotZero(t, guardada.ID)

	encontrada, ok := repo.BuscarEspeciePorID(guardada.ID)
	require.True(t, ok)
	assert.Equal(t, "Corvina", encontrada.NombreComun)

	_, ok = repo.BuscarEspeciePorID(9999)
	assert.False(t, ok)
}

func TestRepositorio_ActualizarEspecie(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creada := repo.CrearEspecie(models.Especie{NombreComun: "Atún", UnidadMedida: "kg", Estado: true})

	actualizada, ok := repo.ActualizarEspecie(creada.ID, models.Especie{NombreComun: "Atún Aleta Amarilla", UnidadMedida: "kg", Estado: true})
	require.True(t, ok)
	assert.Equal(t, "Atún Aleta Amarilla", actualizada.NombreComun)

	_, ok = repo.ActualizarEspecie(9999, models.Especie{NombreComun: "X", UnidadMedida: "kg"})
	assert.False(t, ok)
}

func TestRepositorio_BorrarEspecie(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creada := repo.CrearEspecie(models.Especie{NombreComun: "Atún", UnidadMedida: "kg", Estado: true})

	assert.False(t, repo.BorrarEspecie(9999))
	assert.True(t, repo.BorrarEspecie(creada.ID))
	assert.Len(t, repo.ListarEspecies(), 0)
}

// ═══════════════════════════ PESCADORES ══════════════════════════════════════

func TestRepositorio_CrearYListarPescador(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creado := repo.CrearPescador(models.Pescador{
		UsuarioID: 1, Cedula: "1310045678", ExperienciaAños: 5, PuertoBase: "Tarqui", Estado: true,
	})
	require.NotZero(t, creado.ID)
	lista := repo.ListarPescadores()
	require.Len(t, lista, 1)
	assert.Equal(t, "1310045678", lista[0].Cedula)

	encontrado, ok := repo.BuscarPescadorPorID(creado.ID)
	require.True(t, ok)
	assert.Equal(t, "Tarqui", encontrado.PuertoBase)
}

func TestRepositorio_ActualizarYBorrarPescador(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creado := repo.CrearPescador(models.Pescador{UsuarioID: 1, Cedula: "123", PuertoBase: "Manta", Estado: true})

	actualizado, ok := repo.ActualizarPescador(creado.ID, models.Pescador{UsuarioID: 1, Cedula: "123", PuertoBase: "Tarqui", Estado: true})
	require.True(t, ok)
	assert.Equal(t, "Tarqui", actualizado.PuertoBase)

	_, ok = repo.ActualizarPescador(9999, models.Pescador{})
	assert.False(t, ok)

	assert.False(t, repo.BorrarPescador(9999))
	assert.True(t, repo.BorrarPescador(creado.ID))
}

// ═══════════════════════════ EMBARCACIONES ════════════════════════════════════

func TestRepositorio_CrearYListarEmbarcacion(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creada := repo.CrearEmbarcacion(models.Embarcacion{
		PescadorID: 1, Nombre: "La Esperanza", Matricula: "MB-0451", CapacidadKG: 800, Estado: true,
	})
	require.NotZero(t, creada.ID)
	lista := repo.ListarEmbarcaciones()
	require.Len(t, lista, 1)
	assert.Equal(t, "La Esperanza", lista[0].Nombre)

	encontrada, ok := repo.BuscarEmbarcacionPorID(creada.ID)
	require.True(t, ok)
	assert.Equal(t, "MB-0451", encontrada.Matricula)

	_, ok = repo.BuscarEmbarcacionPorID(9999)
	assert.False(t, ok)
}

func TestRepositorio_BorrarEmbarcacion(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creada := repo.CrearEmbarcacion(models.Embarcacion{PescadorID: 1, Nombre: "X", Matricula: "Y", CapacidadKG: 100, Estado: true})

	assert.False(t, repo.BorrarEmbarcacion(9999))
	assert.True(t, repo.BorrarEmbarcacion(creada.ID))
}

// ═══════════════════════════ BODEGAS ═════════════════════════════════════════

func TestRepositorio_CrearYListarBodega(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creada := repo.CrearBodega(models.Bodega{
		Nombre: "Bodega Central", Ubicacion: "Puerto Tarqui", CapacidadKG: 5000, Estado: true,
	})
	require.NotZero(t, creada.ID)
	lista := repo.ListarBodegas()
	require.Len(t, lista, 1)
	assert.Equal(t, "Bodega Central", lista[0].Nombre)

	encontrada, ok := repo.BuscarBodegaPorID(creada.ID)
	require.True(t, ok)
	assert.Equal(t, "Puerto Tarqui", encontrada.Ubicacion)

	_, ok = repo.BuscarBodegaPorID(9999)
	assert.False(t, ok)
}

func TestRepositorio_BorrarBodega(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creada := repo.CrearBodega(models.Bodega{Nombre: "B1", Ubicacion: "Manta", CapacidadKG: 100, Estado: true})

	assert.False(t, repo.BorrarBodega(9999))
	assert.True(t, repo.BorrarBodega(creada.ID))
}

// ═══════════════════════════ STOCKS ══════════════════════════════════════════

func TestRepositorio_CrearYListarStock(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creado := repo.CrearStock(models.Stock{
		BodegaID: 1, EspecieID: 1, CantidadKG: 120.5, FechaIngreso: "2026-06-28", Estado: true,
	})
	require.NotZero(t, creado.ID)
	lista := repo.ListarStocks()
	require.Len(t, lista, 1)
	assert.Equal(t, 120.5, lista[0].CantidadKG)

	encontrado, ok := repo.BuscarStockPorID(creado.ID)
	require.True(t, ok)
	assert.Equal(t, 1, encontrado.BodegaID)

	_, ok = repo.BuscarStockPorID(9999)
	assert.False(t, ok)
}

func TestRepositorio_BorrarStock(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creado := repo.CrearStock(models.Stock{BodegaID: 1, EspecieID: 1, CantidadKG: 50, FechaIngreso: "2026-06-28", Estado: true})

	assert.False(t, repo.BorrarStock(9999))
	assert.True(t, repo.BorrarStock(creado.ID))
}

// ═══════════════════════════ CAPTURAS ════════════════════════════════════════

func TestRepositorio_CrearYListarCaptura(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creada := repo.CrearCaptura(models.Captura{
		EmbarcacionID: 1, EspecieID: 1, Fecha: "2026-06-28",
		CantidadKG: 80, PrecioSugerido: 2.5, EstadoFrescura: "fresco",
	})
	require.NotZero(t, creada.ID)
	lista := repo.ListarCapturas()
	require.Len(t, lista, 1)
	assert.Equal(t, "fresco", lista[0].EstadoFrescura)

	encontrada, ok := repo.BuscarCapturaPorID(creada.ID)
	require.True(t, ok)
	assert.Equal(t, 80.0, encontrada.CantidadKG)

	_, ok = repo.BuscarCapturaPorID(9999)
	assert.False(t, ok)
}

func TestRepositorio_BorrarCaptura(t *testing.T) {
	repo := NuevoAlmacenSQLitePesca(abrirDBEnMemoria(t))
	creada := repo.CrearCaptura(models.Captura{EmbarcacionID: 1, EspecieID: 1, Fecha: "2026-06-28", CantidadKG: 50, EstadoFrescura: "fresco"})

	assert.False(t, repo.BorrarCaptura(9999))
	assert.True(t, repo.BorrarCaptura(creada.ID))
}

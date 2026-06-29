package storage

// Test del repositorio real con GORM contra SQLite :memory: (usando testify).
//
// Qué comprueba: que AlmacenSQLite crea registros en la base de datos
// y los recupera correctamente (Crear → Buscar/Listar lo refleja).
//
// Por qué ":memory:": la base desaparece al terminar el test, sin
// ensuciar el disco ni depender de pedidos.db. Cada test arranca limpio.

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// ─── Helper: abre SQLite en memoria y migra las tablas ───────────────────────

func abrirDBPedidosEnMemoria(t *testing.T) *gorm.DB {
	t.Helper()
	gdb, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err, "no se pudo abrir SQLite en memoria")

	err = gdb.AutoMigrate(
		&models.Cliente{},
		&models.Pedido{},
		&models.DetallePedido{},
	)
	require.NoError(t, err, "AutoMigrate falló")
	return gdb
}

// ─── Tests ────────────────────────────────────────────────────────────────────

// TestRepositorioPedidos_CrearYListarCliente: crear un cliente → listar → está ahí.
func TestRepositorioPedidos_CrearYListarCliente(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBPedidosEnMemoria(t))

	creado := repo.CrearCliente(models.Cliente{
		UsuarioID:     1,
		TipoCliente:   "restaurante",
		NombreNegocio: "Sushi Koi",
		Direccion:     "Av. Flavio Reyes",
		Telefono:      "0991234567",
		Estado:        "activo",
	})

	require.NotZero(t, creado.ID, "GORM no asignó ID al cliente creado")

	lista := repo.ListarClientes()
	require.Len(t, lista, 1, "esperaba exactamente 1 cliente en la lista")
	assert.Equal(t, "Sushi Koi", lista[0].NombreNegocio)
}

// TestRepositorioPedidos_BuscarClientePorID: buscar un ID existente y uno inexistente.
func TestRepositorioPedidos_BuscarClientePorID(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBPedidosEnMemoria(t))

	guardado := repo.CrearCliente(models.Cliente{
		UsuarioID:     1,
		TipoCliente:   "mayorista",
		NombreNegocio: "Distribuidora El Puerto",
		Direccion:     "Calle 10 de Agosto",
		Telefono:      "0997654321",
		Estado:        "activo",
	})
	require.NotZero(t, guardado.ID)

	encontrado, ok := repo.BuscarClientePorID(guardado.ID)
	require.True(t, ok, "BuscarClientePorID devolvió false para un ID que existe")
	assert.Equal(t, "Distribuidora El Puerto", encontrado.NombreNegocio)

	_, ok = repo.BuscarClientePorID(9999)
	assert.False(t, ok, "BuscarClientePorID devolvió true para un ID que no existe")
}

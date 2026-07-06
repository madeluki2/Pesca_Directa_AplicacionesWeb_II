package storage

import (
	"fmt"
	"log"
	"time"

	_ "github.com/glebarez/go-sqlite" // Driver tradicional database/sql para SQLite por si usan sqlc
	"github.com/glebarez/sqlite"      // Dialector GORM para SQLite
	"gorm.io/driver/postgres"         // Dialector GORM para PostgreSQL
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// Recursos agrupa todo lo que la capa de almacenamiento expone a la aplicación de Pesca.
type Recursos struct {
	Pesca        AlmacenPesca
	Pedidos      Almacen
	Rutas        AlmacenRutas
	Usuarios     UserRepository // Repositorio de usuarios para el AuthService
	BackendUsado string
	Cerrar       func() error
}

// Inicializar centraliza todo el plumbing de almacenamiento siguiendo el patrón del profesor.
func Inicializar(driver, dsn, rutaDB, storageCfg string) (*Recursos, error) {
	// ── CASO A: ALMACENAMIENTO VOLÁTIL EN MEMORIA ────────────────────────
	if storageCfg == "memoria" {
		m := NewMemoria()
		m.Seed()

		cerrarVacio := func() error { return nil }

		return &Recursos{
			Pesca:        NuevaMemoriaPesca(),
			Pedidos:      m,
			Rutas:        NuevaMemoriaRutas(),
			Usuarios:     NewUsuarioGORM(nil), // Instancia de respaldo sin BD real
			BackendUsado: "Memoria (datos volátiles)",
			Cerrar:       cerrarVacio,
		}, nil
	}

	// ── CASO B: INFRAESTRUCTURA PERSISTENTE (GORM) ───────────────────────
	// 1. GORM abre según el motor (postgres o sqlite) y maneja la lógica de reintentos
	gdb, err := abrirGorm(driver, dsn, rutaDB)
	if err != nil {
		return nil, err
	}

	// El Factory se encarga de migrar de golpe todo el esquema de Pesca-Directa Tarqui
	if err := gdb.AutoMigrate(
		&models.Usuario{},
		&models.Pescador{},
		&models.Embarcacion{},
		&models.Especie{},
		&models.Captura{},
		&models.Bodega{},
		&models.Stock{},
		&models.Cliente{},
		&models.Pedido{},
		&models.DetallePedido{},
		&models.Ruta{},
		&models.Punto{},
		&models.Transportista{},
		&models.EntregaPedido{},
	); err != nil {
		return nil, fmt.Errorf("falló AutoMigrate en Factory: %w", err)
	}

	// 2. Inicializamos tus repositorios reales usando las funciones de tu proyecto
	almacenPesca := NuevoAlmacenSQLitePesca(gdb)
	almacenPedidos := NuevoAlmacenSQLite(gdb)
	almacenRutas := NuevoAlmacenSQLiteRutas(gdb)
	usuarioRepo := NewUsuarioGORM(gdb)

	// 3. Cierre ordenado de conexiones
	cerrar := func() error {
		sqlDB, err := gdb.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}

	backendInfo := "GORM + PostgreSQL"
	if driver == "sqlite" {
		backendInfo = fmt.Sprintf("GORM + SQLite ( %s )", rutaDB)
	}

	return &Recursos{
		Pesca:        almacenPesca,
		Pedidos:      almacenPedidos,
		Rutas:        almacenRutas,
		Usuarios:     usuarioRepo,
		BackendUsado: backendInfo,
		Cerrar:       cerrar,
	}, nil
}

// abrirGorm elige el Dialector según el driver y reintenta la conexión de forma robusta
func abrirGorm(driver, dsn, rutaDB string) (*gorm.DB, error) {
	switch driver {
	case "postgres":
		var gdb *gorm.DB
		var err error
		for intento := 1; intento <= 10; intento++ {
			gdb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				return gdb, nil
			}
			log.Printf("PostgreSQL (Pesca API) no está listo (intento %d/10): %v", intento, err)
			time.Sleep(2 * time.Second)
		}
		return nil, fmt.Errorf("conectar a PostgreSQL tras reintentos: %w", err)
	default: // "sqlite"
		gdb, err := gorm.Open(sqlite.Open(rutaDB), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("abrir SQLite: %w", err)
		}
		return gdb, nil
	}
}

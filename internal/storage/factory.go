package storage

import (
	"fmt"
	"log"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// Recursos agrupa todo lo que la capa de almacenamiento expone a la aplicación.
// Los campos son INTERFACES (no tipos concretos) para que tanto MemoriaPesca
// como AlmacenSQLitePesca puedan asignarse sin error de tipo.
type Recursos struct {
	Pesca        AlmacenPesca   // interfaz — acepta MemoriaPesca o AlmacenSQLitePesca
	Pedidos      Almacen        // interfaz — acepta Memoria o AlmacenSQLite
	Rutas        AlmacenRutas   // interfaz — acepta MemoriaRutas o AlmacenSQLiteRutas
	Usuarios     UserRepository // siempre GORM
	BackendUsado string
	Cerrar       func() error
}

// Inicializar centraliza TODO el plumbing de almacenamiento (patrón Factory).
func Inicializar(driver, dsn, rutaDB, backend string) (*Recursos, error) {
	// 1. GORM abre la BD, migra el esquema y siembra datos iniciales.
	gdb, err := abrirGorm(driver, dsn, rutaDB)
	if err != nil {
		return nil, err
	}

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
		return nil, fmt.Errorf("AutoMigrate: %w", err)
	}

	// 2. Elegir backend según variable STORAGE.
	// Las variables son de tipo interfaz → cualquier implementación es válida.
	var almacenPesca AlmacenPesca
	var almacenPedidos Almacen
	var almacenRutas AlmacenRutas
	backendUsado := "gorm"

	if backend == "memoria" {
		almacenPesca = NuevaMemoriaPesca() // *MemoriaPesca implementa AlmacenPesca ✓
		m := NewMemoria()
		m.Seed()
		almacenPedidos = m                 // *Memoria implementa Almacen ✓
		almacenRutas = NuevaMemoriaRutas() // *MemoriaRutas implementa AlmacenRutas ✓
		backendUsado = "memoria"
	} else {
		almacenPesca = NuevoAlmacenSQLitePesca(gdb)
		almacenPedidos = NuevoAlmacenSQLite(gdb)
		almacenRutas = NuevoAlmacenSQLiteRutas(gdb)
	}

	// 3. Usuarios siempre en GORM.
	usuarios := NewUsuarioGORM(gdb)

	// 4. Cierre ordenado.
	cerrar := func() error {
		sqlDB, err := gdb.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}

	return &Recursos{
		Pesca:        almacenPesca,
		Pedidos:      almacenPedidos,
		Rutas:        almacenRutas,
		Usuarios:     usuarios,
		BackendUsado: backendUsado,
		Cerrar:       cerrar,
	}, nil
}

// abrirGorm elige el Dialector según el driver y abre la conexión.
// Para PostgreSQL reintenta 10 veces con 2s de espera entre intentos.
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
			log.Printf("PostgreSQL no está listo (intento %d/10): %v", intento, err)
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

package storage

import (
	"fmt"
	"log"
	"time"

	"github.com/glebarez/sqlite" // dialector GORM para SQLite (pure-Go)
	"gorm.io/driver/postgres"    // dialector GORM para PostgreSQL
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// Recursos agrupa todo lo que la capa de almacenamiento expone a main.go.
type Recursos struct {
	AlmacenPesca   AlmacenPesca
	AlmacenPedidos Almacen
	AlmacenRutas   AlmacenRutas
	Usuarios       UserRepository
	BackendUsado   string
	Cerrar         func() error
}

// Inicializar centraliza el plumbing de almacenamiento (patron Factory).
// El motor (sqlite en local, postgres en Docker) se elige por config; el
// backend de pesca/pedidos (memoria o gorm) se elige igual que antes.
func Inicializar(driver, dsn, rutaDB, backend string) (*Recursos, error) {
	// 1. GORM es el dueño del esquema.
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

	// 2. Elegir el backend de pesca/pedidos (rutas siempre vive en GORM,
	//    igual que en tu main.go original).
	var almacenPesca AlmacenPesca
	var almacenPedidos Almacen
	almacenRutas := NuevoAlmacenSQLiteRutas(gdb)
	backendUsado := "gorm"

	switch backend {
	case "memoria":
		almacenPesca = NuevaMemoriaPesca()
		m := NewMemoria()
		m.Seed()
		almacenPedidos = m
		backendUsado = "memoria"
	default:
		almacenPesca = NuevoAlmacenSQLitePesca(gdb)
		almacenPedidos = NuevoAlmacenSQLite(gdb)
	}

	// 3. Usuarios viven siempre en GORM.
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
		AlmacenPesca:   almacenPesca,
		AlmacenPedidos: almacenPedidos,
		AlmacenRutas:   almacenRutas,
		Usuarios:       usuarios,
		BackendUsado:   backendUsado,
		Cerrar:         cerrar,
	}, nil
}

// abrirGorm elige el Dialector segun el driver. Para PostgreSQL reintenta
// unos segundos: dentro de docker compose la base puede tardar en aceptar
// conexiones aunque el contenedor ya este arriba.
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
			log.Printf("PostgreSQL no esta listo (intento %d/10): %v", intento, err)
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

package storage

import (
	"fmt"
	"log"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	pescaStorage "Pesca_Directa_AplicacionesWeb_II/internal/storage/gestion_pesca"
)

type Recursos struct {
	Pesca        pescaStorage.AlmacenPesca
	Pedidos      Almacen
	Rutas        AlmacenRutas
	Usuarios     UserRepository
	BackendUsado string
	Cerrar       func() error
}

func Inicializar(driver, dsn, rutaDB, backend string) (*Recursos, error) {
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

	var almacenPedidos Almacen
	var almacenRutas AlmacenRutas
	backendUsado := "gorm"

	if backend == "memoria" {
		m := NewMemoria()
		m.Seed()
		almacenPedidos = m
		almacenRutas = NuevaMemoriaRutas()
		backendUsado = "memoria"
	} else {
		almacenPedidos = NuevoAlmacenSQLite(gdb)
		almacenRutas = NuevoAlmacenSQLiteRutas(gdb)
	}

	cerrar := func() error {
		sqlDB, err := gdb.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}

	return &Recursos{
		Pesca:        pescaStorage.NuevoAlmacenPesca(gdb, backend),
		Pedidos:      almacenPedidos,
		Rutas:        almacenRutas,
		Usuarios:     NewUsuarioGORM(gdb),
		BackendUsado: backendUsado,
		Cerrar:       cerrar,
	}, nil
}

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
	default:
		gdb, err := gorm.Open(sqlite.Open(rutaDB), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("abrir SQLite: %w", err)
		}
		return gdb, nil
	}
}

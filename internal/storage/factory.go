package storage

import (
	"fmt"
	"log"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	pedidosStorage "Pesca_Directa_AplicacionesWeb_II/internal/storage/gestion_pedidos"
<<<<<<< HEAD
<<<<<<< HEAD
	pescaStorage "Pesca_Directa_AplicacionesWeb_II/internal/storage/gestion_pesca"
	rutasStorage "Pesca_Directa_AplicacionesWeb_II/internal/storage/rutas_de_distribucion"
=======
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
=======
	pescaStorage "Pesca_Directa_AplicacionesWeb_II/internal/storage/gestion_pesca"
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d
)

// Recursos agrupa los almacenes de los modulos y el repositorio compartido de usuarios.
type Recursos struct {
<<<<<<< HEAD
<<<<<<< HEAD
	Pesca        pescaStorage.AlmacenPesca
	Pedidos      pedidosStorage.Almacen
	Rutas        rutasStorage.AlmacenRutas
=======
	Pesca        AlmacenPesca
=======
	Pesca        pescaStorage.AlmacenPesca
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d
	Pedidos      pedidosStorage.Almacen
	Rutas        AlmacenRutas
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
	Usuarios     UserRepository
	BackendUsado string
	Cerrar       func() error
}

// Inicializar abre la base de datos y construye los repositorios de pesca, pedidos y rutas.
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

<<<<<<< HEAD
<<<<<<< HEAD
	var almacenRutas rutasStorage.AlmacenRutas
	backendUsado := "gorm"

	if backend == "memoria" {
		almacenRutas = rutasStorage.NuevaMemoriaRutas()
		backendUsado = "memoria"
	} else {
		almacenRutas = rutasStorage.NuevoAlmacenSQLiteRutas(gdb)
=======
	var almacenRutas AlmacenRutas
	var almacenPesca AlmacenPesca
=======
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d
	backendUsado := "gorm"
	almacenRutas := AlmacenRutas(NuevoAlmacenSQLiteRutas(gdb))
	if backend == "memoria" {
		almacenRutas = NuevaMemoriaRutas()
		backendUsado = "memoria"
<<<<<<< HEAD
	} else {
		almacenRutas = NuevoAlmacenSQLiteRutas(gdb)
		almacenPesca = NuevoAlmacenSQLitePesca(gdb)
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
=======
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d
	}

	cerrar := func() error {
		sqlDB, err := gdb.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}

	return &Recursos{
<<<<<<< HEAD
<<<<<<< HEAD
		Pesca:        pescaStorage.NuevoAlmacenPesca(gdb, backend),
=======
		Pesca:        almacenPesca,
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
		Pedidos:      almacenPedidos,
=======
		Pesca:        pescaStorage.NuevoAlmacenPesca(gdb, backend),
		Pedidos:      pedidosStorage.NuevoAlmacenSQLite(gdb),
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d
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

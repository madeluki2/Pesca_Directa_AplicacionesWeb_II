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
	pescaStorage "Pesca_Directa_AplicacionesWeb_II/internal/storage/gestion_pesca"
	rutasStorage "Pesca_Directa_AplicacionesWeb_II/internal/storage/rutas_de_distribucion"
=======
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
)

// Recursos agrupa todo lo que main.go necesita para arrancar: los tres
// almacenes (uno por módulo), el repositorio de usuarios (compartido) y
// una función para cerrar la conexión a la base de datos limpiamente.
type Recursos struct {
<<<<<<< HEAD
	Pesca        pescaStorage.AlmacenPesca
	Pedidos      pedidosStorage.Almacen
	Rutas        rutasStorage.AlmacenRutas
=======
	Pesca        AlmacenPesca
	Pedidos      pedidosStorage.Almacen
	Rutas        AlmacenRutas
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
	Usuarios     UserRepository
	BackendUsado string
	Cerrar       func() error
}

// Inicializar abre la base de datos (sqlite en local, postgres en Docker),
// corre AutoMigrate para TODOS los modelos de los 3 módulos, y arma los
// almacenes correspondientes usando una única conexión *gorm.DB inyectada
// en cada uno (Pesca, Pedidos y Rutas).
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
	backendUsado := "gorm"

	if backend == "memoria" {
		almacenRutas = NuevaMemoriaRutas()
		almacenPesca = NuevaMemoriaPesca()
		backendUsado = "memoria"
	} else {
		almacenRutas = NuevoAlmacenSQLiteRutas(gdb)
		almacenPesca = NuevoAlmacenSQLitePesca(gdb)
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
	}

	// Pedidos aún no tiene backend en memoria migrado; siempre usa GORM.
	almacenPedidos := pedidosStorage.NuevoAlmacenSQLite(gdb)

	cerrar := func() error {
		sqlDB, err := gdb.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}

	return &Recursos{
<<<<<<< HEAD
		Pesca:        pescaStorage.NuevoAlmacenPesca(gdb, backend),
=======
		Pesca:        almacenPesca,
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
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

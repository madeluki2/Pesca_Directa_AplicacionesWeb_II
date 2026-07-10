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
	. "Pesca_Directa_AplicacionesWeb_II/internal/storage/rutas_de_distribucion"
)

// Recursos agrupa todo lo que main.go necesita para arrancar.
type Recursos struct {
	Pesca        AlmacenPesca
	Pedidos      pedidosStorage.Almacen
	Rutas        AlmacenRutas
	Usuarios     UserRepository
	BackendUsado string
	Cerrar       func() error
}

// Inicializar abre la base de datos y arma los almacenes.
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

	var almacenPesca AlmacenPesca
	var almacenRutas AlmacenRutas
	backendUsado := "gorm"

	if backend == "memoria" {
		almacenPesca = NuevaMemoriaPesca()
		almacenRutas = NuevaMemoriaRutas()
		backendUsado = "memoria"
	} else {
		almacenPesca = NuevoAlmacenSQLitePesca(gdb)
		almacenRutas = NuevoAlmacenSQLiteRutas(gdb)
	}

	// Pedidos aún no tiene backend en memoria migrado; siempre usa GORM.
	almacenPedidos := pedidosStorage.NuevoAlmacenSQLite(gdb)
	usuarios := NewUsuarioGORM(gdb)

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
	default: // sqlite
		gdb, err := gorm.Open(sqlite.Open(rutaDB), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("abrir SQLite: %w", err)
		}
		return gdb, nil
	}
}

package gestion_pedidos

import (
	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"fmt"

	"gorm.io/driver/sqlite" // Asegúrate de tener este driver
	"gorm.io/gorm"
)

// Recursos ahora solo expone Pedidos
type Recursos struct {
	Pedidos      Almacen
	BackendUsado string
	Cerrar       func() error
}

func Inicializar(rutaDB string, backend string) (*Recursos, error) {
	// 1. Abrir conexión única
	gdb, err := gorm.Open(sqlite.Open(rutaDB), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al abrir BD: %w", err)
	}

	// 2. Migraciones específicas de Pedidos
	if err := gdb.AutoMigrate(
		&models.Cliente{},
		&models.Pedido{},
		&models.DetallePedido{},
	); err != nil {
		return nil, fmt.Errorf("AutoMigrate: %w", err)
	}

	var almacenPedidos Almacen
	backendUsado := "gorm"

	// 3. Inicialización exclusiva de tu rama
	if backend == "memoria" {
		// Asumiendo que NuevoAlmacenMemoria() es tu constructor de memoria
		almacenPedidos = NuevoAlmacenSQLite(gdb)
		almacenPedidos.Seed()
		backendUsado = "memoria"
	} else {
		almacenPedidos = NuevoAlmacenSQLite(gdb)
	}

	cerrar := func() error {
		sqlDB, err := gdb.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}

	return &Recursos{
		Pedidos:      almacenPedidos,
		BackendUsado: backendUsado,
		Cerrar:       cerrar,
	}, nil
}

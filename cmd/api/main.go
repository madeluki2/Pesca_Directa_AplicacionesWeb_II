package main

import (
	"log"
	"net/http"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

func main() {
	// 1. GORM abre la DB y migra las tablas automáticamente.
	//    AutoMigrate crea las tablas en pedidos.db sin borrar datos existentes.
	gdb, err := gorm.Open(sqlite.Open("pedidos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}
	if err := gdb.AutoMigrate(
		&models.Usuario{},
		&models.Cliente{},
		&models.Pedido{},
		&models.DetallePedido{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	// 2. Elegir el backend según la variable de entorno STORAGE.
	//    STORAGE=memoria → usa almacenamiento en RAM (datos volátiles)
	//    por defecto     → usa GORM + SQLite (pedidos.db)
	var almacen storage.Almacen
	switch os.Getenv("STORAGE") {
	case "memoria":
		m := storage.NewMemoria()
		m.Seed()
		almacen = m
		log.Println("Backend de almacenamiento: Memoria (datos volátiles)")
	default:
		almacen = storage.NuevoAlmacenSQLite(gdb)
		log.Println("Backend de almacenamiento: GORM + SQLite (pedidos.db)")
	}

	// 3. UsuarioRepository siempre usa GORM porque el login necesita persistencia real.
	usuarioRepo := storage.NewUsuarioGORM(gdb)

	// 4. Services con inyección de dependencias.
	authService := service.NewAuthService(usuarioRepo)
	pedidoService := service.NewPedidoService(almacen)

	// 5. Server: punto único de entrada a todos los services para los handlers.
	servidor := handlers.NewServer(pedidoService, authService)

	// 6. Router + middlewares globales.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 7. Rutas versionadas /api/v1/
	r.Route("/api/v1", func(r chi.Router) {

		// Rutas públicas — sin token
		r.Post("/auth/register", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		// Rutas protegidas — requieren header: Authorization: Bearer <token>
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			// Clientes
			r.Get("/clientes", servidor.ListarClientes)
			r.Post("/clientes", servidor.CrearCliente)
			r.Get("/clientes/{id}", servidor.ObtenerCliente)
			r.Put("/clientes/{id}", servidor.ActualizarCliente)
			r.Delete("/clientes/{id}", servidor.EliminarCliente)
			r.Patch("/clientes/{id}/tipo", servidor.CambiarTipoCliente)

			// Pedidos
			r.Get("/pedidos", servidor.ListarPedidos)
			r.Post("/pedidos", servidor.CrearPedido)
			r.Get("/pedidos/{id}", servidor.ObtenerPedido)
			r.Put("/pedidos/{id}", servidor.ActualizarPedido)
			r.Delete("/pedidos/{id}", servidor.EliminarPedido)

			// Detalles de Pedido
			r.Get("/detalles-pedido", servidor.ListarDetalles)
			r.Post("/detalles-pedido", servidor.CrearDetalle)
			r.Get("/detalles-pedido/{id}", servidor.ObtenerDetalle)
			r.Put("/detalles-pedido/{id}", servidor.ActualizarDetalle)
			r.Delete("/detalles-pedido/{id}", servidor.EliminarDetalle)
		})
	})

	log.Println("Servidor Pesca-Directa Tarqui — Gestión de Pedidos escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

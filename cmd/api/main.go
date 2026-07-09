package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/config"
	"Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	pedidosHandlers "Pesca_Directa_AplicacionesWeb_II/internal/handlers/gestion_pedidos"
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	pedidosService "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pedidos"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
	pedidosStorage "Pesca_Directa_AplicacionesWeb_II/internal/storage/gestion_pedidos"
)

func main() {
	cfg, err := config.Cargar()
	if err != nil {
		log.Fatal("error cargando configuración: ", err)
	}
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config.Config) error {
	// ── 1. Inicializar BD ───────────────────────────
	gdb, err := gorm.Open(sqlite.Open(cfg.RutaDB), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("no se pudo abrir la base de datos: %w", err)
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
		return fmt.Errorf("falló AutoMigrate: %w", err)
	}

	// ── 2. Repositorios ───────────────────────────
	usuarioRepo := storage.NewUsuarioGORM(gdb)
	almacenPesca := storage.NuevoAlmacenSQLitePesca(gdb)
	almacenRutas := storage.NuevoAlmacenSQLiteRutas(gdb)
	almacenPedidos := pedidosStorage.NuevoAlmacenSQLite(gdb)

	log.Println("Infraestructura de Datos GORM + SQLite inicializada.")

	// ── 3. Services ───────────────────────────
	authService := service.NewAuthService(usuarioRepo)
	pescaService := service.NewPescaService(almacenPesca)
	rutasService := service.NewRutasService(almacenRutas)
	pedidoService := pedidosService.NewPedidoService(almacenPedidos)

	// ── 4. Servers ───────────────────────────
	// Server principal (solo pesca y rutas)
	servidor := handlers.NewServer(pescaService, rutasService)

	// Server de pedidos (con auth)
	pedidoServer := pedidosHandlers.NewServer(pedidoService, authService)

	// ── 5. Router ───────────────────────────
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// ── 6. Rutas ───────────────────────────
	r.Route("/api/v1", func(r chi.Router) {
		// Auth
		r.Post("/auth/register", pedidoServer.Registrar)
		r.Post("/auth/login", pedidoServer.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			// Anthony — Gestión de Pesca
			r.Get("/pescadores", servidor.ListarPescadores)
			// ... resto de pesca igual

			// Madelyn — Rutas
			r.Get("/rutas", servidor.ListarRutas)
			// ... resto de rutas igual

			// Ilaria — Gestión de Pedidos
			r.Get("/clientes", pedidoServer.ListarClientes)
			r.Post("/clientes", pedidoServer.CrearCliente)
			r.Get("/clientes/{id}", pedidoServer.ObtenerCliente)
			r.Put("/clientes/{id}", pedidoServer.ActualizarCliente)
			r.Delete("/clientes/{id}", pedidoServer.EliminarCliente)
			r.Patch("/clientes/{id}/tipo", pedidoServer.CambiarTipoCliente)

			r.Get("/pedidos", pedidoServer.ListarPedidos)
			r.Post("/pedidos", pedidoServer.CrearPedido)
			r.Get("/pedidos/{id}", pedidoServer.ObtenerPedido)
			r.Put("/pedidos/{id}", pedidoServer.ActualizarPedido)
			r.Delete("/pedidos/{id}", pedidoServer.EliminarPedido)

			r.Get("/detalles-pedido", pedidoServer.ListarDetalles)
			r.Post("/detalles-pedido", pedidoServer.CrearDetalle)
			r.Get("/detalles-pedido/{id}", pedidoServer.ObtenerDetalle)
			r.Put("/detalles-pedido/{id}", pedidoServer.ActualizarDetalle)
			r.Delete("/detalles-pedido/{id}", pedidoServer.EliminarDetalle)
		})
	})

	// ── 7. Servidor HTTP ───────────────────────────
	srv := &http.Server{
		Addr:         ":" + cfg.Puerto,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Servidor Pesca-Directa Tarqui escuchando en http://localhost:" + cfg.Puerto)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("error del servidor: %v", err)
		}
	}()

	<-stop
	log.Println("Apagando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("error en shutdown: %w", err)
	}

	log.Println("Servidor detenido limpiamente.")
	return nil
}

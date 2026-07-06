// Command pesca-api arranca el servidor HTTP de Pesca-Directa Tarqui.
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
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

func main() {
	// main.go solo hace dos cosas: cargar la config y llamar a run().
	// run() devuelve error en vez de llamar log.Fatal en cada paso.
	cfg, err := config.Cargar()
	if err != nil {
		log.Fatal("error cargando configuración: ", err)
	}

	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

// run contiene toda la lógica de arranque.
// Separar run() de main() facilita los tests de integración del servidor.
func run(cfg config.Config) error {

	// ── 1. GORM: dueño del esquema ───────────────────────────────────────
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

	// ── 2. Elegir backend según cfg.Storage ──────────────────────────────
	var almacenPesca storage.AlmacenPesca
	var almacenPedidos storage.Almacen
	var almacenRutas storage.AlmacenRutas

	switch cfg.Storage {
	case "memoria":
		almacenPesca = storage.NuevaMemoriaPesca()
		m := storage.NewMemoria()
		m.Seed()
		almacenPedidos = m
		almacenRutas = storage.NuevaMemoriaRutas()
		log.Println("Backend: Memoria (datos volátiles)")
	default:
		almacenPesca = storage.NuevoAlmacenSQLitePesca(gdb)
		almacenPedidos = storage.NuevoAlmacenSQLite(gdb)
		almacenRutas = storage.NuevoAlmacenSQLiteRutas(gdb)
		log.Println("Backend: GORM + SQLite (", cfg.RutaDB, ")")
	}

	// ── 3. Services con configuración inyectada ──────────────────────────
	// El secreto JWT y la duración vienen del .env, no están hardcodeados.
	usuarioRepo := storage.NewUsuarioGORM(gdb)
	authService := service.NewAuthService(usuarioRepo,
		service.WithSecreto(cfg.JWTSecreto),
		service.WithDuracionToken(cfg.JWTDuracion),
	)
	pescaService := service.NewPescaService(almacenPesca)
	pedidoService := service.NewPedidoService(almacenPedidos)
	rutasService := service.NewRutasService(almacenRutas)

	// ── 4. Server con Deps struct ────────────────────────────────────────
	servidor := handlers.NewServer(handlers.Deps{
		Pesca:   pescaService,
		Pedidos: pedidoService,
		Rutas:   rutasService,
		Auth:    authService,
	})

	// ── 5. Router + middlewares ──────────────────────────────────────────
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// ── 6. Rutas ─────────────────────────────────────────────────────────
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			// Anthony — Gestión de Pesca
			r.Get("/pescadores", servidor.ListarPescadores)
			r.Post("/pescadores", servidor.CrearPescador)
			r.Get("/pescadores/{id}", servidor.ObtenerPescador)
			r.Put("/pescadores/{id}", servidor.ActualizarPescador)
			r.Delete("/pescadores/{id}", servidor.BorrarPescador)

			r.Get("/embarcaciones", servidor.ListarEmbarcaciones)
			r.Post("/embarcaciones", servidor.CrearEmbarcacion)
			r.Get("/embarcaciones/{id}", servidor.ObtenerEmbarcacion)
			r.Put("/embarcaciones/{id}", servidor.ActualizarEmbarcacion)
			r.Delete("/embarcaciones/{id}", servidor.BorrarEmbarcacion)

			r.Get("/especies", servidor.ListarEspecies)
			r.Post("/especies", servidor.CrearEspecie)
			r.Get("/especies/{id}", servidor.ObtenerEspecie)
			r.Put("/especies/{id}", servidor.ActualizarEspecie)
			r.Delete("/especies/{id}", servidor.BorrarEspecie)

			r.Get("/capturas", servidor.ListarCapturas)
			r.Post("/capturas", servidor.CrearCaptura)
			r.Get("/capturas/{id}", servidor.ObtenerCaptura)
			r.Put("/capturas/{id}", servidor.ActualizarCaptura)
			r.Delete("/capturas/{id}", servidor.BorrarCaptura)

			r.Get("/bodegas", servidor.ListarBodegas)
			r.Post("/bodegas", servidor.CrearBodega)
			r.Get("/bodegas/{id}", servidor.ObtenerBodega)
			r.Put("/bodegas/{id}", servidor.ActualizarBodega)
			r.Delete("/bodegas/{id}", servidor.BorrarBodega)

			r.Get("/stocks", servidor.ListarStocks)
			r.Post("/stocks", servidor.CrearStock)
			r.Get("/stocks/{id}", servidor.ObtenerStock)
			r.Put("/stocks/{id}", servidor.ActualizarStock)
			r.Delete("/stocks/{id}", servidor.BorrarStock)

			// Ilaria — Gestión de Pedidos
			r.Get("/clientes", servidor.ListarClientes)
			r.Post("/clientes", servidor.CrearCliente)
			r.Get("/clientes/{id}", servidor.ObtenerCliente)
			r.Put("/clientes/{id}", servidor.ActualizarCliente)
			r.Delete("/clientes/{id}", servidor.EliminarCliente)
			r.Patch("/clientes/{id}/tipo", servidor.CambiarTipoCliente)

			r.Get("/pedidos", servidor.ListarPedidos)
			r.Post("/pedidos", servidor.CrearPedido)
			r.Get("/pedidos/{id}", servidor.ObtenerPedido)
			r.Put("/pedidos/{id}", servidor.ActualizarPedido)
			r.Delete("/pedidos/{id}", servidor.EliminarPedido)

			r.Get("/detalles-pedido", servidor.ListarDetalles)
			r.Post("/detalles-pedido", servidor.CrearDetalle)
			r.Get("/detalles-pedido/{id}", servidor.ObtenerDetalle)
			r.Put("/detalles-pedido/{id}", servidor.ActualizarDetalle)
			r.Delete("/detalles-pedido/{id}", servidor.EliminarDetalle)

			// Madelyn — Rutas de Distribución
			r.Get("/rutas", servidor.ListarRutas)
			r.Post("/rutas", servidor.CrearRuta)
			r.Get("/rutas/{id}", servidor.ObtenerRuta)
			r.Put("/rutas/{id}", servidor.ActualizarRuta)
			r.Delete("/rutas/{id}", servidor.BorrarRuta)

			r.Get("/puntos", servidor.ListarPuntos)
			r.Post("/puntos", servidor.CrearPunto)
			r.Get("/puntos/{id}", servidor.ObtenerPunto)
			r.Put("/puntos/{id}", servidor.ActualizarPunto)
			r.Delete("/puntos/{id}", servidor.BorrarPunto)

			r.Get("/transportistas", servidor.ListarTransportistas)
			r.Post("/transportistas", servidor.CrearTransportista)
			r.Get("/transportistas/{id}", servidor.ObtenerTransportista)
			r.Put("/transportistas/{id}", servidor.ActualizarTransportista)
			r.Delete("/transportistas/{id}", servidor.BorrarTransportista)

			r.Get("/entregas", servidor.ListarEntregas)
			r.Post("/entregas", servidor.CrearEntrega)
			r.Get("/entregas/{id}", servidor.ObtenerEntrega)
			r.Put("/entregas/{id}", servidor.ActualizarEntrega)
			r.Delete("/entregas/{id}", servidor.BorrarEntrega)
		})
	})

	// ── 7. Servidor HTTP con timeouts y graceful shutdown ────────────────
	srv := &http.Server{
		Addr:         ":" + cfg.Puerto,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Canal que escucha señales de sistema (Ctrl+C, kill)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Arrancar el servidor en una goroutine para poder escuchar la señal
	go func() {
		log.Println("Servidor Pesca-Directa Tarqui escuchando en http://localhost:" + cfg.Puerto)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("error del servidor: %v", err)
		}
	}()

	// Bloquear hasta recibir señal de parada
	<-stop
	log.Println("Apagando servidor...")

	// Dar 10 segundos para que terminen las requests en curso
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("error en shutdown: %w", err)
	}

	log.Println("Servidor detenido limpiamente.")
	return nil
}

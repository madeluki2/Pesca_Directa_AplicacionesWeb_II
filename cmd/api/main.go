package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"Pesca_Directa_AplicacionesWeb_II/internal/config"
	"Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	handlersPedidos "Pesca_Directa_AplicacionesWeb_II/internal/handlers/gestion_pedidos"
	handlersPesca "Pesca_Directa_AplicacionesWeb_II/internal/handlers/gestion_pesca"
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	servicePedidos "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pedidos"
	servicePesca "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pesca"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

func main() {
	cfg, err := config.Cargar()
	if err != nil {
		log.Fatal(err)
	}
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config.Config) error {
	// driver y dsn: SQLite por defecto, postgres si Storage=="postgres"
	driver := "sqlite"
	dsn := ""
	if cfg.Storage == "postgres" {
		driver = "postgres"
		dsn = "host=db user=pesca password=pesca dbname=pesca_directa port=5432 sslmode=disable"
	}

	recursos, err := storage.Inicializar(driver, dsn, cfg.RutaDB, cfg.Storage)
	if err != nil {
		return err
	}
	defer func() { _ = recursos.Cerrar() }()
	log.Printf("Motor: %s | Backend: %s", driver, recursos.BackendUsado)

	authService := service.NewAuthService(recursos.Usuarios)
	pescaService := servicePesca.NewPescaService(recursos.Pesca)
	pedidoService := servicePedidos.NewPedidoService(recursos.Pedidos)
	rutasService := service.NewRutasService(recursos.Rutas)

	servidorComun := handlers.NewServer(handlers.Deps{
		Rutas: rutasService,
	})

	hPesca := handlersPesca.NewServer(handlersPesca.Deps{Pesca: pescaService})
	hPedidos := handlersPedidos.NewServer(pedidoService, authService)

	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", hPedidos.Registrar)
		r.Post("/auth/login", hPedidos.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			// ── Pesca (Anthony) ───────────────────────────────────────
			r.Get("/pescadores", hPesca.ListarPescadores)
			r.Post("/pescadores", hPesca.CrearPescador)
			r.Get("/pescadores/{id}", hPesca.ObtenerPescador)
			r.Put("/pescadores/{id}", hPesca.ActualizarPescador)
			r.Delete("/pescadores/{id}", hPesca.BorrarPescador)

			r.Get("/embarcaciones", hPesca.ListarEmbarcaciones)
			r.Post("/embarcaciones", hPesca.CrearEmbarcacion)
			r.Get("/embarcaciones/{id}", hPesca.ObtenerEmbarcacion)
			r.Put("/embarcaciones/{id}", hPesca.ActualizarEmbarcacion)
			r.Delete("/embarcaciones/{id}", hPesca.BorrarEmbarcacion)

			r.Get("/especies", hPesca.ListarEspecies)
			r.Post("/especies", hPesca.CrearEspecie)
			r.Get("/especies/{id}", hPesca.ObtenerEspecie)
			r.Put("/especies/{id}", hPesca.ActualizarEspecie)
			r.Delete("/especies/{id}", hPesca.BorrarEspecie)

			r.Get("/capturas", hPesca.ListarCapturas)
			r.Post("/capturas", hPesca.CrearCaptura)
			r.Get("/capturas/{id}", hPesca.ObtenerCaptura)
			r.Put("/capturas/{id}", hPesca.ActualizarCaptura)
			r.Delete("/capturas/{id}", hPesca.BorrarCaptura)

			r.Get("/bodegas", hPesca.ListarBodegas)
			r.Post("/bodegas", hPesca.CrearBodega)
			r.Get("/bodegas/{id}", hPesca.ObtenerBodega)
			r.Put("/bodegas/{id}", hPesca.ActualizarBodega)
			r.Delete("/bodegas/{id}", hPesca.BorrarBodega)

			r.Get("/stocks", hPesca.ListarStocks)
			r.Post("/stocks", hPesca.CrearStock)
			r.Get("/stocks/{id}", hPesca.ObtenerStock)
			r.Put("/stocks/{id}", hPesca.ActualizarStock)
			r.Delete("/stocks/{id}", hPesca.BorrarStock)

			// ── Pedidos (Ilaria) ──────────────────────────────────────
			r.Get("/clientes", hPedidos.ListarClientes)
			r.Post("/clientes", hPedidos.CrearCliente)
			r.Get("/clientes/{id}", hPedidos.ObtenerCliente)
			r.Put("/clientes/{id}", hPedidos.ActualizarCliente)
			r.Delete("/clientes/{id}", hPedidos.EliminarCliente)
			r.Patch("/clientes/{id}/tipo", hPedidos.CambiarTipoCliente)

			r.Get("/pedidos", hPedidos.ListarPedidos)
			r.Post("/pedidos", hPedidos.CrearPedido)
			r.Get("/pedidos/{id}", hPedidos.ObtenerPedido)
			r.Put("/pedidos/{id}", hPedidos.ActualizarPedido)
			r.Delete("/pedidos/{id}", hPedidos.EliminarPedido)

			r.Get("/detalles-pedido", hPedidos.ListarDetalles)
			r.Post("/detalles-pedido", hPedidos.CrearDetalle)
			r.Get("/detalles-pedido/{id}", hPedidos.ObtenerDetalle)
			r.Put("/detalles-pedido/{id}", hPedidos.ActualizarDetalle)
			r.Delete("/detalles-pedido/{id}", hPedidos.EliminarDetalle)

			// ── Rutas (Madelyn) ───────────────────────────────────────
			r.Get("/rutas", servidorComun.ListarRutas)
			r.Post("/rutas", servidorComun.CrearRuta)
			r.Get("/rutas/{id}", servidorComun.ObtenerRuta)
			r.Put("/rutas/{id}", servidorComun.ActualizarRuta)
			r.Delete("/rutas/{id}", servidorComun.BorrarRuta)

			r.Get("/puntos", servidorComun.ListarPuntos)
			r.Post("/puntos", servidorComun.CrearPunto)
			r.Get("/puntos/{id}", servidorComun.ObtenerPunto)
			r.Put("/puntos/{id}", servidorComun.ActualizarPunto)
			r.Delete("/puntos/{id}", servidorComun.BorrarPunto)

			r.Get("/transportistas", servidorComun.ListarTransportistas)
			r.Post("/transportistas", servidorComun.CrearTransportista)
			r.Get("/transportistas/{id}", servidorComun.ObtenerTransportista)
			r.Put("/transportistas/{id}", servidorComun.ActualizarTransportista)
			r.Delete("/transportistas/{id}", servidorComun.BorrarTransportista)

			r.Get("/entregas", servidorComun.ListarEntregas)
			r.Post("/entregas", servidorComun.CrearEntrega)
			r.Get("/entregas/{id}", servidorComun.ObtenerEntrega)
			r.Put("/entregas/{id}", servidorComun.ActualizarEntrega)
			r.Delete("/entregas/{id}", servidorComun.BorrarEntrega)
		})
	})

	srv := &http.Server{
		Addr:         ":" + cfg.Puerto,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		log.Printf("Servidor escuchando en http://localhost:%s", cfg.Puerto)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		log.Println("Apagando servidor...")
	}

	ctxStop, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return srv.Shutdown(ctxStop)
}

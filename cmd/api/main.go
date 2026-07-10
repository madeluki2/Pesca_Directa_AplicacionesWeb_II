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
	handlersPedidos "Pesca_Directa_AplicacionesWeb_II/internal/handlers/gestion_pedidos"
	handlersPesca "Pesca_Directa_AplicacionesWeb_II/internal/handlers/gestion_pesca"
	handlersRutas "Pesca_Directa_AplicacionesWeb_II/internal/handlers/rutas_de_distribucion"
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	servicePedidos "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pedidos"
	servicePesca "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pesca"
	serviceRutas "Pesca_Directa_AplicacionesWeb_II/internal/service/rutas_de_distribucion"
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
	// Inicializar almacenamiento
	recursos, err := storage.Inicializar(cfg.DBDriver, cfg.DBDsn, cfg.RutaDB, cfg.Storage)
	if err != nil {
		return err
	}
	defer func() { _ = recursos.Cerrar() }()
	log.Printf("Motor de base de datos: %s | Backend: %s", cfg.DBDriver, recursos.BackendUsado)

	// Servicios
	authService := service.NewAuthService(
		recursos.Usuarios,
		service.WithSecreto(cfg.JWTSecreto),
		service.WithDuracionToken(cfg.JWTDuracion),
	)
	pescaSvc := servicePesca.NewPescaService(recursos.Pesca)
	pedidoSvc := servicePedidos.NewPedidoService(recursos.Pedidos)
	rutasSvc := serviceRutas.NewRutasService(recursos.Rutas)

	// Handlers
	hPesca := handlersPesca.NewServer(handlersPesca.Deps{Pesca: pescaSvc})
	hPedidos := handlersPedidos.NewServer0(pedidoSvc, authService)
	hRutas := handlersRutas.NewServer0(rutasSvc, authService)

	// Router + middlewares
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", hPedidos.Registrar)
		r.Post("/auth/login", hPedidos.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			// ── Pesca ──
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

			// ── Pedidos ──
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

			// ── Rutas ──
			r.Get("/rutas", hRutas.ListarRutas)
			r.Post("/rutas", hRutas.CrearRuta)
			r.Get("/rutas/{id}", hRutas.ObtenerRuta)
			r.Put("/rutas/{id}", hRutas.ActualizarRuta)
			r.Delete("/rutas/{id}", hRutas.BorrarRuta)

			r.Get("/puntos", hRutas.ListarPuntos)
			r.Post("/puntos", hRutas.CrearPunto)
			r.Get("/puntos/{id}", hRutas.ObtenerPunto)
			r.Put("/puntos/{id}", hRutas.ActualizarPunto)
			r.Delete("/puntos/{id}", hRutas.BorrarPunto)

			r.Get("/transportistas", hRutas.ListarTransportistas)
			r.Post("/transportistas", hRutas.CrearTransportista)
			r.Get("/transportistas/{id}", hRutas.ObtenerTransportista)
			r.Put("/transportistas/{id}", hRutas.ActualizarTransportista)
			r.Delete("/transportistas/{id}", hRutas.BorrarTransportista)

			r.Get("/entregas", hRutas.ListarEntregas)
			r.Post("/entregas", hRutas.CrearEntrega)
			r.Get("/entregas/{id}", hRutas.ObtenerEntrega)
			r.Put("/entregas/{id}", hRutas.ActualizarEntrega)
			r.Delete("/entregas/{id}", hRutas.BorrarEntrega)
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

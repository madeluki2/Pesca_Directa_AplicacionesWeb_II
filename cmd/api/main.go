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
<<<<<<< HEAD
	pedidosHandlers "Pesca_Directa_AplicacionesWeb_II/internal/handlers/gestion_pedidos"
	pescaHandlers "Pesca_Directa_AplicacionesWeb_II/internal/handlers/gestion_pesca"
	rutasHandlers "Pesca_Directa_AplicacionesWeb_II/internal/handlers/rutas_de_distribucion"
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	pedidosService "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pedidos"
	pescaService "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pesca"
	rutasService "Pesca_Directa_AplicacionesWeb_II/internal/service/rutas_de_distribucion"
=======
	"Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	handlersPedidos "Pesca_Directa_AplicacionesWeb_II/internal/handlers/gestion_pedidos"
	handlersPesca "Pesca_Directa_AplicacionesWeb_II/internal/handlers/gestion_pesca"
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
<<<<<<< HEAD
	pedidosService "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pedidos"
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
=======
	servicePedidos "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pedidos"
	servicePesca "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pesca"
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

func main() {
<<<<<<< HEAD
	cfg := config.Cargar()
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config.Config) error {
	// 1. Recursos de almacenamiento (Factory): sqlite en local, postgres en Docker.
	//    Una única conexión *gorm.DB compartida entre los 3 módulos.
	recursos, err := storage.Inicializar(cfg.DBDriver, cfg.DBDsn, cfg.RutaDB, cfg.Backend)
	if err != nil {
		return err
	}
	defer func() { _ = recursos.Cerrar() }()
	log.Printf("Motor de base de datos: %s | Backend: %s", cfg.DBDriver, recursos.BackendUsado)

	// 2. Services con inyección de dependencias.
	authService := service.NewAuthService(
		recursos.Usuarios,
		service.WithSecreto(string(cfg.JWTSecreto)),
		service.WithDuracionToken(cfg.JWTDuracion),
	)
	pescaSvc := pescaService.NewPescaService(recursos.Pesca)
	pedidoSvc := pedidosService.NewPedidoService(recursos.Pedidos)
	rutasSvc := rutasService.NewRutasService(recursos.Rutas)

	// 3. Servers: cada módulo tiene el suyo, cada uno vive en su propio
	//    subpaquete (gestion_pesca, gestion_pedidos, rutas_de_distribucion).
	//    Auth (registro/login) se expone desde el server de Pedidos.
	servidorPesca := pescaHandlers.NewServer(pescaHandlers.Deps{Pesca: pescaSvc})
	servidorPedidos := pedidosHandlers.NewServer(pedidoSvc, authService)
	servidorRutas := rutasHandlers.NewServer0(rutasSvc, authService)

=======
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

<<<<<<< HEAD
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
	// 4. Router + middlewares globales.
=======
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", hPedidos.Registrar)
		r.Post("/auth/login", hPedidos.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

<<<<<<< HEAD
			// ── Anthony: Gestión de Pesca ─────────────────────────────
<<<<<<< HEAD
			r.Get("/pescadores", servidorPesca.ListarPescadores)
			r.Post("/pescadores", servidorPesca.CrearPescador)
			r.Get("/pescadores/{id}", servidorPesca.ObtenerPescador)
			r.Put("/pescadores/{id}", servidorPesca.ActualizarPescador)
			r.Delete("/pescadores/{id}", servidorPesca.BorrarPescador)

			r.Get("/embarcaciones", servidorPesca.ListarEmbarcaciones)
			r.Post("/embarcaciones", servidorPesca.CrearEmbarcacion)
			r.Get("/embarcaciones/{id}", servidorPesca.ObtenerEmbarcacion)
			r.Put("/embarcaciones/{id}", servidorPesca.ActualizarEmbarcacion)
			r.Delete("/embarcaciones/{id}", servidorPesca.BorrarEmbarcacion)

			r.Get("/especies", servidorPesca.ListarEspecies)
			r.Post("/especies", servidorPesca.CrearEspecie)
			r.Get("/especies/{id}", servidorPesca.ObtenerEspecie)
			r.Put("/especies/{id}", servidorPesca.ActualizarEspecie)
			r.Delete("/especies/{id}", servidorPesca.BorrarEspecie)

			r.Get("/capturas", servidorPesca.ListarCapturas)
			r.Post("/capturas", servidorPesca.CrearCaptura)
			r.Get("/capturas/{id}", servidorPesca.ObtenerCaptura)
			r.Put("/capturas/{id}", servidorPesca.ActualizarCaptura)
			r.Delete("/capturas/{id}", servidorPesca.BorrarCaptura)

			r.Get("/bodegas", servidorPesca.ListarBodegas)
			r.Post("/bodegas", servidorPesca.CrearBodega)
			r.Get("/bodegas/{id}", servidorPesca.ObtenerBodega)
			r.Put("/bodegas/{id}", servidorPesca.ActualizarBodega)
			r.Delete("/bodegas/{id}", servidorPesca.BorrarBodega)

			r.Get("/stocks", servidorPesca.ListarStocks)
			r.Post("/stocks", servidorPesca.CrearStock)
			r.Get("/stocks/{id}", servidorPesca.ObtenerStock)
			r.Put("/stocks/{id}", servidorPesca.ActualizarStock)
			r.Delete("/stocks/{id}", servidorPesca.BorrarStock)
=======
			r.Get("/pescadores", servidorComun.ListarPescadores)
			r.Post("/pescadores", servidorComun.CrearPescador)
			r.Get("/pescadores/{id}", servidorComun.ObtenerPescador)
			r.Put("/pescadores/{id}", servidorComun.ActualizarPescador)
			r.Delete("/pescadores/{id}", servidorComun.BorrarPescador)
=======
			// ── Pesca (Anthony) ───────────────────────────────────────
			r.Get("/pescadores", hPesca.ListarPescadores)
			r.Post("/pescadores", hPesca.CrearPescador)
			r.Get("/pescadores/{id}", hPesca.ObtenerPescador)
			r.Put("/pescadores/{id}", hPesca.ActualizarPescador)
			r.Delete("/pescadores/{id}", hPesca.BorrarPescador)
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d

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

<<<<<<< HEAD
			r.Get("/stocks", servidorComun.ListarStocks)
			r.Post("/stocks", servidorComun.CrearStock)
			r.Get("/stocks/{id}", servidorComun.ObtenerStock)
			r.Put("/stocks/{id}", servidorComun.ActualizarStock)
			r.Delete("/stocks/{id}", servidorComun.BorrarStock)
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
=======
			r.Get("/stocks", hPesca.ListarStocks)
			r.Post("/stocks", hPesca.CrearStock)
			r.Get("/stocks/{id}", hPesca.ObtenerStock)
			r.Put("/stocks/{id}", hPesca.ActualizarStock)
			r.Delete("/stocks/{id}", hPesca.BorrarStock)
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d

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

<<<<<<< HEAD
			// ── Madelyn: Rutas de Distribución ────────────────────────
<<<<<<< HEAD
			r.Get("/rutas", servidorRutas.ListarRutas)
			r.Post("/rutas", servidorRutas.CrearRuta)
			r.Get("/rutas/{id}", servidorRutas.ObtenerRuta)
			r.Put("/rutas/{id}", servidorRutas.ActualizarRuta)
			r.Delete("/rutas/{id}", servidorRutas.BorrarRuta)

			r.Get("/puntos", servidorRutas.ListarPuntos)
			r.Post("/puntos", servidorRutas.CrearPunto)
			r.Get("/puntos/{id}", servidorRutas.ObtenerPunto)
			r.Put("/puntos/{id}", servidorRutas.ActualizarPunto)
			r.Delete("/puntos/{id}", servidorRutas.BorrarPunto)

			r.Get("/transportistas", servidorRutas.ListarTransportistas)
			r.Post("/transportistas", servidorRutas.CrearTransportista)
			r.Get("/transportistas/{id}", servidorRutas.ObtenerTransportista)
			r.Put("/transportistas/{id}", servidorRutas.ActualizarTransportista)
			r.Delete("/transportistas/{id}", servidorRutas.BorrarTransportista)

			r.Get("/entregas", servidorRutas.ListarEntregas)
			r.Post("/entregas", servidorRutas.CrearEntrega)
			r.Get("/entregas/{id}", servidorRutas.ObtenerEntrega)
			r.Put("/entregas/{id}", servidorRutas.ActualizarEntrega)
			r.Delete("/entregas/{id}", servidorRutas.BorrarEntrega)
=======
=======
			// ── Rutas (Madelyn) ───────────────────────────────────────
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d
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
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
		})
	})

	srv := &http.Server{
		Addr:         ":" + cfg.Puerto,
		Handler:      r,
<<<<<<< HEAD
<<<<<<< HEAD
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
=======
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
=======
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d
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

<<<<<<< HEAD
	// 10. Graceful shutdown: 10s para terminar requests en curso.
	ctxApagado, cancelar := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelar()
	if err := srv.Shutdown(ctxApagado); err != nil {
		return err
	}
	log.Println("Servidor detenido limpiamente.")
	return nil
<<<<<<< HEAD
}
=======
}
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
=======
	ctxStop, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return srv.Shutdown(ctxStop)
}
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d

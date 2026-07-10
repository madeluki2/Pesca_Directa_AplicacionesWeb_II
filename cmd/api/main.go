// Command pesca-api arranca el servidor HTTP de Pesca-Directa Tarqui.
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"strings"
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
	pedidosHandlers "Pesca_Directa_AplicacionesWeb_II/internal/handlers/gestion_pedidos"
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	pedidosService "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pedidos"
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
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
	// 1. Recursos de almacenamiento (Factory): sqlite en local, postgres en Docker.
	//    Una única conexión *gorm.DB compartida entre los 3 módulos.
	recursos, err := storage.Inicializar("sqlite", "", cfg.RutaDB, cfg.Storage)
	if err != nil {
		return err
	}
	defer func() { _ = recursos.Cerrar() }()
	log.Printf("Motor de base de datos: sqlite | Backend: %s", recursos.BackendUsado)

	// 2. Services con inyección de dependencias.
	authService := service.NewAuthService(recursos.Usuarios)
	pescaSvc := service.NewPescaService(recursos.Pesca)
	pedidoSvc := pedidosService.NewPedidoService(recursos.Pedidos)
	rutasSvc := service.NewRutasService(recursos.Rutas)

	// 3. Servers: cada módulo tiene el suyo, cada uno vive en su propio
	//    subpaquete (gestion_pesca, gestion_pedidos, rutas_de_distribucion).
	//    Auth (registro/login) se expone desde el server de Pedidos.
	servidorComun := handlers.NewServer(pescaSvc, rutasSvc)
	servidorPedidos := pedidosHandlers.NewServer(pedidoSvc, authService)

>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
	// 4. Router + middlewares globales.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 5. Rutas versionadas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", servidorPedidos.Registrar)
		r.Post("/auth/login", servidorPedidos.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

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

			r.Get("/embarcaciones", servidorComun.ListarEmbarcaciones)
			r.Post("/embarcaciones", servidorComun.CrearEmbarcacion)
			r.Get("/embarcaciones/{id}", servidorComun.ObtenerEmbarcacion)
			r.Put("/embarcaciones/{id}", servidorComun.ActualizarEmbarcacion)
			r.Delete("/embarcaciones/{id}", servidorComun.BorrarEmbarcacion)

			r.Get("/especies", servidorComun.ListarEspecies)
			r.Post("/especies", servidorComun.CrearEspecie)
			r.Get("/especies/{id}", servidorComun.ObtenerEspecie)
			r.Put("/especies/{id}", servidorComun.ActualizarEspecie)
			r.Delete("/especies/{id}", servidorComun.BorrarEspecie)

			r.Get("/capturas", servidorComun.ListarCapturas)
			r.Post("/capturas", servidorComun.CrearCaptura)
			r.Get("/capturas/{id}", servidorComun.ObtenerCaptura)
			r.Put("/capturas/{id}", servidorComun.ActualizarCaptura)
			r.Delete("/capturas/{id}", servidorComun.BorrarCaptura)

			r.Get("/bodegas", servidorComun.ListarBodegas)
			r.Post("/bodegas", servidorComun.CrearBodega)
			r.Get("/bodegas/{id}", servidorComun.ObtenerBodega)
			r.Put("/bodegas/{id}", servidorComun.ActualizarBodega)
			r.Delete("/bodegas/{id}", servidorComun.BorrarBodega)

			r.Get("/stocks", servidorComun.ListarStocks)
			r.Post("/stocks", servidorComun.CrearStock)
			r.Get("/stocks/{id}", servidorComun.ObtenerStock)
			r.Put("/stocks/{id}", servidorComun.ActualizarStock)
			r.Delete("/stocks/{id}", servidorComun.BorrarStock)
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327

			// ── Ilaria: Gestión de Pedidos ────────────────────────────
			r.Get("/clientes", servidorPedidos.ListarClientes)
			r.Post("/clientes", servidorPedidos.CrearCliente)
			r.Get("/clientes/{id}", servidorPedidos.ObtenerCliente)
			r.Put("/clientes/{id}", servidorPedidos.ActualizarCliente)
			r.Delete("/clientes/{id}", servidorPedidos.EliminarCliente)
			r.Patch("/clientes/{id}/tipo", servidorPedidos.CambiarTipoCliente)

			r.Get("/pedidos", servidorPedidos.ListarPedidos)
			r.Post("/pedidos", servidorPedidos.CrearPedido)
			r.Get("/pedidos/{id}", servidorPedidos.ObtenerPedido)
			r.Put("/pedidos/{id}", servidorPedidos.ActualizarPedido)
			r.Delete("/pedidos/{id}", servidorPedidos.EliminarPedido)

			r.Get("/detalles-pedido", servidorPedidos.ListarDetalles)
			r.Post("/detalles-pedido", servidorPedidos.CrearDetalle)
			r.Get("/detalles-pedido/{id}", servidorPedidos.ObtenerDetalle)
			r.Put("/detalles-pedido/{id}", servidorPedidos.ActualizarDetalle)
			r.Delete("/detalles-pedido/{id}", servidorPedidos.EliminarDetalle)

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

	// 6. http.Server con timeouts desde config.
	direccionPuerto := cfg.Puerto
	if !strings.HasPrefix(direccionPuerto, ":") {
		direccionPuerto = ":" + direccionPuerto
	}

	srv := &http.Server{
		Addr:         direccionPuerto,
		Handler:      r,
<<<<<<< HEAD
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
=======
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
	}

	// 7. Contexto que se cancela con Ctrl+C / SIGTERM.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 8. Arrancar en goroutine.
	errServidor := make(chan error, 1)
	go func() {
		log.Printf("Servidor Pesca-Directa Tarqui escuchando en http://localhost:%s", strings.TrimPrefix(cfg.Puerto, ":"))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errServidor <- err
		}
	}()

	// 9. Esperar señal o error.
	select {
	case err := <-errServidor:
		return err
	case <-ctx.Done():
		log.Println("Señal de apagado recibida, cerrando ordenadamente...")
	}

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

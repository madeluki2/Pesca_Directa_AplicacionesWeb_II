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
	"Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	pedidosHandlers "Pesca_Directa_AplicacionesWeb_II/internal/handlers/gestion_pedidos"
	rutasHandlers "Pesca_Directa_AplicacionesWeb_II/internal/handlers/rutas_de_distribucion"
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	pedidosService "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pedidos"
	rutasSvc "Pesca_Directa_AplicacionesWeb_II/internal/service/rutas_de_distribucion"
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
	// 1. Recursos de almacenamiento (Factory).
	recursos, err := storage.Inicializar(cfg.DBDriver, cfg.DBDsn, cfg.RutaDB, cfg.Storage)
	if err != nil {
		return err
	}
	defer func() { _ = recursos.Cerrar() }()
	log.Printf("Motor de base de datos: %s | Backend: %s", cfg.DBDriver, recursos.BackendUsado)

	// 2. Services con inyección de dependencias.
	authService := service.NewAuthService(recursos.Usuarios)
	pescaSvc := service.NewPescaService(recursos.Pesca)
	pedidoSvc := pedidosService.NewPedidoService(recursos.Pedidos)
	rutasSvc := rutasSvc.NewRutasService(recursos.Rutas)

	// 3. Servers para cada módulo.
	servidorPescaYAuth := handlers.NewServer(pescaSvc, nil, nil, authService)
	servidorPedidos := pedidosHandlers.NewServer(pedidoSvc, authService)
	servidorRutas := rutasHandlers.NewServer0(rutasSvc, authService)

	// 4. Router + middlewares globales.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 5. Rutas versionadas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", servidorPescaYAuth.Registrar)
		r.Post("/auth/login", servidorPescaYAuth.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			// ── Anthony: Gestión de Pesca ─────────────────────────────
			r.Get("/pescadores", servidorPescaYAuth.ListarPescadores)
			r.Post("/pescadores", servidorPescaYAuth.CrearPescador)
			r.Get("/pescadores/{id}", servidorPescaYAuth.ObtenerPescador)
			r.Put("/pescadores/{id}", servidorPescaYAuth.ActualizarPescador)
			r.Delete("/pescadores/{id}", servidorPescaYAuth.BorrarPescador)

			r.Get("/embarcaciones", servidorPescaYAuth.ListarEmbarcaciones)
			r.Post("/embarcaciones", servidorPescaYAuth.CrearEmbarcacion)
			r.Get("/embarcaciones/{id}", servidorPescaYAuth.ObtenerEmbarcacion)
			r.Put("/embarcaciones/{id}", servidorPescaYAuth.ActualizarEmbarcacion)
			r.Delete("/embarcaciones/{id}", servidorPescaYAuth.BorrarEmbarcacion)

			r.Get("/especies", servidorPescaYAuth.ListarEspecies)
			r.Post("/especies", servidorPescaYAuth.CrearEspecie)
			r.Get("/especies/{id}", servidorPescaYAuth.ObtenerEspecie)
			r.Put("/especies/{id}", servidorPescaYAuth.ActualizarEspecie)
			r.Delete("/especies/{id}", servidorPescaYAuth.BorrarEspecie)

			r.Get("/capturas", servidorPescaYAuth.ListarCapturas)
			r.Post("/capturas", servidorPescaYAuth.CrearCaptura)
			r.Get("/capturas/{id}", servidorPescaYAuth.ObtenerCaptura)
			r.Put("/capturas/{id}", servidorPescaYAuth.ActualizarCaptura)
			r.Delete("/capturas/{id}", servidorPescaYAuth.BorrarCaptura)

			r.Get("/bodegas", servidorPescaYAuth.ListarBodegas)
			r.Post("/bodegas", servidorPescaYAuth.CrearBodega)
			r.Get("/bodegas/{id}", servidorPescaYAuth.ObtenerBodega)
			r.Put("/bodegas/{id}", servidorPescaYAuth.ActualizarBodega)
			r.Delete("/bodegas/{id}", servidorPescaYAuth.BorrarBodega)

			r.Get("/stocks", servidorPescaYAuth.ListarStocks)
			r.Post("/stocks", servidorPescaYAuth.CrearStock)
			r.Get("/stocks/{id}", servidorPescaYAuth.ObtenerStock)
			r.Put("/stocks/{id}", servidorPescaYAuth.ActualizarStock)
			r.Delete("/stocks/{id}", servidorPescaYAuth.BorrarStock)

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
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
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
}

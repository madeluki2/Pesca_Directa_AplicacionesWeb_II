// Command pesca-api arranca el servidor HTTP de Pesca-Directa Tarqui.
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"strings" // Importación añadida para limpiar el prefijo del puerto en el log
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"Pesca_Directa_AplicacionesWeb_II/internal/config"
	"Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	pescaHandlers "Pesca_Directa_AplicacionesWeb_II/internal/handlers/gestion_pesca"
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	pescaService "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pesca"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

func main() {
	cfg := config.Cargar()
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config.Config) error {
	// 1. Factory: abre BD, migra y elige backend.
	recursos, err := storage.Inicializar(cfg.DBDriver, cfg.DBDsn, cfg.RutaDB, cfg.Backend)
	if err != nil {
		return err
	}
	defer func() { _ = recursos.Cerrar() }()
	log.Printf("Motor: %s | Backend: %s", cfg.DBDriver, recursos.BackendUsado)

	// 2. Services — AuthService usa SecretJWT (var global actual).
	authService := service.NewAuthService(recursos.Usuarios)
	pescaSvc := pescaService.NewPescaService(recursos.Pesca)
	pedidoService := service.NewPedidoService(recursos.Pedidos)
	rutasService := service.NewRutasService(recursos.Rutas)

	// 3. Server con Deps.
	servidor := handlers.NewServer(handlers.Deps{
		Pedidos: pedidoService,
		Rutas:   rutasService,
		Auth:    authService,
	})
	servidorPesca := pescaHandlers.NewServer(pescaHandlers.Deps{Pesca: pescaSvc})

	// 4. Router + middlewares.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 5. Rutas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			// ── Anthony: Gestión de Pesca ─────────────────────────────
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

			// ── Ilaria: Gestión de Pedidos ────────────────────────────
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

			// ── Madelyn: Rutas de Distribución ────────────────────────
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

	// 6. http.Server con timeouts desde config - Ajustado con prefijo seguro para la red TCP
	direccionPuerto := cfg.Puerto
	if !strings.HasPrefix(direccionPuerto, ":") {
		direccionPuerto = ":" + direccionPuerto
	}

	srv := &http.Server{
		Addr:         direccionPuerto,
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
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

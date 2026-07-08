// Command pesca-api arranca el servidor HTTP de Pesca-Directa Tarqui.
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
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
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
	//    Cuando se aplique el refactor de Options, esta línea cambia a:
	//    service.NewAuthService(recursos.Usuarios, service.WithSecreto(cfg.JWTSecreto), ...)
	authService := service.NewAuthService(recursos.Usuarios)
	pescaService := service.NewPescaService(recursos.Pesca)
	pedidoService := service.NewPedidoService(recursos.Pedidos)
	rutasService := service.NewRutasService(recursos.Rutas)

	// 3. Server con Deps.
	servidor := handlers.NewServer(handlers.Deps{
		Pesca:   pescaService,
		Pedidos: pedidoService,
		Rutas:   rutasService,
		Auth:    authService,
	})

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

	// 6. http.Server con timeouts desde config (sin paquete httpserver).
	srv := &http.Server{
		Addr:         cfg.Puerto,
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
		log.Printf("Servidor Pesca-Directa Tarqui escuchando en http://localhost%s", cfg.Puerto)
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

// Command pesca-api arranca el servidor HTTP de Pesca-Directa Tarqui.
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"Pesca_Directa_AplicacionesWeb_II/internal/config"
	"Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	handlersRutas "Pesca_Directa_AplicacionesWeb_II/internal/handlers/rutas_de_distribucion"
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	rutasSvc "Pesca_Directa_AplicacionesWeb_II/internal/service/rutas_de_distribucion"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

func main() {
	cfg := config.Cargar()
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config.Config) error {
	// 1. Recursos de almacenamiento (Factory): sqlite en local, postgres en Docker.
	recursos, err := storage.Inicializar(cfg.DBDriver, cfg.DBDsn, cfg.RutaDB, cfg.Backend)
	if err != nil {
		return err
	}
	defer func() { _ = recursos.Cerrar() }()
	log.Printf("Motor de base de datos: %s | Backend: %s", cfg.DBDriver, recursos.BackendUsado)

	// 2. Services con inyeccion de dependencias.
	authService := service.NewAuthService(recursos.Usuarios)
	pescaService := service.NewPescaService(recursos.Pesca)
	pedidoService := service.NewPedidoService(recursos.Pedidos)
	rutasService := rutasSvc.NewRutasService(recursos.Rutas)

	// 3. Servers: punto de entrada para los handlers.
	// Existen dos "Server" separados porque cada equipo definio el suyo:
	// - handlers.Server        -> pescadores, embarcaciones, especies, capturas, bodegas, stocks, clientes, pedidos, detalles-pedido
	// - handlersRutas.Server0  -> auth (registro/login), rutas, puntos, transportistas, entregas
	servidor := handlers.NewServer(pescaService, pedidoService)
	servidorRutas := handlersRutas.NewServer0(rutasService, authService)

	// 4. Router + middlewares globales.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 5. Rutas versionadas /api/v1/ (idénticas a las que ya tenías).
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", servidorRutas.Registrar)
		r.Post("/auth/login", servidorRutas.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

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

	log.Printf("Servidor Pesca-Directa Tarqui escuchando en http://localhost%s", cfg.Puerto)
	return http.ListenAndServe(cfg.Puerto, r)
}

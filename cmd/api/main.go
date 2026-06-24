package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

func main() {
	// 1. Inicializamos el almacén en memoria y cargamos los datos de prueba
	almacen := storage.NewMemoria()
	almacen.Seed()

	// 2. Inyectamos el almacén en el servidor de handlers
	servidor := handlers.NewServer(almacen)

	// 3. Inicializamos el router principal de Chi
	r := chi.NewRouter()

	// 4. Middlewares: Logger para ver peticiones, Recoverer para evitar caídas
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// 5. Registramos las rutas del módulo de Gestión de Pedidos
	r.Route("/api/v1/clientes", func(r chi.Router) {
		r.Get("/", servidor.ListarClientes)
		r.Post("/", servidor.CrearCliente)
		r.Get("/{id}", servidor.ObtenerCliente)
		r.Put("/{id}", servidor.ActualizarCliente)
		r.Patch("/{id}/tipo", servidor.CambiarTipoCliente)
		r.Delete("/{id}", servidor.EliminarCliente)
	})

	r.Route("/api/v1/pedidos", func(r chi.Router) {
		r.Get("/", servidor.ListarPedidos)
		r.Post("/", servidor.CrearPedido)
		r.Get("/{id}", servidor.ObtenerPedido)
		r.Put("/{id}", servidor.ActualizarPedido)
		r.Delete("/{id}", servidor.EliminarPedido)
	})

	r.Route("/api/v1/detalles-pedido", func(r chi.Router) {
		r.Get("/", servidor.ListarDetalles)
		r.Post("/", servidor.CrearDetalle)
		r.Get("/{id}", servidor.ObtenerDetalle)
		r.Put("/{id}", servidor.ActualizarDetalle)
		r.Delete("/{id}", servidor.EliminarDetalle)
	})

	// 6. Levantamos el servidor
	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

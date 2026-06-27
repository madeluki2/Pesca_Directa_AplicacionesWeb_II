package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

func main() {

	// =========================
	// Gestión de Pesca (SQLite)
	// =========================

	db, err := storage.NuevaConexionSQLite()
	if err != nil {
		log.Fatal(err)
	}

	storePesca := storage.NuevaSQLitePesca(db)

	// =========================
	// Gestión de Pedidos
	// =========================

	almacenPedidos := storage.NewMemoria()
	almacenPedidos.Seed()

	servidorPedidos := handlers.NewServer(almacenPedidos)

	// =========================
	// Router principal
	// =========================

	r := chi.NewRouter()

	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)

	// =========================
	// Rutas Pesca
	// =========================

	r.Route("/api/v1", func(r chi.Router) {
		handlers.MontarRutasPesca(r, storePesca)
	})

	// =========================
	// Rutas Pedidos
	// =========================

	r.Route("/api/v1/clientes", func(r chi.Router) {
		r.Get("/", servidorPedidos.ListarClientes)
		r.Post("/", servidorPedidos.CrearCliente)
		// ...
	})

	r.Route("/api/v1/pedidos", func(r chi.Router) {
		r.Get("/", servidorPedidos.ListarPedidos)
		// ...
	})

	// =========================
	// Servidor
	// =========================

	log.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

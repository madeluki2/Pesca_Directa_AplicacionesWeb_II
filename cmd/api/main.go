package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

func main() {
	// ── Instanciar store único en memoria ──────────────
	store := storage.NuevaMemoriaRutas()

	// ── Instanciar handlers ────────────────────────────
	rutaHandler := handlers.NewRutaHandler(store)
	puntoHandler := handlers.NewPuntoHandler(store)
	transportistaHandler := handlers.NewTransportistaHandler(store)
	entregaHandler := handlers.NewEntregaHandler(store)

	// ── Router principal ───────────────────────────────
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// ── Subrouter: módulo Rutas de Distribución ─────────
	r.Route("/api/v1", func(r chi.Router) {

		// Rutas
		r.Route("/rutas", func(r chi.Router) {
			r.Post("/", rutaHandler.Crear)
			r.Get("/", rutaHandler.ObtenerTodos)
			r.Get("/{id}", rutaHandler.ObtenerUno)
			r.Put("/{id}", rutaHandler.Actualizar)
			r.Delete("/{id}", rutaHandler.Eliminar)
		})

		// Puntos
		r.Route("/puntos", func(r chi.Router) {
			r.Post("/", puntoHandler.Crear)
			r.Get("/", puntoHandler.ObtenerTodos)
			r.Get("/{id}", puntoHandler.ObtenerUno)
			r.Put("/{id}", puntoHandler.Actualizar)
			r.Delete("/{id}", puntoHandler.Eliminar)
		})

		// Transportistas
		r.Route("/transportistas", func(r chi.Router) {
			r.Post("/", transportistaHandler.Crear)
			r.Get("/", transportistaHandler.ObtenerTodos)
			r.Get("/{id}", transportistaHandler.ObtenerUno)
			r.Put("/{id}", transportistaHandler.Actualizar)
			r.Delete("/{id}", transportistaHandler.Eliminar)
		})

		// Entregas
		r.Route("/entregas", func(r chi.Router) {
			r.Post("/", entregaHandler.Crear)
			r.Get("/", entregaHandler.ObtenerTodos)
			r.Get("/{id}", entregaHandler.ObtenerUno)
			r.Put("/{id}", entregaHandler.Actualizar)
			r.Delete("/{id}", entregaHandler.Eliminar)
		})
	})

	// ── Arrancar servidor ──────────────────────────────
	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

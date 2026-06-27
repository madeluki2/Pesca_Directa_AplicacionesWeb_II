package main

import (
	"log"
	"net/http"

	"Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

func main() {

	// ============================
	// Almacenamiento en memoria
	// ============================
	storePesca := storage.NuevaMemoriaPesca()

	// ============================
	// Router principal
	// ============================
	r := chi.NewRouter()

	// ============================
	// Middlewares
	// ============================
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)

	// ============================
	// API Version 1
	// ============================
	r.Route("/api/v1", func(r chi.Router) {

		// Gestión de Pesca
		handlers.MontarRutasPesca(r, storePesca)

	})

	// ============================
	// Iniciar servidor
	// ============================
	log.Println("Servidor iniciado en http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}

package handlers

import (
	"net/http"

	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

// Server expone los handlers del paquete raiz.
type Server struct {
	Rutas *service.RutasService
}

// Deps agrupa los servicios requeridos por los handlers de rutas.
type Deps struct {
	Rutas *service.RutasService
}

// NewServer inicializa el controlador de rutas de distribucion.
func NewServer(d Deps) *Server {
	return &Server{
		Rutas: d.Rutas,
	}
}

// BaseHandler es un ping de sanidad para verificar la disponibilidad de la API.
func (s *Server) BaseHandler(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, map[string]string{
		"app":     "Pesca-Directa Tarqui API",
		"status":  "running",
		"version": "1.0.0",
	})
}

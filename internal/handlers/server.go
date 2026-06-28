package handlers

import (
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

// Server agrupa todos los services que los handlers necesitan.
// Al inyectarlos aquí, los handlers nunca dependen del storage directamente.
type Server struct {
	Rutas *service.RutasService
	Auth  *service.AuthService
}

func NewServer(rutas *service.RutasService, auth *service.AuthService) *Server {
	return &Server{
		Rutas: rutas,
		Auth:  auth,
	}
}

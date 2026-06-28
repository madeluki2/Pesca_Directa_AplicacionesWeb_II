package handlers

import (
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

// Server agrupa todos los services que los handlers necesitan.
// Al inyectarlos aquí, los handlers nunca dependen del storage directamente.
type Server struct {
	Pesca *service.PescaService
	Auth  *service.AuthService
}

func NewServer(pesca *service.PescaService, auth *service.AuthService) *Server {
	return &Server{
		Pesca: pesca,
		Auth:  auth,
	}
}

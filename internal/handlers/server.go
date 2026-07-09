package handlers

import (
	// Mantenemos la importación original de Pesca tal cual estaba
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

type Server struct {
	Pesca *service.PescaService // SIN MODIFICAR
	Rutas *service.RutasService
}

func NewServer(
	pesca *service.PescaService,
	rutas *service.RutasService,
) *Server {
	return &Server{
		Pesca: pesca,
		Rutas: rutas,
	}
}

package handlers

import (
	// Mantenemos la importación original de Pesca tal cual estaba
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

type Server struct {
<<<<<<< HEAD
	Pesca   *service.PescaService  // Anthony  — Gestión de Pesca
	Pedidos *service.PedidoService // Ilaria   — Gestión de Pedidos
=======
	Pesca *service.PescaService // SIN MODIFICAR
	Rutas *service.RutasService
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
}

func NewServer(
	pesca *service.PescaService,
<<<<<<< HEAD
	pedidos *service.PedidoService,
) *Server {
	return &Server{
		Pesca:   pesca,
		Pedidos: pedidos,
=======
	rutas *service.RutasService,
) *Server {
	return &Server{
		Pesca: pesca,
		Rutas: rutas,
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
	}
}

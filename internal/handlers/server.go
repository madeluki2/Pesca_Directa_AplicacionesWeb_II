package handlers

import (
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

// Server agrupa los tres services del proyecto.
// Los handlers nunca dependen del storage directamente —
// solo conocen los services que reciben aquí.
type Server struct {
	Pesca   *service.PescaService  // Anthony  — Gestión de Pesca
	Pedidos *service.PedidoService // Ilaria   — Gestión de Pedidos
}

// NewServer crea el servidor con todos los services inyectados.
func NewServer(
	pesca *service.PescaService,
	pedidos *service.PedidoService,
) *Server {
	return &Server{
		Pesca:   pesca,
		Pedidos: pedidos,
	}
}

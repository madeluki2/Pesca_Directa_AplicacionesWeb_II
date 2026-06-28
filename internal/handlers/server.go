package handlers

import (
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

// Server agrupa todos los services que los handlers necesitan.
// Al inyectarlos aquí, los handlers nunca dependen del storage directamente.
type Server struct {
	Pedidos *service.PedidoService
	Auth    *service.AuthService
}

// NewServer crea un servidor con los services inyectados.
func NewServer(pedidos *service.PedidoService, auth *service.AuthService) *Server {
	return &Server{
		Pedidos: pedidos,
		Auth:    auth,
	}
}

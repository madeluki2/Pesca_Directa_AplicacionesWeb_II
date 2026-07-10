package gestion_pedidos

import (
	// Importamos tu rama de Pedidos correctamente
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	pedidosSvc "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pedidos"
)

type Server struct {
	Pedidos *pedidosSvc.PedidoService // Tu rama (Ajusta PedidoService si tu struct tiene otro nombre)
	Auth    *service.AuthService
}

func NewServer(
	pedidos *pedidosSvc.PedidoService,
	auth *service.AuthService,
) *Server {
	return &Server{
		Pedidos: pedidos,
		Auth:    auth,
	}
}

package gestion_pedidos

import (
	// Importamos tu rama de Pedidos correctamente
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	pedidosSvc "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pedidos"
)

type Server0 struct {
	Pedidos *pedidosSvc.PedidoService // Tu rama (Ajusta PedidoService si tu struct tiene otro nombre)
	Auth    *service.AuthService
}

func NewServer0(
	pedidos *pedidosSvc.PedidoService,
	auth *service.AuthService,
) *Server0 {
	return &Server0{
		Pedidos: pedidos,
		Auth:    auth,
	}
}

package handlers

import (
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	gestion_pedidos "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pedidos"
	rutas_de_distribucion "Pesca_Directa_AplicacionesWeb_II/internal/service/rutas_de_distribucion"
)

type Server struct {
	Pesca   *service.PescaService
	Pedidos *gestion_pedidos.PedidoService
	Rutas   *rutas_de_distribucion.RutasService
	Auth    *service.AuthService
}

func NewServer(
	pesca *service.PescaService,
	pedidos *gestion_pedidos.PedidoService,
	rutas *rutas_de_distribucion.RutasService,
	auth *service.AuthService,
) *Server {
	return &Server{
		Pesca:   pesca,
		Pedidos: pedidos,
		Rutas:   rutas,
		Auth:    auth,
	}
}

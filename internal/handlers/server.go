package handlers

import (
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

// Deps agrupa todas las dependencias que los handlers necesitan.
type Deps struct {
	Pedidos *service.PedidoService
	Rutas   *service.RutasService
	Auth    *service.AuthService
}

// Server es el punto único de entrada a todos los handlers.
type Server struct {
	Pedidos *service.PedidoService
	Rutas   *service.RutasService
	Auth    *service.AuthService
}

// NewServer crea el servidor con todas las dependencias inyectadas.
func NewServer(deps Deps) *Server {
	return &Server{
		Pedidos: deps.Pedidos,
		Rutas:   deps.Rutas,
		Auth:    deps.Auth,
	}
}

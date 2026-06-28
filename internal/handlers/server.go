package handlers

import (
    "Pesca_Directa_AplicacionesWeb_II/internal/service"
)

// Server agrupa los tres services del proyecto.
// Los handlers nunca dependen del storage directamente —
// solo conocen los services que reciben aquí.
type Server struct {
    Pesca   service.PescaService   // Anthony  — Gestión de Pesca
    Pedidosservice.PedidoService  // Ilaria   — Gestión de Pedidos
    Rutas   *service.RutasService   // Madelyn  — Rutas de Distribución
    Auth    service.AuthService    // Compartido — autenticación JWT
}

// NewServer crea el servidor con todos los services inyectados.
func NewServer(
    pescaservice.PescaService,
    pedidos service.PedidoService,
    rutasservice.RutasService,
    auth service.AuthService,
	)Server {
		return &Server{
			Pesca:   pesca,
			Pedidos: pedidos,
			Rutas:   rutas,
			Auth:    auth,
    }
}
package handlers

import (
	"net/http"

	"Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pedidos"
	"Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pesca"
	"Pesca_Directa_AplicacionesWeb_II/internal/service/rutas_de_distribucion"
)

// Server expone los handlers del paquete raíz.
type Server struct {
	Pesca   *gestion_pesca.PescaService
	Pedidos *gestion_pedidos.PedidoService      // Revisa si Ilaria le puso PedidoService o PedidosService
	Rutas   *rutas_de_distribucion.RutasService // Madelyn — Gestión de Rutas de Distribución
}

// NewServer inicializa el servidor con los servicios requeridos.
func NewServer(
	pesca *gestion_pesca.PescaService,
	pedidos *gestion_pedidos.PedidoService,
	rutas *rutas_de_distribucion.RutasService,
) *Server {
	return &Server{
		Pesca:   pesca,
		Pedidos: pedidos,
		Rutas:   rutas,
	}
}

// BaseHandler es un ping de sanidad para verificar la disponibilidad de la API.
func (s *Server) BaseHandler(w http.ResponseWriter, r *http.Request) {
	// Asegúrate de que en tu archivo de respuestas la función empiece con Mayúscula
	RespondJSON(w, http.StatusOK, map[string]string{
		"app":     "Pesca-Directa Tarqui API",
		"status":  "running",
		"version": "1.0.0",
	})
}

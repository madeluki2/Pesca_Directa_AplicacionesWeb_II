package handlers

import (
	"net/http"

	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

// Server expone los handlers del paquete raiz.
type Server struct {
<<<<<<< HEAD
<<<<<<< HEAD
	Pesca   *service.PescaService  // Anthony  — Gestión de Pesca
	Pedidos *service.PedidoService // Ilaria   — Gestión de Pedidos
=======
	Pesca *service.PescaService // SIN MODIFICAR
=======
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d
	Rutas *service.RutasService
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
}

<<<<<<< HEAD
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
=======
// Deps agrupa los servicios requeridos por los handlers de rutas.
type Deps struct {
	Rutas *service.RutasService
}

// NewServer inicializa el controlador de rutas de distribucion.
func NewServer(d Deps) *Server {
	return &Server{
		Rutas: d.Rutas,
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d
	}
}

// BaseHandler es un ping de sanidad para verificar la disponibilidad de la API.
func (s *Server) BaseHandler(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, map[string]string{
		"app":     "Pesca-Directa Tarqui API",
		"status":  "running",
		"version": "1.0.0",
	})
}

package rutas_de_distribucion

import (

	// Importamos tu rama de Pedidos correctamente
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	rutasSvc "Pesca_Directa_AplicacionesWeb_II/internal/service/rutas_de_distribucion"
)

// Server0 agrupa los tres services del proyecto.
// Los handlers nunca dependen del storage directamente —
// solo conocen los services que reciben aquí.
type Server0 struct {
	Rutas *rutasSvc.RutasService // Madelyn  — Rutas de Distribución
	Auth  *service.AuthService   // Compartido — autenticación JWT
}

// NewServer0 crea el servidor con todos los services inyectados.
func NewServer0(
	rutas *rutasSvc.RutasService,
	auth *service.AuthService,
) *Server0 {
	return &Server0{
		Rutas: rutas,
		Auth:  auth,
	}
}

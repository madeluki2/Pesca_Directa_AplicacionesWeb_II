package rutas_de_distribucion

import (
	"errors"
	"net/http"

	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

// statusDeError convierte un error del dominio al código HTTP correspondiente.
// Cubre los errores de los tres módulos: Pesca, Pedidos y Rutas.
// Los handlers solo hacen: RespondError(w, statusDeError(err), err.Error())
func statusDeError(err error) int {
	switch {

	// ── 400 Bad Request ──────────────────────────────────────────────────
	// Auth compartido
	case errors.Is(err, service.ErrCredencialesInvalidas),
		errors.Is(err, service.ErrEmailEnUso):
		return http.StatusBadRequest

	// Madelyn — Rutas de Distribución
	case errors.Is(err, service.ErrOrigenVacio),
		errors.Is(err, service.ErrDestinoVacio),
		errors.Is(err, service.ErrRutaIDVacio),
		errors.Is(err, service.ErrPlacaVacia),
		errors.Is(err, service.ErrPlacaDuplicada),
		errors.Is(err, service.ErrPedidoIDVacio),
		errors.Is(err, service.ErrPuntoIDVacio),
		errors.Is(err, service.ErrTransportistaIDVacio):
		return http.StatusBadRequest

	// ── 404 Not Found ────────────────────────────────────────────────────
	// Genérico compartido
	case errors.Is(err, service.ErrNoEncontrado):
		return http.StatusNotFound

	// Madelyn — Rutas de Distribución
	case errors.Is(err, service.ErrRutaNoEncontrada),
		errors.Is(err, service.ErrPuntoNoEncontrado),
		errors.Is(err, service.ErrTransportistaNoEncontrado),
		errors.Is(err, service.ErrEntregaNoEncontrada):
		return http.StatusNotFound

	// ── 500 Internal Server Error ─────────────────────────────────────────
	default:
		return http.StatusInternalServerError
	}
}

package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

// RespondJSON escribe data como JSON con el status HTTP indicado.
// Centraliza el Content-Type y WriteHeader para que todos los handlers
// devuelvan respuestas consistentes.
func RespondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// RespondError escribe un error en formato {"error": "..."}.
func RespondError(w http.ResponseWriter, status int, mensaje string) {
	RespondJSON(w, status, map[string]string{"error": mensaje})
}

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

	// Anthony — Gestión de Pesca
	case errors.Is(err, service.ErrCedulaVacia),
		errors.Is(err, service.ErrPuertoVacio),
		errors.Is(err, service.ErrNombreVacio),
		errors.Is(err, service.ErrMatriculaVacia),
		errors.Is(err, service.ErrNombreComunVacio),
		errors.Is(err, service.ErrUnidadMedidaVacia),
		errors.Is(err, service.ErrCantidadInvalida),
		errors.Is(err, service.ErrCantidadNegativa),
		errors.Is(err, service.ErrFechaVacia),
		errors.Is(err, service.ErrFechaIngresoVacia),
		errors.Is(err, service.ErrFrescuraInvalida),
		errors.Is(err, service.ErrUbicacionVacia),
		errors.Is(err, service.ErrCapacidadInvalida):
		return http.StatusBadRequest

	// Ilaria — Gestión de Pedidos
	case errors.Is(err, service.ErrNombreNegocioVacio),
		errors.Is(err, service.ErrTipoClienteInvalido),
		errors.Is(err, service.ErrTelefonoVacio),
		errors.Is(err, service.ErrDireccionVacia),
		errors.Is(err, service.ErrClienteIDInvalido),
		errors.Is(err, service.ErrPedidoIDInvalido),
		errors.Is(err, service.ErrEspecieIDInvalido),
		errors.Is(err, service.ErrPrecioInvalido):
		return http.StatusBadRequest

	// ── 401 Unauthorized ─────────────────────────────────────────────────
	// (solo Ilaria lo definió explícito, pero aplica a todos)
	case errors.Is(err, service.ErrCredencialesInvalidas):
		return http.StatusUnauthorized

	// ── 404 Not Found ────────────────────────────────────────────────────
	// Genérico compartido
	case errors.Is(err, service.ErrNoEncontrado):
		return http.StatusNotFound

	// Anthony — Gestión de Pesca
	case errors.Is(err, service.ErrPescadorNoEncontrado),
		errors.Is(err, service.ErrEmbarcacionNoEncontrada),
		errors.Is(err, service.ErrEspecieNoEncontrada),
		errors.Is(err, service.ErrCapturaNoEncontrada),
		errors.Is(err, service.ErrBodegaNoEncontrada),
		errors.Is(err, service.ErrStockNoEncontrado):
		return http.StatusNotFound

	// Ilaria — Gestión de Pedidos
	case errors.Is(err, service.ErrClienteNoEncontrado),
		errors.Is(err, service.ErrPedidoNoEncontrado),
		errors.Is(err, service.ErrDetalleNoEncontrado):
		return http.StatusNotFound

	// ── 500 Internal Server Error ─────────────────────────────────────────
	default:
		return http.StatusInternalServerError
	}
}

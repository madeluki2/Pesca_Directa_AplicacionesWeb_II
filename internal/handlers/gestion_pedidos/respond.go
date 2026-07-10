package gestion_pedidos

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
func RespondJson(w http.ResponseWriter, status int, data any) {
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
func RespondERROR(w http.ResponseWriter, status int, mensaje string) {
	RespondJson(w, status, map[string]string{"error": mensaje})
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
	case errors.Is(err, service.ErrCredencialesInvalidas):
		return http.StatusUnauthorized

	// ── 404 Not Found ────────────────────────────────────────────────────
	// Genérico compartido
	case errors.Is(err, service.ErrNoEncontrado):
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

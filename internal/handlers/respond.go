package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

// RespondJSON escribe data como JSON con el status HTTP indicado.
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
// Así los handlers solo hacen: RespondError(w, statusDeError(err), err.Error())
func statusDeError(err error) int {
	switch {
	// 400 Bad Request — datos inválidos enviados por el cliente
	case errors.Is(err, service.ErrEmailEnUso),
		errors.Is(err, service.ErrNombreNegocioVacio),
		errors.Is(err, service.ErrTipoClienteInvalido),
		errors.Is(err, service.ErrTelefonoVacio),
		errors.Is(err, service.ErrDireccionVacia),
		errors.Is(err, service.ErrClienteIDInvalido),
		errors.Is(err, service.ErrFechaVacia),
		errors.Is(err, service.ErrPedidoIDInvalido),
		errors.Is(err, service.ErrEspecieIDInvalido),
		errors.Is(err, service.ErrCantidadInvalida),
		errors.Is(err, service.ErrPrecioInvalido):
		return http.StatusBadRequest

	// 401 Unauthorized — credenciales inválidas
	case errors.Is(err, service.ErrCredencialesInvalidas):
		return http.StatusUnauthorized

	// 404 Not Found — el recurso pedido no existe
	case errors.Is(err, service.ErrNoEncontrado),
		errors.Is(err, service.ErrClienteNoEncontrado),
		errors.Is(err, service.ErrPedidoNoEncontrado),
		errors.Is(err, service.ErrDetalleNoEncontrado):
		return http.StatusNotFound

	// 500 Internal Server Error — cualquier otro error inesperado
	default:
		return http.StatusInternalServerError
	}
}

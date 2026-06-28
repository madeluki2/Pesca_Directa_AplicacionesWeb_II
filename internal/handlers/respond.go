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
	case errors.Is(err, service.ErrCredencialesInvalidas),
		errors.Is(err, service.ErrEmailEnUso),
		errors.Is(err, service.ErrCedulaVacia),
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

	// 404 Not Found — el recurso pedido no existe
	case errors.Is(err, service.ErrNoEncontrado),
		errors.Is(err, service.ErrPescadorNoEncontrado),
		errors.Is(err, service.ErrEmbarcacionNoEncontrada),
		errors.Is(err, service.ErrEspecieNoEncontrada),
		errors.Is(err, service.ErrCapturaNoEncontrada),
		errors.Is(err, service.ErrBodegaNoEncontrada),
		errors.Is(err, service.ErrStockNoEncontrado):
		return http.StatusNotFound

	// 500 Internal Server Error — cualquier otro error inesperado
	default:
		return http.StatusInternalServerError
	}
}

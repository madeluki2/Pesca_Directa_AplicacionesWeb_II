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
		errors.Is(err, service.ErrNombreVacio),
		errors.Is(err, service.ErrOrigenVacio),
		errors.Is(err, service.ErrDestinoVacio),
		errors.Is(err, service.ErrDireccionVacia),
		errors.Is(err, service.ErrRutaIDVacio),
		errors.Is(err, service.ErrTelefonoVacio),
		errors.Is(err, service.ErrPlacaVacia),
		errors.Is(err, service.ErrPlacaDuplicada),
		errors.Is(err, service.ErrPedidoIDVacio),
		errors.Is(err, service.ErrPuntoIDVacio),
		errors.Is(err, service.ErrTransportistaIDVacio):
		return http.StatusBadRequest

	// 404 Not Found — el recurso pedido no existe
	case errors.Is(err, service.ErrNoEncontrado),
		errors.Is(err, service.ErrRutaNoEncontrada),
		errors.Is(err, service.ErrPuntoNoEncontrado),
		errors.Is(err, service.ErrTransportistaNoEncontrado),
		errors.Is(err, service.ErrEntregaNoEncontrada):
		return http.StatusNotFound

	// 500 Internal Server Error — cualquier otro error inesperado
	default:
		return http.StatusInternalServerError
	}
}

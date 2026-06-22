package handlers

import (
	"encoding/json"
	"net/http"
)

// RespondJSON serializa `data` como JSON y escribe el status code en la respuesta.
func RespondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// RespondError envía un JSON de error con el formato estándar { "error": "mensaje" }.
func RespondError(w http.ResponseWriter, status int, mensaje string) {
	RespondJSON(w, status, map[string]string{"error": mensaje})
}

<<<<<<<< HEAD:internal/handlers/rutas_de_distribucion/auth.go
package rutas_de_distribucion
========
package gestion_pedidos
>>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327:internal/handlers/gestion_pedidos/auth.go

import (
	"encoding/json"
	"net/http"

	. "Pesca_Directa_AplicacionesWeb_II/internal/handlers"
)

// credenciales es el body que reciben Registrar y Login.
type credenciales struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Registrar atiende POST /api/v1/auth/register
func (s *Server0) Registrar(w http.ResponseWriter, r *http.Request) {
	var creds credenciales
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	usuario, err := s.Auth.Registrar(creds.Email, creds.Password)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusCreated, usuario)
}

// Login atiende POST /api/v1/auth/login
func (s *Server0) Login(w http.ResponseWriter, r *http.Request) {
	var creds credenciales
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	token, err := s.Auth.Login(creds.Email, creds.Password)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, map[string]string{"token": token})
}

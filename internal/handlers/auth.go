package handlers

import (
	"encoding/json"
	"net/http"
)

// credenciales representa el body de registro y login.
type credenciales struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Registrar crea un nuevo usuario en el sistema (POST /api/v1/auth/register).
func (s *Server) Registrar(w http.ResponseWriter, r *http.Request) {
	var creds credenciales
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	// Validamos que vengan los dos campos obligatorios
	if creds.Email == "" || creds.Password == "" {
		RespondError(w, http.StatusBadRequest, "email y password son obligatorios")
		return
	}

	usuario, err := s.Auth.Registrar(creds.Email, creds.Password)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusCreated, usuario)
}

// Login verifica credenciales y devuelve un token JWT (POST /api/v1/auth/login).
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var creds credenciales
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	// Validamos que vengan los dos campos obligatorios
	if creds.Email == "" || creds.Password == "" {
		RespondError(w, http.StatusBadRequest, "email y password son obligatorios")
		return
	}

	token, err := s.Auth.Login(creds.Email, creds.Password)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	// Devolvemos el token en un objeto JSON
	RespondJSON(w, http.StatusOK, map[string]string{"token": token})
}

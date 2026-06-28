package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// parseID extrae el parámetro {id} de la URL y lo convierte a uint.
// Devuelve false si el valor no es un número entero positivo.
func parseID(r *http.Request) (uint, bool) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		return 0, false
	}
	return uint(id), true
}

// ════════════════════════════ RUTAS ═══════════════════════════════════════

func (s *Server) ListarRutas(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Rutas.ListarRutas())
}

func (s *Server) ObtenerRuta(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	ruta, err := s.Rutas.ObtenerRuta(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, ruta)
}

func (s *Server) CrearRuta(w http.ResponseWriter, r *http.Request) {
	var body models.Ruta
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	ruta, err := s.Rutas.CrearRuta(body)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, ruta)
}

func (s *Server) ActualizarRuta(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Ruta
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	ruta, err := s.Rutas.ActualizarRuta(id, body)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, ruta)
}

func (s *Server) BorrarRuta(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if err := s.Rutas.BorrarRuta(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ PUNTOS ══════════════════════════════════════

func (s *Server) ListarPuntos(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Rutas.ListarPuntos())
}

func (s *Server) ObtenerPunto(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	punto, err := s.Rutas.ObtenerPunto(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, punto)
}

func (s *Server) CrearPunto(w http.ResponseWriter, r *http.Request) {
	var body models.Punto
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	punto, err := s.Rutas.CrearPunto(body)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, punto)
}

func (s *Server) ActualizarPunto(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Punto
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	punto, err := s.Rutas.ActualizarPunto(id, body)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, punto)
}

func (s *Server) BorrarPunto(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if err := s.Rutas.BorrarPunto(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ TRANSPORTISTAS ══════════════════════════════

func (s *Server) ListarTransportistas(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Rutas.ListarTransportistas())
}

func (s *Server) ObtenerTransportista(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	t, err := s.Rutas.ObtenerTransportista(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, t)
}

func (s *Server) CrearTransportista(w http.ResponseWriter, r *http.Request) {
	var body models.Transportista
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	t, err := s.Rutas.CrearTransportista(body)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, t)
}

func (s *Server) ActualizarTransportista(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Transportista
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	t, err := s.Rutas.ActualizarTransportista(id, body)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, t)
}

func (s *Server) BorrarTransportista(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if err := s.Rutas.BorrarTransportista(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ ENTREGAS ════════════════════════════════════

func (s *Server) ListarEntregas(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Rutas.ListarEntregas())
}

func (s *Server) ObtenerEntrega(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	e, err := s.Rutas.ObtenerEntrega(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, e)
}

func (s *Server) CrearEntrega(w http.ResponseWriter, r *http.Request) {
	var body models.EntregaPedido
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	e, err := s.Rutas.CrearEntrega(body)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, e)
}

func (s *Server) ActualizarEntrega(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.EntregaPedido
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	e, err := s.Rutas.ActualizarEntrega(id, body)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, e)
}

func (s *Server) BorrarEntrega(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if err := s.Rutas.BorrarEntrega(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

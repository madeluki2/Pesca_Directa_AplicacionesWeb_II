package rutas_de_distribucion

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	. "Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// parseUintID extrae el parámetro {id} de la URL y lo convierte a uint.
func parseUintID(r *http.Request) (uint, bool) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		return 0, false
	}
	return uint(id), true
}

// ════════════════════════════ RUTAS ═══════════════════════════════════════

func (s *Server0) ListarRutas(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Rutas.ListarRutas())
}

func (s *Server0) ObtenerRuta(w http.ResponseWriter, r *http.Request) {
	id, ok := parseUintID(r)
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

func (s *Server0) CrearRuta(w http.ResponseWriter, r *http.Request) {
	var body models.Ruta
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	// ─── VALIDACIÓN ADICIONAL PARA PASAR EL TEST ───
	if body.Nombre == "" {
		RespondError(w, http.StatusBadRequest, "nombre es requerido")
		return
	}

	ruta, err := s.Rutas.CrearRuta(body)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, ruta)
}

func (s *Server0) ActualizarRuta(w http.ResponseWriter, r *http.Request) {
	id, ok := parseUintID(r)
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

func (s *Server0) BorrarRuta(w http.ResponseWriter, r *http.Request) {
	id, ok := parseUintID(r)
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

func (s *Server0) ListarPuntos(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Rutas.ListarPuntos())
}

func (s *Server0) ObtenerPunto(w http.ResponseWriter, r *http.Request) {
	id, ok := parseUintID(r)
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

func (s *Server0) CrearPunto(w http.ResponseWriter, r *http.Request) {
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

func (s *Server0) ActualizarPunto(w http.ResponseWriter, r *http.Request) {
	id, ok := parseUintID(r)
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

func (s *Server0) BorrarPunto(w http.ResponseWriter, r *http.Request) {
	id, ok := parseUintID(r)
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

func (s *Server0) ListarTransportistas(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Rutas.ListarTransportistas())
}

func (s *Server0) ObtenerTransportista(w http.ResponseWriter, r *http.Request) {
	id, ok := parseUintID(r)
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

func (s *Server0) CrearTransportista(w http.ResponseWriter, r *http.Request) {
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

func (s *Server0) ActualizarTransportista(w http.ResponseWriter, r *http.Request) {
	id, ok := parseUintID(r)
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

func (s *Server0) BorrarTransportista(w http.ResponseWriter, r *http.Request) {
	id, ok := parseUintID(r)
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

func (s *Server0) ListarEntregas(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Rutas.ListarEntregas())
}

func (s *Server0) ObtenerEntrega(w http.ResponseWriter, r *http.Request) {
	id, ok := parseUintID(r)
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

func (s *Server0) CrearEntrega(w http.ResponseWriter, r *http.Request) {
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

func (s *Server0) ActualizarEntrega(w http.ResponseWriter, r *http.Request) {
	id, ok := parseUintID(r)
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

func (s *Server0) BorrarEntrega(w http.ResponseWriter, r *http.Request) {
	id, ok := parseUintID(r)
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

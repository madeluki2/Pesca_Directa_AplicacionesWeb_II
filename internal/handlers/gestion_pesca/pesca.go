package gestion_pesca

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	handlers "Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	service "Pesca_Directa_AplicacionesWeb_II/internal/service/gestion_pesca"
)

type Deps struct {
	Pesca *service.PescaService
}

type Server struct {
	Pesca *service.PescaService
}

func NewServer(deps Deps) *Server {
	return &Server{Pesca: deps.Pesca}
}

func statusDeError(err error) int {
	switch {
	case errors.Is(err, service.ErrCredencialesInvalidas),
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
	case errors.Is(err, service.ErrPescadorNoEncontrado),
		errors.Is(err, service.ErrEmbarcacionNoEncontrada),
		errors.Is(err, service.ErrEspecieNoEncontrada),
		errors.Is(err, service.ErrCapturaNoEncontrada),
		errors.Is(err, service.ErrBodegaNoEncontrada),
		errors.Is(err, service.ErrStockNoEncontrado):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

// parseID extrae el parámetro {id} de la URL y lo convierte a int.
// Devuelve false si el valor no es un número entero positivo.
func parseID(r *http.Request) (int, bool) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return 0, false
	}
	return id, true
}

// ════════════════════════════ PESCADORES ══════════════════════════════════════

func (s *Server) ListarPescadores(w http.ResponseWriter, r *http.Request) {
	handlers.RespondJSON(w, http.StatusOK, s.Pesca.ListarPescadores())
}

func (s *Server) ObtenerPescador(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	p, err := s.Pesca.ObtenerPescador(id)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusOK, p)
}

func (s *Server) CrearPescador(w http.ResponseWriter, r *http.Request) {
	var body models.Pescador
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	p, err := s.Pesca.CrearPescador(body)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusCreated, p)
}

func (s *Server) ActualizarPescador(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Pescador
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	p, err := s.Pesca.ActualizarPescador(id, body)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusOK, p)
}

func (s *Server) BorrarPescador(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if err := s.Pesca.BorrarPescador(id); err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ EMBARCACIONES ═══════════════════════════════════

func (s *Server) ListarEmbarcaciones(w http.ResponseWriter, r *http.Request) {
	handlers.RespondJSON(w, http.StatusOK, s.Pesca.ListarEmbarcaciones())
}

func (s *Server) ObtenerEmbarcacion(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	e, err := s.Pesca.ObtenerEmbarcacion(id)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusOK, e)
}

func (s *Server) CrearEmbarcacion(w http.ResponseWriter, r *http.Request) {
	var body models.Embarcacion
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	e, err := s.Pesca.CrearEmbarcacion(body)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusCreated, e)
}

func (s *Server) ActualizarEmbarcacion(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Embarcacion
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	e, err := s.Pesca.ActualizarEmbarcacion(id, body)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusOK, e)
}

func (s *Server) BorrarEmbarcacion(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if err := s.Pesca.BorrarEmbarcacion(id); err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ ESPECIES ════════════════════════════════════════

func (s *Server) ListarEspecies(w http.ResponseWriter, r *http.Request) {
	handlers.RespondJSON(w, http.StatusOK, s.Pesca.ListarEspecies())
}

func (s *Server) ObtenerEspecie(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	e, err := s.Pesca.ObtenerEspecie(id)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusOK, e)
}

func (s *Server) CrearEspecie(w http.ResponseWriter, r *http.Request) {
	var body models.Especie
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	e, err := s.Pesca.CrearEspecie(body)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusCreated, e)
}

func (s *Server) ActualizarEspecie(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Especie
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	e, err := s.Pesca.ActualizarEspecie(id, body)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusOK, e)
}

func (s *Server) BorrarEspecie(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if err := s.Pesca.BorrarEspecie(id); err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ CAPTURAS ════════════════════════════════════════

func (s *Server) ListarCapturas(w http.ResponseWriter, r *http.Request) {
	handlers.RespondJSON(w, http.StatusOK, s.Pesca.ListarCapturas())
}

func (s *Server) ObtenerCaptura(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	c, err := s.Pesca.ObtenerCaptura(id)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusOK, c)
}

func (s *Server) CrearCaptura(w http.ResponseWriter, r *http.Request) {
	var body models.Captura
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	c, err := s.Pesca.CrearCaptura(body)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusCreated, c)
}

func (s *Server) ActualizarCaptura(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Captura
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	c, err := s.Pesca.ActualizarCaptura(id, body)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusOK, c)
}

func (s *Server) BorrarCaptura(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if err := s.Pesca.BorrarCaptura(id); err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ BODEGAS ═════════════════════════════════════════

func (s *Server) ListarBodegas(w http.ResponseWriter, r *http.Request) {
	handlers.RespondJSON(w, http.StatusOK, s.Pesca.ListarBodegas())
}

func (s *Server) ObtenerBodega(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	b, err := s.Pesca.ObtenerBodega(id)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusOK, b)
}

func (s *Server) CrearBodega(w http.ResponseWriter, r *http.Request) {
	var body models.Bodega
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	b, err := s.Pesca.CrearBodega(body)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusCreated, b)
}

func (s *Server) ActualizarBodega(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Bodega
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	b, err := s.Pesca.ActualizarBodega(id, body)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusOK, b)
}

func (s *Server) BorrarBodega(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if err := s.Pesca.BorrarBodega(id); err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ STOCKS ══════════════════════════════════════════

func (s *Server) ListarStocks(w http.ResponseWriter, r *http.Request) {
	handlers.RespondJSON(w, http.StatusOK, s.Pesca.ListarStocks())
}

func (s *Server) ObtenerStock(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	st, err := s.Pesca.ObtenerStock(id)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusOK, st)
}

func (s *Server) CrearStock(w http.ResponseWriter, r *http.Request) {
	var body models.Stock
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	st, err := s.Pesca.CrearStock(body)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusCreated, st)
}

func (s *Server) ActualizarStock(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Stock
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	st, err := s.Pesca.ActualizarStock(id, body)
	if err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusOK, st)
}

func (s *Server) BorrarStock(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		handlers.RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if err := s.Pesca.BorrarStock(id); err != nil {
		handlers.RespondError(w, statusDeError(err), err.Error())
		return
	}
	handlers.RespondJSON(w, http.StatusNoContent, nil)
}

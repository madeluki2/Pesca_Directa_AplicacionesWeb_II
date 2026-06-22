package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"

	"github.com/go-chi/chi/v5"
)

// ──────────────────────────────────────────────
// Helper: parsear {id} del path
// ──────────────────────────────────────────────
func parseID(r *http.Request) (uint, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// ══════════════════════════════════════════════
// RutaHandler
// ══════════════════════════════════════════════
type RutaHandler struct{ store storage.AlmacenRuta }

func NewRutaHandler(s storage.AlmacenRuta) *RutaHandler { return &RutaHandler{store: s} }

func (h *RutaHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var ruta models.Ruta
	if err := json.NewDecoder(r.Body).Decode(&ruta); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(ruta.Nombre) == "" || strings.TrimSpace(ruta.Origen) == "" || strings.TrimSpace(ruta.Destino) == "" {
		RespondError(w, http.StatusBadRequest, "nombre, origen y destino son obligatorios")
		return
	}
	if ruta.Estado == "" {
		ruta.Estado = "activo"
	}
	creado, err := h.store.CrearRuta(ruta)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (h *RutaHandler) ObtenerTodos(w http.ResponseWriter, r *http.Request) {
	rutas, err := h.store.ObtenerRutas()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, rutas)
}

func (h *RutaHandler) ObtenerUno(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	ruta, err := h.store.ObtenerRutaPorID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, "ruta no encontrada")
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, ruta)
}

func (h *RutaHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	var ruta models.Ruta
	if err := json.NewDecoder(r.Body).Decode(&ruta); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(ruta.Nombre) == "" || strings.TrimSpace(ruta.Origen) == "" || strings.TrimSpace(ruta.Destino) == "" {
		RespondError(w, http.StatusBadRequest, "nombre, origen y destino son obligatorios")
		return
	}
	actualizado, err := h.store.ActualizarRuta(id, ruta)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, "ruta no encontrada")
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (h *RutaHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	if err := h.store.EliminarRuta(id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, "ruta no encontrada")
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "ruta eliminada correctamente"})
}

// ══════════════════════════════════════════════
// PuntoHandler
// ══════════════════════════════════════════════
type PuntoHandler struct{ store storage.AlmacenRuta }

func NewPuntoHandler(s storage.AlmacenRuta) *PuntoHandler { return &PuntoHandler{store: s} }

func (h *PuntoHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var punto models.Punto
	if err := json.NewDecoder(r.Body).Decode(&punto); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(punto.Nombre) == "" || strings.TrimSpace(punto.Direccion) == "" || punto.RutaID == 0 {
		RespondError(w, http.StatusBadRequest, "nombre, direccion y ruta_id son obligatorios")
		return
	}
	if punto.Estado == "" {
		punto.Estado = "activo"
	}
	creado, err := h.store.CrearPunto(punto)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (h *PuntoHandler) ObtenerTodos(w http.ResponseWriter, r *http.Request) {
	puntos, err := h.store.ObtenerPuntos()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, puntos)
}

func (h *PuntoHandler) ObtenerUno(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	punto, err := h.store.ObtenerPuntoPorID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, "punto no encontrado")
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, punto)
}

func (h *PuntoHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	var punto models.Punto
	if err := json.NewDecoder(r.Body).Decode(&punto); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(punto.Nombre) == "" || strings.TrimSpace(punto.Direccion) == "" || punto.RutaID == 0 {
		RespondError(w, http.StatusBadRequest, "nombre, direccion y ruta_id son obligatorios")
		return
	}
	actualizado, err := h.store.ActualizarPunto(id, punto)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, "punto no encontrado")
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (h *PuntoHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	if err := h.store.EliminarPunto(id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, "punto no encontrado")
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "punto eliminado correctamente"})
}

// ══════════════════════════════════════════════
// TransportistaHandler
// ══════════════════════════════════════════════
type TransportistaHandler struct{ store storage.AlmacenRuta }

func NewTransportistaHandler(s storage.AlmacenRuta) *TransportistaHandler {
	return &TransportistaHandler{store: s}
}

func (h *TransportistaHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var t models.Transportista
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(t.Nombre) == "" || strings.TrimSpace(t.Telefono) == "" || strings.TrimSpace(t.PlacaVehiculo) == "" {
		RespondError(w, http.StatusBadRequest, "nombre, telefono y placa_vehiculo son obligatorios")
		return
	}
	if t.Estado == "" {
		t.Estado = "activo"
	}
	creado, err := h.store.CrearTransportista(t)
	if err != nil {
		if errors.Is(err, storage.ErrPlacaDuplicada) {
			RespondError(w, http.StatusBadRequest, "la placa de vehículo ya existe")
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (h *TransportistaHandler) ObtenerTodos(w http.ResponseWriter, r *http.Request) {
	lista, err := h.store.ObtenerTransportistas()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, lista)
}

func (h *TransportistaHandler) ObtenerUno(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	t, err := h.store.ObtenerTransportistaPorID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, "transportista no encontrado")
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, t)
}

func (h *TransportistaHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	var t models.Transportista
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(t.Nombre) == "" || strings.TrimSpace(t.Telefono) == "" || strings.TrimSpace(t.PlacaVehiculo) == "" {
		RespondError(w, http.StatusBadRequest, "nombre, telefono y placa_vehiculo son obligatorios")
		return
	}
	actualizado, err := h.store.ActualizarTransportista(id, t)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, "transportista no encontrado")
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (h *TransportistaHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	if err := h.store.EliminarTransportista(id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, "transportista no encontrado")
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "transportista eliminado correctamente"})
}

// ══════════════════════════════════════════════
// EntregaHandler
// ══════════════════════════════════════════════
type EntregaHandler struct{ store storage.AlmacenRuta }

func NewEntregaHandler(s storage.AlmacenRuta) *EntregaHandler {
	return &EntregaHandler{store: s}
}

func (h *EntregaHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var e models.EntregaPedido
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if e.PedidoID == 0 || e.PuntoID == 0 || e.TransportistaID == 0 {
		RespondError(w, http.StatusBadRequest, "pedido_id, punto_id y transportista_id son obligatorios")
		return
	}
	if e.Estado == "" {
		e.Estado = "pendiente"
	}
	creado, err := h.store.CrearEntrega(e)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (h *EntregaHandler) ObtenerTodos(w http.ResponseWriter, r *http.Request) {
	lista, err := h.store.ObtenerEntregas()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, lista)
}

func (h *EntregaHandler) ObtenerUno(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	e, err := h.store.ObtenerEntregaPorID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, "entrega no encontrada")
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, e)
}

func (h *EntregaHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	var e models.EntregaPedido
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if e.PedidoID == 0 || e.PuntoID == 0 || e.TransportistaID == 0 {
		RespondError(w, http.StatusBadRequest, "pedido_id, punto_id y transportista_id son obligatorios")
		return
	}
	actualizado, err := h.store.ActualizarEntrega(id, e)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, "entrega no encontrada")
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (h *EntregaHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	if err := h.store.EliminarEntrega(id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, "entrega no encontrada")
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "entrega eliminada correctamente"})
}

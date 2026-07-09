package gestion_pedidos

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"

	"github.com/go-chi/chi/v5"
)

// -------------------- CLIENTES --------------------

func (s *Server) ListarClientes(w http.ResponseWriter, _ *http.Request) {
	clientes, err := s.Pedidos.ListarClientes()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "Error al listar clientes: "+err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, clientes)
}

func (s *Server) ObtenerCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	cliente, err := s.Pedidos.ObtenerCliente(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, cliente)
}

func (s *Server) CrearCliente(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	cliente, err := s.Pedidos.CrearCliente(nuevo)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, cliente)
}

func (s *Server) ActualizarCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	var datos models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	cliente, err := s.Pedidos.ActualizarCliente(id, datos)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, cliente)
}

func (s *Server) EliminarCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	if err := s.Pedidos.EliminarCliente(id); err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) CambiarTipoCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	var body struct{ Tipo string }
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	cliente, err := s.Pedidos.CambiarTipoCliente(id, body.Tipo)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, cliente)
}

// -------------------- PEDIDOS --------------------

func (s *Server) ListarPedidos(w http.ResponseWriter, _ *http.Request) {
	pedidos, err := s.Pedidos.ListarPedidos()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "Error al listar pedidos: "+err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, pedidos)
}

func (s *Server) ObtenerPedido(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	pedido, err := s.Pedidos.ObtenerPedido(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, pedido)
}

func (s *Server) CrearPedido(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Pedido
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	pedido, err := s.Pedidos.CrearPedido(nuevo)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, pedido)
}

func (s *Server) ActualizarPedido(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	var datos models.Pedido
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	pedido, err := s.Pedidos.ActualizarPedido(id, datos)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, pedido)
}

func (s *Server) EliminarPedido(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	if err := s.Pedidos.EliminarPedido(id); err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// -------------------- DETALLES --------------------

func (s *Server) ListarDetalles(w http.ResponseWriter, _ *http.Request) {
	detalles, err := s.Pedidos.ListarDetalles()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "Error al listar detalles: "+err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, detalles)
}

func (s *Server) ObtenerDetalle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	detalle, err := s.Pedidos.ObtenerDetalle(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, detalle)
}

func (s *Server) CrearDetalle(w http.ResponseWriter, r *http.Request) {
	var nuevo models.DetallePedido
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	detalle, err := s.Pedidos.CrearDetalle(nuevo)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, detalle)
}

func (s *Server) ActualizarDetalle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	var datos models.DetallePedido
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	detalle, err := s.Pedidos.ActualizarDetalle(id, datos)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, detalle)
}

func (s *Server) EliminarDetalle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	if err := s.Pedidos.EliminarDetalle(id); err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

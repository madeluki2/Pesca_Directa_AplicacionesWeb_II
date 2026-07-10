package gestion_pedidos

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"

	"github.com/go-chi/chi/v5"
)

// -------------------- CLIENTES --------------------

func (s *Server0) ListarClientes(w http.ResponseWriter, _ *http.Request) {
	clientes, err := s.Pedidos.ListarClientes()
	if err != nil {
		RespondERROR(w, http.StatusInternalServerError, "Error al listar clientes: "+err.Error())
		return
	}
	RespondJson(w, http.StatusOK, clientes)
}

func (s *Server0) ObtenerCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondERROR(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	cliente, err := s.Pedidos.ObtenerCliente(id)
	if err != nil {
		RespondERROR(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJson(w, http.StatusOK, cliente)
}

func (s *Server0) CrearCliente(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondERROR(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	cliente, err := s.Pedidos.CrearCliente(nuevo)
	if err != nil {
		RespondERROR(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJson(w, http.StatusCreated, cliente)
}

func (s *Server0) ActualizarCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondERROR(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	var datos models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondERROR(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	cliente, err := s.Pedidos.ActualizarCliente(id, datos)
	if err != nil {
		RespondERROR(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJson(w, http.StatusOK, cliente)
}

func (s *Server0) EliminarCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondERROR(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	if err := s.Pedidos.EliminarCliente(id); err != nil {
		RespondERROR(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server0) CambiarTipoCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondERROR(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	var body struct{ Tipo string }
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondERROR(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	cliente, err := s.Pedidos.CambiarTipoCliente(id, body.Tipo)
	if err != nil {
		RespondERROR(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJson(w, http.StatusOK, cliente)
}

// -------------------- PEDIDOS --------------------

func (s *Server0) ListarPedidos(w http.ResponseWriter, _ *http.Request) {
	pedidos, err := s.Pedidos.ListarPedidos()
	if err != nil {
		RespondERROR(w, http.StatusInternalServerError, "Error al listar pedidos: "+err.Error())
		return
	}
	RespondJson(w, http.StatusOK, pedidos)
}

func (s *Server0) ObtenerPedido(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondERROR(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	pedido, err := s.Pedidos.ObtenerPedido(id)
	if err != nil {
		RespondERROR(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJson(w, http.StatusOK, pedido)
}

func (s *Server0) CrearPedido(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Pedido
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondERROR(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	pedido, err := s.Pedidos.CrearPedido(nuevo)
	if err != nil {
		RespondERROR(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJson(w, http.StatusCreated, pedido)
}

func (s *Server0) ActualizarPedido(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondERROR(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	var datos models.Pedido
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondERROR(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	pedido, err := s.Pedidos.ActualizarPedido(id, datos)
	if err != nil {
		RespondERROR(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJson(w, http.StatusOK, pedido)
}

func (s *Server0) EliminarPedido(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondERROR(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	if err := s.Pedidos.EliminarPedido(id); err != nil {
		RespondERROR(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// -------------------- DETALLES --------------------

func (s *Server0) ListarDetalles(w http.ResponseWriter, _ *http.Request) {
	detalles, err := s.Pedidos.ListarDetalles()
	if err != nil {
		RespondERROR(w, http.StatusInternalServerError, "Error al listar detalles: "+err.Error())
		return
	}
	RespondJson(w, http.StatusOK, detalles)
}

func (s *Server0) ObtenerDetalle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondERROR(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	detalle, err := s.Pedidos.ObtenerDetalle(id)
	if err != nil {
		RespondERROR(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJson(w, http.StatusOK, detalle)
}

func (s *Server0) CrearDetalle(w http.ResponseWriter, r *http.Request) {
	var nuevo models.DetallePedido
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondERROR(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	detalle, err := s.Pedidos.CrearDetalle(nuevo)
	if err != nil {
		RespondERROR(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJson(w, http.StatusCreated, detalle)
}

func (s *Server0) ActualizarDetalle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondERROR(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	var datos models.DetallePedido
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondERROR(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	detalle, err := s.Pedidos.ActualizarDetalle(id, datos)
	if err != nil {
		RespondERROR(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJson(w, http.StatusOK, detalle)
}

func (s *Server0) EliminarDetalle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondERROR(w, http.StatusBadRequest, "El id debe ser un número entero")
		return
	}
	if err := s.Pedidos.EliminarDetalle(id); err != nil {
		RespondERROR(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

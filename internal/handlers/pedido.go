package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// -------------------- CLIENTES --------------------

// ListarClientes devuelve todos los clientes registrados (GET)
func (s *Server) ListarClientes(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Pedidos.ListarClientes())
}

// ObtenerCliente devuelve un cliente por su ID (GET)
func (s *Server) ObtenerCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	cliente, err := s.Pedidos.ObtenerCliente(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, cliente)
}

// CrearCliente registra un nuevo cliente (POST)
func (s *Server) CrearCliente(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	cliente, err := s.Pedidos.CrearCliente(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusCreated, cliente)
}

// ActualizarCliente modifica los datos de un cliente existente (PUT)
func (s *Server) ActualizarCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	var datos models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	cliente, err := s.Pedidos.ActualizarCliente(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, cliente)
}

// EliminarCliente remueve un cliente por su ID (DELETE)
func (s *Server) EliminarCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	if err := s.Pedidos.EliminarCliente(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// CambiarTipoCliente actualiza únicamente el tipo de un cliente (PATCH)
func (s *Server) CambiarTipoCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	var body struct {
		TipoCliente string `json:"tipo_cliente"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	cliente, err := s.Pedidos.CambiarTipoCliente(id, body.TipoCliente)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, cliente)
}

// -------------------- PEDIDOS --------------------

// ListarPedidos devuelve todos los pedidos registrados (GET)
func (s *Server) ListarPedidos(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Pedidos.ListarPedidos())
}

// ObtenerPedido devuelve un pedido por su ID (GET)
func (s *Server) ObtenerPedido(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	pedido, err := s.Pedidos.ObtenerPedido(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, pedido)
}

// CrearPedido registra un nuevo pedido (POST)
func (s *Server) CrearPedido(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Pedido
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	pedido, err := s.Pedidos.CrearPedido(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusCreated, pedido)
}

// ActualizarPedido modifica los datos de un pedido existente (PUT)
func (s *Server) ActualizarPedido(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	var datos models.Pedido
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	pedido, err := s.Pedidos.ActualizarPedido(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, pedido)
}

// EliminarPedido cancela y remueve un pedido por su ID (DELETE)
func (s *Server) EliminarPedido(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	if err := s.Pedidos.EliminarPedido(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// -------------------- DETALLES DE PEDIDO --------------------

// ListarDetalles devuelve todos los detalles de pedido registrados (GET)
func (s *Server) ListarDetalles(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Pedidos.ListarDetalles())
}

// ObtenerDetalle devuelve un detalle de pedido por su ID (GET)
func (s *Server) ObtenerDetalle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	detalle, err := s.Pedidos.ObtenerDetalle(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, detalle)
}

// CrearDetalle registra un nuevo detalle dentro de un pedido (POST)
func (s *Server) CrearDetalle(w http.ResponseWriter, r *http.Request) {
	var nuevo models.DetallePedido
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	detalle, err := s.Pedidos.CrearDetalle(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusCreated, detalle)
}

// ActualizarDetalle modifica los datos de un detalle de pedido existente (PUT)
func (s *Server) ActualizarDetalle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	var datos models.DetallePedido
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	detalle, err := s.Pedidos.ActualizarDetalle(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, detalle)
}

// EliminarDetalle remueve un detalle de pedido por su ID (DELETE)
func (s *Server) EliminarDetalle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	if err := s.Pedidos.EliminarDetalle(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

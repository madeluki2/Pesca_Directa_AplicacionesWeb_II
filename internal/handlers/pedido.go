package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

// Server unifica todos los handlers del módulo de Gestión de Pedidos
type Server struct {
	storage *storage.Memoria
}

// NewServer es el constructor que inyecta el almacén en el servidor
func NewServer(s *storage.Memoria) *Server {
	return &Server{storage: s}
}

// -------------------- CLIENTES --------------------

// ListarClientes devuelve todos los clientes registrados (GET)
func (s *Server) ListarClientes(w http.ResponseWriter, _ *http.Request) {
	clientes := s.storage.ListarClientes()
	RespondJSON(w, http.StatusOK, clientes)
}

// ObtenerCliente devuelve un cliente por su ID (GET)
func (s *Server) ObtenerCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	cliente, encontrado := s.storage.BuscarClientePorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "cliente no encontrado")
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

	// Validación de campos obligatorios
	if strings.TrimSpace(nuevo.NombreNegocio) == "" || strings.TrimSpace(nuevo.TipoCliente) == "" || strings.TrimSpace(nuevo.Telefono) == "" || strings.TrimSpace(nuevo.Direccion) == "" {
		RespondError(w, http.StatusBadRequest, "NombreNegocio, TipoCliente, Telefono y Direccion son obligatorios")
		return
	}

	// El estado siempre inicia como activo
	nuevo.Estado = "activo"
	nuevo = s.storage.CrearCliente(nuevo)
	RespondJSON(w, http.StatusCreated, nuevo)
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

	// Validamos que al menos un campo venga para actualizar
	if strings.TrimSpace(datos.NombreNegocio) == "" && strings.TrimSpace(datos.TipoCliente) == "" && strings.TrimSpace(datos.Telefono) == "" && strings.TrimSpace(datos.Direccion) == "" && strings.TrimSpace(datos.Estado) == "" {
		RespondError(w, http.StatusBadRequest, "debe enviar al menos un campo para actualizar")
		return
	}

	actualizado, encontrado := s.storage.ActualizarCliente(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "cliente no encontrado para actualizar")
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

// EliminarCliente remueve un cliente por su ID (DELETE)
func (s *Server) EliminarCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	if !s.storage.EliminarCliente(id) {
		RespondError(w, http.StatusNotFound, "cliente no encontrado para eliminar")
		return
	}

	// 204 No Content: eliminación exitosa sin cuerpo de respuesta
	w.WriteHeader(http.StatusNoContent)
}

// CambiarTipoCliente actualiza únicamente el tipo de un cliente (PATCH)
func (s *Server) CambiarTipoCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	// Recibimos solo el tipo_cliente en el body
	var body struct {
		TipoCliente string `json:"tipo_cliente"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	// Validamos que el tipo sea uno de los permitidos
	tipo := strings.TrimSpace(body.TipoCliente)
	if tipo != "restaurante" && tipo != "intermediario" && tipo != "mayorista" {
		RespondError(w, http.StatusBadRequest, "tipo_cliente debe ser: restaurante, intermediario o mayorista")
		return
	}

	actualizado, encontrado := s.storage.CambiarTipoCliente(id, tipo)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "cliente no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

// -------------------- PEDIDOS --------------------

// ListarPedidos devuelve todos los pedidos registrados (GET)
func (s *Server) ListarPedidos(w http.ResponseWriter, _ *http.Request) {
	pedidos := s.storage.ListarPedidos()
	RespondJSON(w, http.StatusOK, pedidos)
}

// ObtenerPedido devuelve un pedido por su ID (GET)
func (s *Server) ObtenerPedido(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	pedido, encontrado := s.storage.BuscarPedidoPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "pedido no encontrado")
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

	// Validación de campos obligatorios
	if nuevo.ClienteID == 0 || strings.TrimSpace(nuevo.Fecha) == "" {
		RespondError(w, http.StatusBadRequest, "ClienteID y Fecha son obligatorios")
		return
	}

	// El estado siempre inicia como pendiente
	nuevo.Estado = "pendiente"
	// El total inicia en cero hasta agregar detalles
	nuevo.Total = 0
	nuevo = s.storage.CrearPedido(nuevo)
	RespondJSON(w, http.StatusCreated, nuevo)
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

	// Validamos que al menos un campo venga para actualizar
	if strings.TrimSpace(datos.Estado) == "" && strings.TrimSpace(datos.Fecha) == "" && datos.Total == 0 {
		RespondError(w, http.StatusBadRequest, "debe enviar al menos un campo para actualizar")
		return
	}

	actualizado, encontrado := s.storage.ActualizarPedido(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "pedido no encontrado para actualizar")
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

// EliminarPedido cancela y remueve un pedido por su ID (DELETE)
func (s *Server) EliminarPedido(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	if !s.storage.EliminarPedido(id) {
		RespondError(w, http.StatusNotFound, "pedido no encontrado para eliminar")
		return
	}

	// 204 No Content: eliminación exitosa sin cuerpo de respuesta
	w.WriteHeader(http.StatusNoContent)
}

// -------------------- DETALLES DE PEDIDO --------------------

// ListarDetalles devuelve todos los detalles de pedido registrados (GET)
func (s *Server) ListarDetalles(w http.ResponseWriter, _ *http.Request) {
	detalles := s.storage.ListarDetalles()
	RespondJSON(w, http.StatusOK, detalles)
}

// ObtenerDetalle devuelve un detalle de pedido por su ID (GET)
func (s *Server) ObtenerDetalle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	detalle, encontrado := s.storage.BuscarDetallePorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "detalle de pedido no encontrado")
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

	// Validación de campos obligatorios
	if nuevo.PedidoID == 0 || nuevo.EspecieID == 0 || nuevo.CantidadKg == 0 || nuevo.PrecioUnitario == 0 {
		RespondError(w, http.StatusBadRequest, "PedidoID, EspecieID, CantidadKg y PrecioUnitario son obligatorios")
		return
	}

	// Calculamos el subtotal automáticamente
	nuevo.Subtotal = nuevo.CantidadKg * nuevo.PrecioUnitario
	nuevo = s.storage.CrearDetalle(nuevo)
	RespondJSON(w, http.StatusCreated, nuevo)
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

	// Validamos que al menos un campo venga para actualizar
	if datos.CantidadKg == 0 && datos.PrecioUnitario == 0 && datos.EspecieID == 0 {
		RespondError(w, http.StatusBadRequest, "debe enviar al menos un campo para actualizar")
		return
	}

	// Recalculamos el subtotal si vienen los dos campos necesarios
	if datos.CantidadKg != 0 && datos.PrecioUnitario != 0 {
		datos.Subtotal = datos.CantidadKg * datos.PrecioUnitario
	}

	actualizado, encontrado := s.storage.ActualizarDetalle(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "detalle de pedido no encontrado para actualizar")
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

// EliminarDetalle remueve un detalle de pedido por su ID (DELETE)
func (s *Server) EliminarDetalle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "el id debe ser un número entero")
		return
	}

	if !s.storage.EliminarDetalle(id) {
		RespondError(w, http.StatusNotFound, "detalle de pedido no encontrado para eliminar")
		return
	}

	// 204 No Content: eliminación exitosa sin cuerpo de respuesta
	w.WriteHeader(http.StatusNoContent)
}

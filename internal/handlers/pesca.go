package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"Proyecto_Aplicaciones_Web_II/internal/models"
	"Proyecto_Aplicaciones_Web_II/internal/storage"
)

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

// MontarRutasPesca registra todas las rutas del módulo dentro del subrouter /api/v1.
// main.go llama a esta función y le pasa el store — los handlers no lo conocen directamente.
func MontarRutasPesca(r chi.Router, store storage.AlmacenPesca) {

	// ── Usuarios ─────────────────────────────────────────────────────────
	r.Get("/usuarios", func(w http.ResponseWriter, req *http.Request) {
		GetAllUsuarios(w, req, store)
	})
	r.Post("/usuarios", func(w http.ResponseWriter, req *http.Request) {
		CreateUsuario(w, req, store)
	})
	r.Get("/usuarios/{id}", func(w http.ResponseWriter, req *http.Request) {
		GetUsuario(w, req, store)
	})
	r.Put("/usuarios/{id}", func(w http.ResponseWriter, req *http.Request) {
		UpdateUsuario(w, req, store)
	})
	r.Delete("/usuarios/{id}", func(w http.ResponseWriter, req *http.Request) {
		DeleteUsuario(w, req, store)
	})

	// ── Pescadores ───────────────────────────────────────────────────────
	r.Get("/pescadores", func(w http.ResponseWriter, req *http.Request) {
		GetAllPescadores(w, req, store)
	})
	r.Post("/pescadores", func(w http.ResponseWriter, req *http.Request) {
		CreatePescador(w, req, store)
	})
	r.Get("/pescadores/{id}", func(w http.ResponseWriter, req *http.Request) {
		GetPescador(w, req, store)
	})
	r.Put("/pescadores/{id}", func(w http.ResponseWriter, req *http.Request) {
		UpdatePescador(w, req, store)
	})
	r.Delete("/pescadores/{id}", func(w http.ResponseWriter, req *http.Request) {
		DeletePescador(w, req, store)
	})

	// ── Embarcaciones ────────────────────────────────────────────────────
	r.Get("/embarcaciones", func(w http.ResponseWriter, req *http.Request) {
		GetAllEmbarcaciones(w, req, store)
	})
	r.Post("/embarcaciones", func(w http.ResponseWriter, req *http.Request) {
		CreateEmbarcacion(w, req, store)
	})
	r.Get("/embarcaciones/{id}", func(w http.ResponseWriter, req *http.Request) {
		GetEmbarcacion(w, req, store)
	})
	r.Put("/embarcaciones/{id}", func(w http.ResponseWriter, req *http.Request) {
		UpdateEmbarcacion(w, req, store)
	})
	r.Delete("/embarcaciones/{id}", func(w http.ResponseWriter, req *http.Request) {
		DeleteEmbarcacion(w, req, store)
	})

	// ── Especies ─────────────────────────────────────────────────────────
	r.Get("/especies", func(w http.ResponseWriter, req *http.Request) {
		GetAllEspecies(w, req, store)
	})
	r.Post("/especies", func(w http.ResponseWriter, req *http.Request) {
		CreateEspecie(w, req, store)
	})
	r.Get("/especies/{id}", func(w http.ResponseWriter, req *http.Request) {
		GetEspecie(w, req, store)
	})
	r.Put("/especies/{id}", func(w http.ResponseWriter, req *http.Request) {
		UpdateEspecie(w, req, store)
	})
	r.Delete("/especies/{id}", func(w http.ResponseWriter, req *http.Request) {
		DeleteEspecie(w, req, store)
	})

	// ── Capturas ─────────────────────────────────────────────────────────
	r.Get("/capturas", func(w http.ResponseWriter, req *http.Request) {
		GetAllCapturas(w, req, store)
	})
	r.Post("/capturas", func(w http.ResponseWriter, req *http.Request) {
		CreateCaptura(w, req, store)
	})
	r.Get("/capturas/{id}", func(w http.ResponseWriter, req *http.Request) {
		GetCaptura(w, req, store)
	})
	r.Put("/capturas/{id}", func(w http.ResponseWriter, req *http.Request) {
		UpdateCaptura(w, req, store)
	})
	r.Delete("/capturas/{id}", func(w http.ResponseWriter, req *http.Request) {
		DeleteCaptura(w, req, store)
	})

	// ── Bodegas ──────────────────────────────────────────────────────────
	r.Get("/bodegas", func(w http.ResponseWriter, req *http.Request) {
		GetAllBodegas(w, req, store)
	})
	r.Post("/bodegas", func(w http.ResponseWriter, req *http.Request) {
		CreateBodega(w, req, store)
	})
	r.Get("/bodegas/{id}", func(w http.ResponseWriter, req *http.Request) {
		GetBodega(w, req, store)
	})
	r.Put("/bodegas/{id}", func(w http.ResponseWriter, req *http.Request) {
		UpdateBodega(w, req, store)
	})
	r.Delete("/bodegas/{id}", func(w http.ResponseWriter, req *http.Request) {
		DeleteBodega(w, req, store)
	})

	// ── Stocks ───────────────────────────────────────────────────────────
	r.Get("/stocks", func(w http.ResponseWriter, req *http.Request) {
		GetAllStocks(w, req, store)
	})
	r.Post("/stocks", func(w http.ResponseWriter, req *http.Request) {
		CreateStock(w, req, store)
	})
	r.Get("/stocks/{id}", func(w http.ResponseWriter, req *http.Request) {
		GetStock(w, req, store)
	})
	r.Put("/stocks/{id}", func(w http.ResponseWriter, req *http.Request) {
		UpdateStock(w, req, store)
	})
	r.Delete("/stocks/{id}", func(w http.ResponseWriter, req *http.Request) {
		DeleteStock(w, req, store)
	})
}

// ════════════════════════════ HANDLERS DE USUARIOS ═══════════════════════════

// GetAllUsuarios atiende GET /api/v1/usuarios
func GetAllUsuarios(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	RespondJSON(w, http.StatusOK, store.ListarUsuarios())
}

// GetUsuario atiende GET /api/v1/usuarios/{id}
func GetUsuario(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	usuario, encontrado := store.BuscarUsuarioPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "usuario no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, usuario)
}

// CreateUsuario atiende POST /api/v1/usuarios
// tipo_usuario debe ser "pescador" o "cliente"
func CreateUsuario(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	var body models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(body.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "nombre es requerido")
		return
	}
	if strings.TrimSpace(body.Email) == "" {
		RespondError(w, http.StatusBadRequest, "email es requerido")
		return
	}
	if body.TipoUsuario != "pescador" && body.TipoUsuario != "cliente" {
		RespondError(w, http.StatusBadRequest, "tipo_usuario debe ser 'pescador' o 'cliente'")
		return
	}
	body.Estado = true
	creado := store.CrearUsuario(body)
	RespondJSON(w, http.StatusCreated, creado)
}

// UpdateUsuario atiende PUT /api/v1/usuarios/{id}
func UpdateUsuario(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(body.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "nombre es requerido")
		return
	}
	if body.TipoUsuario != "pescador" && body.TipoUsuario != "cliente" {
		RespondError(w, http.StatusBadRequest, "tipo_usuario debe ser 'pescador' o 'cliente'")
		return
	}
	actualizado, encontrado := store.ActualizarUsuario(id, body)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "usuario no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

// DeleteUsuario atiende DELETE /api/v1/usuarios/{id}
func DeleteUsuario(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if !store.BorrarUsuario(id) {
		RespondError(w, http.StatusNotFound, "usuario no encontrado")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ HANDLERS DE PESCADORES ═════════════════════════

// GetAllPescadores atiende GET /api/v1/pescadores
func GetAllPescadores(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	RespondJSON(w, http.StatusOK, store.ListarPescadores())
}

// GetPescador atiende GET /api/v1/pescadores/{id}
func GetPescador(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	pescador, encontrado := store.BuscarPescadorPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "pescador no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, pescador)
}

// CreatePescador atiende POST /api/v1/pescadores
func CreatePescador(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	var body models.Pescador
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if body.UsuarioID == 0 {
		RespondError(w, http.StatusBadRequest, "usuario_id es requerido")
		return
	}
	if strings.TrimSpace(body.Cedula) == "" {
		RespondError(w, http.StatusBadRequest, "cedula es requerida")
		return
	}
	if strings.TrimSpace(body.PuertoBase) == "" {
		RespondError(w, http.StatusBadRequest, "puerto_base es requerido")
		return
	}
	body.Estado = true
	creado := store.CrearPescador(body)
	RespondJSON(w, http.StatusCreated, creado)
}

// UpdatePescador atiende PUT /api/v1/pescadores/{id}
func UpdatePescador(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Pescador
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(body.Cedula) == "" {
		RespondError(w, http.StatusBadRequest, "cedula es requerida")
		return
	}
	actualizado, encontrado := store.ActualizarPescador(id, body)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "pescador no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

// DeletePescador atiende DELETE /api/v1/pescadores/{id}
func DeletePescador(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if !store.BorrarPescador(id) {
		RespondError(w, http.StatusNotFound, "pescador no encontrado")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ HANDLERS DE EMBARCACIONES ══════════════════════

// GetAllEmbarcaciones atiende GET /api/v1/embarcaciones
func GetAllEmbarcaciones(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	RespondJSON(w, http.StatusOK, store.ListarEmbarcaciones())
}

// GetEmbarcacion atiende GET /api/v1/embarcaciones/{id}
func GetEmbarcacion(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	embarcacion, encontrada := store.BuscarEmbarcacionPorID(id)
	if !encontrada {
		RespondError(w, http.StatusNotFound, "embarcacion no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, embarcacion)
}

// CreateEmbarcacion atiende POST /api/v1/embarcaciones
func CreateEmbarcacion(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	var body models.Embarcacion
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if body.PescadorID == 0 {
		RespondError(w, http.StatusBadRequest, "pescador_id es requerido")
		return
	}
	if strings.TrimSpace(body.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "nombre es requerido")
		return
	}
	if strings.TrimSpace(body.Matricula) == "" {
		RespondError(w, http.StatusBadRequest, "matricula es requerida")
		return
	}
	body.Estado = true
	creada := store.CrearEmbarcacion(body)
	RespondJSON(w, http.StatusCreated, creada)
}

// UpdateEmbarcacion atiende PUT /api/v1/embarcaciones/{id}
func UpdateEmbarcacion(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Embarcacion
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(body.Matricula) == "" {
		RespondError(w, http.StatusBadRequest, "matricula es requerida")
		return
	}
	actualizada, encontrada := store.ActualizarEmbarcacion(id, body)
	if !encontrada {
		RespondError(w, http.StatusNotFound, "embarcacion no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

// DeleteEmbarcacion atiende DELETE /api/v1/embarcaciones/{id}
func DeleteEmbarcacion(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if !store.BorrarEmbarcacion(id) {
		RespondError(w, http.StatusNotFound, "embarcacion no encontrada")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ HANDLERS DE ESPECIES ═══════════════════════════

// GetAllEspecies atiende GET /api/v1/especies
func GetAllEspecies(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	RespondJSON(w, http.StatusOK, store.ListarEspecies())
}

// GetEspecie atiende GET /api/v1/especies/{id}
func GetEspecie(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	especie, encontrada := store.BuscarEspeciePorID(id)
	if !encontrada {
		RespondError(w, http.StatusNotFound, "especie no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, especie)
}

// CreateEspecie atiende POST /api/v1/especies
func CreateEspecie(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	var body models.Especie
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(body.NombreComun) == "" {
		RespondError(w, http.StatusBadRequest, "nombre_comun es requerido")
		return
	}
	if strings.TrimSpace(body.UnidadMedida) == "" {
		RespondError(w, http.StatusBadRequest, "unidad_medida es requerida")
		return
	}
	body.Estado = true
	creada := store.CrearEspecie(body)
	RespondJSON(w, http.StatusCreated, creada)
}

// UpdateEspecie atiende PUT /api/v1/especies/{id}
func UpdateEspecie(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Especie
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(body.NombreComun) == "" {
		RespondError(w, http.StatusBadRequest, "nombre_comun es requerido")
		return
	}
	actualizada, encontrada := store.ActualizarEspecie(id, body)
	if !encontrada {
		RespondError(w, http.StatusNotFound, "especie no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

// DeleteEspecie atiende DELETE /api/v1/especies/{id}
func DeleteEspecie(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if !store.BorrarEspecie(id) {
		RespondError(w, http.StatusNotFound, "especie no encontrada")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ HANDLERS DE CAPTURAS ═══════════════════════════

// GetAllCapturas atiende GET /api/v1/capturas
func GetAllCapturas(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	RespondJSON(w, http.StatusOK, store.ListarCapturas())
}

// GetCaptura atiende GET /api/v1/capturas/{id}
func GetCaptura(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	captura, encontrada := store.BuscarCapturaPorID(id)
	if !encontrada {
		RespondError(w, http.StatusNotFound, "captura no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, captura)
}

// CreateCaptura atiende POST /api/v1/capturas
func CreateCaptura(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	var body models.Captura
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if body.EmbarcacionID == 0 {
		RespondError(w, http.StatusBadRequest, "embarcacion_id es requerido")
		return
	}
	if body.EspecieID == 0 {
		RespondError(w, http.StatusBadRequest, "especie_id es requerido")
		return
	}
	if body.CantidadKG <= 0 {
		RespondError(w, http.StatusBadRequest, "cantidad_kg debe ser mayor a 0")
		return
	}
	if strings.TrimSpace(body.Fecha) == "" {
		RespondError(w, http.StatusBadRequest, "fecha es requerida")
		return
	}
	if body.EstadoFrescura != "fresco" && body.EstadoFrescura != "refrigerado" && body.EstadoFrescura != "congelado" {
		RespondError(w, http.StatusBadRequest, "estado_frescura debe ser 'fresco', 'refrigerado' o 'congelado'")
		return
	}
	creada := store.CrearCaptura(body)
	RespondJSON(w, http.StatusCreated, creada)
}

// UpdateCaptura atiende PUT /api/v1/capturas/{id}
func UpdateCaptura(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Captura
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if body.EmbarcacionID == 0 {
		RespondError(w, http.StatusBadRequest, "embarcacion_id es requerido")
		return
	}
	if body.CantidadKG <= 0 {
		RespondError(w, http.StatusBadRequest, "cantidad_kg debe ser mayor a 0")
		return
	}
	actualizada, encontrada := store.ActualizarCaptura(id, body)
	if !encontrada {
		RespondError(w, http.StatusNotFound, "captura no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

// DeleteCaptura atiende DELETE /api/v1/capturas/{id}
func DeleteCaptura(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if !store.BorrarCaptura(id) {
		RespondError(w, http.StatusNotFound, "captura no encontrada")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ HANDLERS DE BODEGAS ════════════════════════════

// GetAllBodegas atiende GET /api/v1/bodegas
func GetAllBodegas(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	RespondJSON(w, http.StatusOK, store.ListarBodegas())
}

// GetBodega atiende GET /api/v1/bodegas/{id}
func GetBodega(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	bodega, encontrada := store.BuscarBodegaPorID(id)
	if !encontrada {
		RespondError(w, http.StatusNotFound, "bodega no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, bodega)
}

// CreateBodega atiende POST /api/v1/bodegas
func CreateBodega(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	var body models.Bodega
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(body.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "nombre es requerido")
		return
	}
	if strings.TrimSpace(body.Ubicacion) == "" {
		RespondError(w, http.StatusBadRequest, "ubicacion es requerida")
		return
	}
	if body.CapacidadKG <= 0 {
		RespondError(w, http.StatusBadRequest, "capacidad_kg debe ser mayor a 0")
		return
	}
	body.Estado = true
	creada := store.CrearBodega(body)
	RespondJSON(w, http.StatusCreated, creada)
}

// UpdateBodega atiende PUT /api/v1/bodegas/{id}
func UpdateBodega(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Bodega
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(body.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "nombre es requerido")
		return
	}
	actualizada, encontrada := store.ActualizarBodega(id, body)
	if !encontrada {
		RespondError(w, http.StatusNotFound, "bodega no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

// DeleteBodega atiende DELETE /api/v1/bodegas/{id}
func DeleteBodega(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if !store.BorrarBodega(id) {
		RespondError(w, http.StatusNotFound, "bodega no encontrada")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// ════════════════════════════ HANDLERS DE STOCKS ═════════════════════════════

// GetAllStocks atiende GET /api/v1/stocks
func GetAllStocks(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	RespondJSON(w, http.StatusOK, store.ListarStocks())
}

// GetStock atiende GET /api/v1/stocks/{id}
func GetStock(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	stock, encontrado := store.BuscarStockPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "stock no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, stock)
}

// CreateStock atiende POST /api/v1/stocks
func CreateStock(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	var body models.Stock
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if body.BodegaID == 0 {
		RespondError(w, http.StatusBadRequest, "bodega_id es requerido")
		return
	}
	if body.EspecieID == 0 {
		RespondError(w, http.StatusBadRequest, "especie_id es requerido")
		return
	}
	if body.CantidadKG < 0 {
		RespondError(w, http.StatusBadRequest, "cantidad_kg no puede ser negativa")
		return
	}
	if strings.TrimSpace(body.FechaIngreso) == "" {
		RespondError(w, http.StatusBadRequest, "fecha_ingreso es requerida")
		return
	}
	body.Estado = body.CantidadKG > 0
	creado := store.CrearStock(body)
	RespondJSON(w, http.StatusCreated, creado)
}

// UpdateStock atiende PUT /api/v1/stocks/{id}
func UpdateStock(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	var body models.Stock
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if body.BodegaID == 0 {
		RespondError(w, http.StatusBadRequest, "bodega_id es requerido")
		return
	}
	if body.CantidadKG < 0 {
		RespondError(w, http.StatusBadRequest, "cantidad_kg no puede ser negativa")
		return
	}
	// Recalcular estado automáticamente según cantidad
	body.Estado = body.CantidadKG > 0
	actualizado, encontrado := store.ActualizarStock(id, body)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "stock no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

// DeleteStock atiende DELETE /api/v1/stocks/{id}
func DeleteStock(w http.ResponseWriter, r *http.Request, store storage.AlmacenPesca) {
	id, ok := parseID(r)
	if !ok {
		RespondError(w, http.StatusBadRequest, "ID debe ser un número entero positivo")
		return
	}
	if !store.BorrarStock(id) {
		RespondError(w, http.StatusNotFound, "stock no encontrado")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

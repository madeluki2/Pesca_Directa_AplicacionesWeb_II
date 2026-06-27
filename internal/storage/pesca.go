package storage

import (
	"sync"

	"Proyecto_Aplicaciones_Web_II/internal/models"
)

// MemoriaPesca guarda todos los datos en RAM mientras la aplicación esté en ejecución. No es persistente.
type MemoriaPesca struct {
	usuarios      []models.Usuario
	nextUsuarioID int

	pescadores     []models.Pescador
	nextPescadorID int

	embarcaciones     []models.Embarcacion
	nextEmbarcacionID int

	capturas      []models.Captura
	nextCapturaID int

	bodegas      []models.Bodega
	nextBodegaID int

	stocks      []models.Stock
	nextStockID int

	especies      []models.Especie
	nextEspecieID int

	mu sync.Mutex
}

// NuevaMemoriaPesca crea una nueva instancia de MemoriaPesca con datos iniciales vacíos, comenzando los IDs en 1.
func NuevaMemoriaPesca() *MemoriaPesca {
	return &MemoriaPesca{
		usuarios:          []models.Usuario{},
		nextUsuarioID:     1,
		pescadores:        []models.Pescador{},
		nextPescadorID:    1,
		embarcaciones:     []models.Embarcacion{},
		nextEmbarcacionID: 1,
		capturas:          []models.Captura{},
		nextCapturaID:     1,
		bodegas:           []models.Bodega{},
		nextBodegaID:      1,
		stocks:            []models.Stock{},
		nextStockID:       1,
		especies:          []models.Especie{},
		nextEspecieID:     1,
	}
}

// Crearemos métodos para agregar, obtener, actualizar y eliminar cada tipo de entidad (Usuario, Pescador, Embarcacion, Captura, Bodega, Stock, Especie) aquí.

// --------------------------- Usuarios --------------------------
func (m *MemoriaPesca) ListarUsuarios() []models.Usuario {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Usuario, len(m.usuarios))
	copy(copia, m.usuarios)
	return copia
}

func (m *MemoriaPesca) BuscarUsuarioPorID(id int) (models.Usuario, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, u := range m.usuarios {
		if u.ID == id {
			return u, true
		}
	}
	return models.Usuario{}, false
}

func (m *MemoriaPesca) CrearUsuario(u models.Usuario) models.Usuario {
	m.mu.Lock()
	defer m.mu.Unlock()
	u.ID = m.nextUsuarioID
	m.nextUsuarioID++
	m.usuarios = append(m.usuarios, u)
	return u
}

func (m *MemoriaPesca) ActualizarUsuario(id int, datos models.Usuario) (models.Usuario, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, existing := range m.usuarios {
		if existing.ID == id {
			m.usuarios[i] = datos
			return datos, true
		}
	}
	return models.Usuario{}, false
}

func (m *MemoriaPesca) BorrarUsuario(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, u := range m.usuarios {
		if u.ID == id {
			m.usuarios = append(m.usuarios[:i], m.usuarios[i+1:]...)
			return true
		}
	}
	return false
}

// --------------------------- Pescadores --------------------------
func (m *MemoriaPesca) ListarPescadores() []models.Pescador {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Pescador, len(m.pescadores))
	copy(copia, m.pescadores)
	return copia
}

func (m *MemoriaPesca) BuscarPescadorPorID(id int) (models.Pescador, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, p := range m.pescadores {
		if p.ID == id {
			return p, true
		}
	}
	return models.Pescador{}, false
}

func (m *MemoriaPesca) CrearPescador(p models.Pescador) models.Pescador {
	m.mu.Lock()
	defer m.mu.Unlock()
	p.ID = m.nextPescadorID
	m.nextPescadorID++
	m.pescadores = append(m.pescadores, p)
	return p
}

func (m *MemoriaPesca) ActualizarPescador(id int, datos models.Pescador) (models.Pescador, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, existing := range m.pescadores {
		if existing.ID == id {
			m.pescadores[i] = datos
			return datos, true
		}
	}
	return models.Pescador{}, false
}

func (m *MemoriaPesca) BorrarPescador(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.pescadores {
		if p.ID == id {
			m.pescadores = append(m.pescadores[:i], m.pescadores[i+1:]...)
			return true
		}
	}
	return false
}

// --------------------------- Embarcaciones --------------------------
func (m *MemoriaPesca) ListarEmbarcaciones() []models.Embarcacion {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Embarcacion, len(m.embarcaciones))
	copy(copia, m.embarcaciones)
	return copia
}

func (m *MemoriaPesca) BuscarEmbarcacionPorID(id int) (models.Embarcacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, e := range m.embarcaciones {
		if e.ID == id {
			return e, true
		}
	}
	return models.Embarcacion{}, false
}

func (m *MemoriaPesca) CrearEmbarcacion(e models.Embarcacion) models.Embarcacion {
	m.mu.Lock()
	defer m.mu.Unlock()
	e.ID = m.nextEmbarcacionID
	m.nextEmbarcacionID++
	m.embarcaciones = append(m.embarcaciones, e)
	return e
}

func (m *MemoriaPesca) ActualizarEmbarcacion(id int, datos models.Embarcacion) (models.Embarcacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, e := range m.embarcaciones {
		if e.ID == id {
			datos.ID = id
			m.embarcaciones[i] = datos
			return datos, true
		}
	}
	return models.Embarcacion{}, false
}

func (m *MemoriaPesca) BorrarEmbarcacion(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, e := range m.embarcaciones {
		if e.ID == id {
			m.embarcaciones = append(m.embarcaciones[:i], m.embarcaciones[i+1:]...)
			return true
		}
	}
	return false
}

// --------------------------- Capturas --------------------------
func (m *MemoriaPesca) ListarCapturas() []models.Captura {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Captura, len(m.capturas))
	copy(copia, m.capturas)
	return copia
}

func (m *MemoriaPesca) BuscarCapturaPorID(id int) (models.Captura, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, c := range m.capturas {
		if c.ID == id {
			return c, true
		}
	}
	return models.Captura{}, false
}

func (m *MemoriaPesca) CrearCaptura(c models.Captura) models.Captura {
	m.mu.Lock()
	defer m.mu.Unlock()
	c.ID = m.nextCapturaID
	m.nextCapturaID++
	m.capturas = append(m.capturas, c)
	return c
}

func (m *MemoriaPesca) ActualizarCaptura(id int, datos models.Captura) (models.Captura, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.capturas {
		if c.ID == id {
			datos.ID = id
			m.capturas[i] = datos
			return datos, true
		}
	}
	return models.Captura{}, false
}

func (m *MemoriaPesca) BorrarCaptura(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.capturas {
		if c.ID == id {
			m.capturas = append(m.capturas[:i], m.capturas[i+1:]...)
			return true
		}
	}
	return false
}

// --------------------------- Bodegas --------------------------
func (m *MemoriaPesca) ListarBodegas() []models.Bodega {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Bodega, len(m.bodegas))
	copy(copia, m.bodegas)
	return copia
}

func (m *MemoriaPesca) BuscarBodegaPorID(id int) (models.Bodega, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, b := range m.bodegas {
		if b.ID == id {
			return b, true
		}
	}
	return models.Bodega{}, false
}

func (m *MemoriaPesca) CrearBodega(b models.Bodega) models.Bodega {
	m.mu.Lock()
	defer m.mu.Unlock()
	b.ID = m.nextBodegaID
	m.nextBodegaID++
	m.bodegas = append(m.bodegas, b)
	return b
}

func (m *MemoriaPesca) ActualizarBodega(id int, datos models.Bodega) (models.Bodega, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, b := range m.bodegas {
		if b.ID == id {
			datos.ID = id
			m.bodegas[i] = datos
			return datos, true
		}
	}
	return models.Bodega{}, false
}

func (m *MemoriaPesca) BorrarBodega(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, b := range m.bodegas {
		if b.ID == id {
			m.bodegas = append(m.bodegas[:i], m.bodegas[i+1:]...)
			return true
		}
	}
	return false
}

// --------------------------- Stocks --------------------------
func (m *MemoriaPesca) ListarStocks() []models.Stock {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Stock, len(m.stocks))
	copy(copia, m.stocks)
	return copia
}

func (m *MemoriaPesca) BuscarStockPorID(id int) (models.Stock, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, s := range m.stocks {
		if s.ID == id {
			return s, true
		}
	}
	return models.Stock{}, false
}

func (m *MemoriaPesca) CrearStock(s models.Stock) models.Stock {
	m.mu.Lock()
	defer m.mu.Unlock()
	s.ID = m.nextStockID
	m.nextStockID++
	m.stocks = append(m.stocks, s)
	return s
}

func (m *MemoriaPesca) ActualizarStock(id int, datos models.Stock) (models.Stock, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.stocks {
		if s.ID == id {
			datos.ID = id
			m.stocks[i] = datos
			return datos, true
		}
	}
	return models.Stock{}, false
}

func (m *MemoriaPesca) BorrarStock(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.stocks {
		if s.ID == id {
			m.stocks = append(m.stocks[:i], m.stocks[i+1:]...)
			return true
		}
	}
	return false
}

// --------------------------- Especies --------------------------
func (m *MemoriaPesca) ListarEspecies() []models.Especie {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Especie, len(m.especies))
	copy(copia, m.especies)
	return copia
}

func (m *MemoriaPesca) BuscarEspeciePorID(id int) (models.Especie, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, e := range m.especies {
		if e.ID == id {
			return e, true
		}
	}
	return models.Especie{}, false
}

func (m *MemoriaPesca) CrearEspecie(e models.Especie) models.Especie {
	m.mu.Lock()
	defer m.mu.Unlock()
	e.ID = m.nextEspecieID
	m.nextEspecieID++
	m.especies = append(m.especies, e)
	return e
}

func (m *MemoriaPesca) ActualizarEspecie(id int, datos models.Especie) (models.Especie, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, e := range m.especies {
		if e.ID == id {
			datos.ID = id
			m.especies[i] = datos
			return datos, true
		}
	}
	return models.Especie{}, false
}

func (m *MemoriaPesca) BorrarEspecie(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, e := range m.especies {
		if e.ID == id {
			m.especies = append(m.especies[:i], m.especies[i+1:]...)
			return true
		}
	}
	return false
}

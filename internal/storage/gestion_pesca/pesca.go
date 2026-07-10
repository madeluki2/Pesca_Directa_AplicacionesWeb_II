package gestion_pesca

import (
	"sync"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

type MemoriaPesca struct {
	pescadores     []models.Pescador
	embarcaciones  []models.Embarcacion
	especies       []models.Especie
	capturas       []models.Captura
	bodegas        []models.Bodega
	stocks         []models.Stock
	nextPescadorID int
	nextEmbarcaID  int
	nextEspecieID  int
	nextCapturaID  int
	nextBodegaID   int
	nextStockID    int
	mu             sync.Mutex
}

func NuevaMemoriaPesca() *MemoriaPesca {
	return &MemoriaPesca{
		pescadores:     []models.Pescador{},
		embarcaciones:  []models.Embarcacion{},
		especies:       []models.Especie{},
		capturas:       []models.Captura{},
		bodegas:        []models.Bodega{},
		stocks:         []models.Stock{},
		nextPescadorID: 1,
		nextEmbarcaID:  1,
		nextEspecieID:  1,
		nextCapturaID:  1,
		nextBodegaID:   1,
		nextStockID:    1,
	}
}

func (m *MemoriaPesca) ListarPescadores() []models.Pescador {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]models.Pescador(nil), m.pescadores...)
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
	for i := range m.pescadores {
		if m.pescadores[i].ID == id {
			datos.ID = id
			m.pescadores[i] = datos
			return datos, true
		}
	}
	return models.Pescador{}, false
}

func (m *MemoriaPesca) BorrarPescador(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := range m.pescadores {
		if m.pescadores[i].ID == id {
			m.pescadores = append(m.pescadores[:i], m.pescadores[i+1:]...)
			return true
		}
	}
	return false
}

func (m *MemoriaPesca) ListarEmbarcaciones() []models.Embarcacion {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]models.Embarcacion(nil), m.embarcaciones...)
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
	e.ID = m.nextEmbarcaID
	m.nextEmbarcaID++
	m.embarcaciones = append(m.embarcaciones, e)
	return e
}

func (m *MemoriaPesca) ActualizarEmbarcacion(id int, datos models.Embarcacion) (models.Embarcacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := range m.embarcaciones {
		if m.embarcaciones[i].ID == id {
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
	for i := range m.embarcaciones {
		if m.embarcaciones[i].ID == id {
			m.embarcaciones = append(m.embarcaciones[:i], m.embarcaciones[i+1:]...)
			return true
		}
	}
	return false
}

func (m *MemoriaPesca) ListarEspecies() []models.Especie {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]models.Especie(nil), m.especies...)
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
	for i := range m.especies {
		if m.especies[i].ID == id {
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
	for i := range m.especies {
		if m.especies[i].ID == id {
			m.especies = append(m.especies[:i], m.especies[i+1:]...)
			return true
		}
	}
	return false
}

func (m *MemoriaPesca) ListarCapturas() []models.Captura {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]models.Captura(nil), m.capturas...)
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
	for i := range m.capturas {
		if m.capturas[i].ID == id {
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
	for i := range m.capturas {
		if m.capturas[i].ID == id {
			m.capturas = append(m.capturas[:i], m.capturas[i+1:]...)
			return true
		}
	}
	return false
}

func (m *MemoriaPesca) ListarBodegas() []models.Bodega {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]models.Bodega(nil), m.bodegas...)
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
	for i := range m.bodegas {
		if m.bodegas[i].ID == id {
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
	for i := range m.bodegas {
		if m.bodegas[i].ID == id {
			m.bodegas = append(m.bodegas[:i], m.bodegas[i+1:]...)
			return true
		}
	}
	return false
}

func (m *MemoriaPesca) ListarStocks() []models.Stock {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]models.Stock(nil), m.stocks...)
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
	for i := range m.stocks {
		if m.stocks[i].ID == id {
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
	for i := range m.stocks {
		if m.stocks[i].ID == id {
			m.stocks = append(m.stocks[:i], m.stocks[i+1:]...)
			return true
		}
	}
	return false
}

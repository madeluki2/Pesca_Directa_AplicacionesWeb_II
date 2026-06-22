package storage

import (
	"sync"
	"time"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// ══════════════════════════════════════════════
// MemoriaRutas (equivalente a MemoriaPesca)
// ══════════════════════════════════════════════

type MemoriaRutas struct {
	rutas      []models.Ruta
	nextRutaID uint

	puntos      []models.Punto
	nextPuntoID uint

	transportistas []models.Transportista
	nextTransID    uint

	entregas      []models.EntregaPedido
	nextEntregaID uint

	mu sync.RWMutex
}

// NuevaMemoriaRutas crea una nueva instancia con IDs iniciando en 1.
func NuevaMemoriaRutas() *MemoriaRutas {
	return &MemoriaRutas{
		rutas:          []models.Ruta{},
		nextRutaID:     1,
		puntos:         []models.Punto{},
		nextPuntoID:    1,
		transportistas: []models.Transportista{},
		nextTransID:    1,
		entregas:       []models.EntregaPedido{},
		nextEntregaID:  1,
	}
}

// ══════════════════════════════════════════════
// IMPLEMENTACIONES EN MEMORIA
// ══════════════════════════════════════════════

// ── Rutas ─────────────────────────────────────
// ── CRUD de Rutas ─────────────────────────────
func (m *MemoriaRutas) CrearRuta(r models.Ruta) (models.Ruta, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	r.ID = m.nextRutaID
	r.CreadoEn = time.Now()
	m.nextRutaID++
	m.rutas = append(m.rutas, r)
	return r, nil
}

func (m *MemoriaRutas) ObtenerRutas() ([]models.Ruta, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.rutas, nil
}

func (m *MemoriaRutas) ObtenerRutaPorID(id uint) (models.Ruta, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, r := range m.rutas {
		if r.ID == id {
			return r, nil
		}
	}
	return models.Ruta{}, ErrNotFound
}

func (m *MemoriaRutas) ActualizarRuta(id uint, nuevo models.Ruta) (models.Ruta, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, r := range m.rutas {
		if r.ID == id {
			nuevo.ID = id
			nuevo.CreadoEn = r.CreadoEn
			m.rutas[i] = nuevo
			return nuevo, nil
		}
	}
	return models.Ruta{}, ErrNotFound
}

func (m *MemoriaRutas) EliminarRuta(id uint) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, r := range m.rutas {
		if r.ID == id {
			m.rutas = append(m.rutas[:i], m.rutas[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// ── CRUD de Puntos ────────────────────────────
func (m *MemoriaRutas) CrearPunto(p models.Punto) (models.Punto, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	p.ID = m.nextPuntoID
	p.CreadoEn = time.Now()
	m.nextPuntoID++
	m.puntos = append(m.puntos, p)
	return p, nil
}

func (m *MemoriaRutas) ObtenerPuntos() ([]models.Punto, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.puntos, nil
}

func (m *MemoriaRutas) ObtenerPuntoPorID(id uint) (models.Punto, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, p := range m.puntos {
		if p.ID == id {
			return p, nil
		}
	}
	return models.Punto{}, ErrNotFound
}

func (m *MemoriaRutas) ActualizarPunto(id uint, nuevo models.Punto) (models.Punto, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.puntos {
		if p.ID == id {
			nuevo.ID = id
			nuevo.CreadoEn = p.CreadoEn
			m.puntos[i] = nuevo
			return nuevo, nil
		}
	}
	return models.Punto{}, ErrNotFound
}

func (m *MemoriaRutas) EliminarPunto(id uint) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.puntos {
		if p.ID == id {
			m.puntos = append(m.puntos[:i], m.puntos[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// ── CRUD de Transportistas ────────────────────
func (m *MemoriaRutas) CrearTransportista(t models.Transportista) (models.Transportista, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, existing := range m.transportistas {
		if existing.PlacaVehiculo == t.PlacaVehiculo {
			return models.Transportista{}, ErrPlacaDuplicada
		}
	}
	t.ID = m.nextTransID
	t.CreadoEn = time.Now()
	m.nextTransID++
	m.transportistas = append(m.transportistas, t)
	return t, nil
}

func (m *MemoriaRutas) ObtenerTransportistas() ([]models.Transportista, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.transportistas, nil
}

func (m *MemoriaRutas) ObtenerTransportistaPorID(id uint) (models.Transportista, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, t := range m.transportistas {
		if t.ID == id {
			return t, nil
		}
	}
	return models.Transportista{}, ErrNotFound
}

func (m *MemoriaRutas) ActualizarTransportista(id uint, nuevo models.Transportista) (models.Transportista, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, t := range m.transportistas {
		if t.ID == id {
			nuevo.ID = id
			nuevo.CreadoEn = t.CreadoEn
			m.transportistas[i] = nuevo
			return nuevo, nil
		}
	}
	return models.Transportista{}, ErrNotFound
}

func (m *MemoriaRutas) EliminarTransportista(id uint) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, t := range m.transportistas {
		if t.ID == id {
			m.transportistas = append(m.transportistas[:i], m.transportistas[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// ── CRUD de Entregas ──────────────────────────
func (m *MemoriaRutas) CrearEntrega(e models.EntregaPedido) (models.EntregaPedido, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	e.ID = m.nextEntregaID
	e.CreadoEn = time.Now()
	m.nextEntregaID++
	m.entregas = append(m.entregas, e)
	return e, nil
}

func (m *MemoriaRutas) ObtenerEntregas() ([]models.EntregaPedido, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.entregas, nil
}

func (m *MemoriaRutas) ObtenerEntregaPorID(id uint) (models.EntregaPedido, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, e := range m.entregas {
		if e.ID == id {
			return e, nil
		}
	}
	return models.EntregaPedido{}, ErrNotFound
}

func (m *MemoriaRutas) ActualizarEntrega(id uint, nuevo models.EntregaPedido) (models.EntregaPedido, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, e := range m.entregas {
		if e.ID == id {
			nuevo.ID = id
			nuevo.CreadoEn = e.CreadoEn
			m.entregas[i] = nuevo
			return nuevo, nil
		}
	}
	return models.EntregaPedido{}, ErrNotFound
}

func (m *MemoriaRutas) EliminarEntrega(id uint) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, e := range m.entregas {
		if e.ID == id {
			m.entregas = append(m.entregas[:i], m.entregas[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

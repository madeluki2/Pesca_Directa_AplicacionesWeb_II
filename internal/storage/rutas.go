package storage

import (
	"sync"
	"time"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// MemoriaRutas guarda todos los datos en RAM mientras la aplicación
// esté en ejecución. No es persistente.
type MemoriaRutas struct {
	rutas      []models.Ruta
	nextRutaID uint

	puntos      []models.Punto
	nextPuntoID uint

	transportistas []models.Transportista
	nextTransID    uint

	entregas      []models.EntregaPedido
	nextEntregaID uint

	mu sync.Mutex
}

// NuevaMemoriaRutas crea una nueva instancia con datos iniciales
// vacíos, comenzando los IDs en 1.
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

// --------------------------- Rutas --------------------------

func (m *MemoriaRutas) ListarRutas() []models.Ruta {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Ruta, len(m.rutas))
	copy(copia, m.rutas)
	return copia
}

func (m *MemoriaRutas) BuscarRutaPorID(id uint) (models.Ruta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, r := range m.rutas {
		if r.ID == id {
			return r, true
		}
	}
	return models.Ruta{}, false
}

func (m *MemoriaRutas) CrearRuta(r models.Ruta) models.Ruta {
	m.mu.Lock()
	defer m.mu.Unlock()
	r.ID = m.nextRutaID
	r.CreadoEn = time.Now()
	m.nextRutaID++
	m.rutas = append(m.rutas, r)
	return r
}

func (m *MemoriaRutas) ActualizarRuta(id uint, datos models.Ruta) (models.Ruta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, r := range m.rutas {
		if r.ID == id {
			datos.ID = id
			datos.CreadoEn = r.CreadoEn
			m.rutas[i] = datos
			return datos, true
		}
	}
	return models.Ruta{}, false
}

func (m *MemoriaRutas) BorrarRuta(id uint) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, r := range m.rutas {
		if r.ID == id {
			m.rutas = append(m.rutas[:i], m.rutas[i+1:]...)
			return true
		}
	}
	return false
}

// --------------------------- Puntos --------------------------

func (m *MemoriaRutas) ListarPuntos() []models.Punto {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Punto, len(m.puntos))
	copy(copia, m.puntos)
	return copia
}

func (m *MemoriaRutas) BuscarPuntoPorID(id uint) (models.Punto, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, p := range m.puntos {
		if p.ID == id {
			return p, true
		}
	}
	return models.Punto{}, false
}

func (m *MemoriaRutas) CrearPunto(p models.Punto) models.Punto {
	m.mu.Lock()
	defer m.mu.Unlock()
	p.ID = m.nextPuntoID
	p.CreadoEn = time.Now()
	m.nextPuntoID++
	m.puntos = append(m.puntos, p)
	return p
}

func (m *MemoriaRutas) ActualizarPunto(id uint, datos models.Punto) (models.Punto, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.puntos {
		if p.ID == id {
			datos.ID = id
			datos.CreadoEn = p.CreadoEn
			m.puntos[i] = datos
			return datos, true
		}
	}
	return models.Punto{}, false
}

func (m *MemoriaRutas) BorrarPunto(id uint) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.puntos {
		if p.ID == id {
			m.puntos = append(m.puntos[:i], m.puntos[i+1:]...)
			return true
		}
	}
	return false
}

// --------------------------- Transportistas --------------------------

func (m *MemoriaRutas) ListarTransportistas() []models.Transportista {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Transportista, len(m.transportistas))
	copy(copia, m.transportistas)
	return copia
}

func (m *MemoriaRutas) BuscarTransportistaPorID(id uint) (models.Transportista, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, t := range m.transportistas {
		if t.ID == id {
			return t, true
		}
	}
	return models.Transportista{}, false
}

func (m *MemoriaRutas) CrearTransportista(t models.Transportista) models.Transportista {
	m.mu.Lock()
	defer m.mu.Unlock()
	t.ID = m.nextTransID
	t.CreadoEn = time.Now()
	m.nextTransID++
	m.transportistas = append(m.transportistas, t)
	return t
}

func (m *MemoriaRutas) ActualizarTransportista(id uint, datos models.Transportista) (models.Transportista, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, t := range m.transportistas {
		if t.ID == id {
			datos.ID = id
			datos.CreadoEn = t.CreadoEn
			m.transportistas[i] = datos
			return datos, true
		}
	}
	return models.Transportista{}, false
}

func (m *MemoriaRutas) BorrarTransportista(id uint) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, t := range m.transportistas {
		if t.ID == id {
			m.transportistas = append(m.transportistas[:i], m.transportistas[i+1:]...)
			return true
		}
	}
	return false
}

// --------------------------- Entregas --------------------------

func (m *MemoriaRutas) ListarEntregas() []models.EntregaPedido {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.EntregaPedido, len(m.entregas))
	copy(copia, m.entregas)
	return copia
}

func (m *MemoriaRutas) BuscarEntregaPorID(id uint) (models.EntregaPedido, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, e := range m.entregas {
		if e.ID == id {
			return e, true
		}
	}
	return models.EntregaPedido{}, false
}

func (m *MemoriaRutas) CrearEntrega(e models.EntregaPedido) models.EntregaPedido {
	m.mu.Lock()
	defer m.mu.Unlock()
	e.ID = m.nextEntregaID
	e.CreadoEn = time.Now()
	m.nextEntregaID++
	m.entregas = append(m.entregas, e)
	return e
}

func (m *MemoriaRutas) ActualizarEntrega(id uint, datos models.EntregaPedido) (models.EntregaPedido, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, e := range m.entregas {
		if e.ID == id {
			datos.ID = id
			datos.CreadoEn = e.CreadoEn
			m.entregas[i] = datos
			return datos, true
		}
	}
	return models.EntregaPedido{}, false
}

func (m *MemoriaRutas) BorrarEntrega(id uint) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, e := range m.entregas {
		if e.ID == id {
			m.entregas = append(m.entregas[:i], m.entregas[i+1:]...)
			return true
		}
	}
	return false
}

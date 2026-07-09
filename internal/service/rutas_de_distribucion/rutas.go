package rutas_de_distribucion

import (
	"strings"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	. "Pesca_Directa_AplicacionesWeb_II/internal/service"
	storage "Pesca_Directa_AplicacionesWeb_II/internal/storage/rutas_de_distribucion"
)

// RutasService concentra toda la lógica de negocio del módulo de rutas de distribución.
type RutasService struct {
	repo storage.AlmacenRutas
}

func NewRutasService(repo storage.AlmacenRutas) *RutasService {
	return &RutasService{repo: repo}
}

// ═══════════════════════════ RUTAS ═══════════════════════════════════════

func (s *RutasService) ListarRutas() []models.Ruta {
	return s.repo.ListarRutas()
}

func (s *RutasService) ObtenerRuta(id uint) (models.Ruta, error) {
	r, ok := s.repo.BuscarRutaPorID(id)
	if !ok {
		return models.Ruta{}, ErrRutaNoEncontrada
	}
	return r, nil
}

func (s *RutasService) CrearRuta(r models.Ruta) (models.Ruta, error) {
	if err := validarRuta(r); err != nil {
		return models.Ruta{}, err
	}
	if r.Estado == "" {
		r.Estado = "activo"
	}
	return s.repo.CrearRuta(r), nil
}

func (s *RutasService) ActualizarRuta(id uint, datos models.Ruta) (models.Ruta, error) {
	if err := validarRuta(datos); err != nil {
		return models.Ruta{}, err
	}
	r, ok := s.repo.ActualizarRuta(id, datos)
	if !ok {
		return models.Ruta{}, ErrRutaNoEncontrada
	}
	return r, nil
}

func (s *RutasService) BorrarRuta(id uint) error {
	if !s.repo.BorrarRuta(id) {
		return ErrRutaNoEncontrada
	}
	return nil
}

func validarRuta(r models.Ruta) error {
	if strings.TrimSpace(r.Nombre) == "" {
		return ErrNombreVacio
	}
	if strings.TrimSpace(r.Origen) == "" {
		return ErrOrigenVacio
	}
	if strings.TrimSpace(r.Destino) == "" {
		return ErrDestinoVacio
	}
	return nil
}

// ═══════════════════════════ PUNTOS ══════════════════════════════════════

func (s *RutasService) ListarPuntos() []models.Punto {
	return s.repo.ListarPuntos()
}

func (s *RutasService) ObtenerPunto(id uint) (models.Punto, error) {
	p, ok := s.repo.BuscarPuntoPorID(id)
	if !ok {
		return models.Punto{}, ErrPuntoNoEncontrado
	}
	return p, nil
}

func (s *RutasService) CrearPunto(p models.Punto) (models.Punto, error) {
	if err := validarPunto(p); err != nil {
		return models.Punto{}, err
	}
	if p.Estado == "" {
		p.Estado = "activo"
	}
	return s.repo.CrearPunto(p), nil
}

func (s *RutasService) ActualizarPunto(id uint, datos models.Punto) (models.Punto, error) {
	if err := validarPunto(datos); err != nil {
		return models.Punto{}, err
	}
	p, ok := s.repo.ActualizarPunto(id, datos)
	if !ok {
		return models.Punto{}, ErrPuntoNoEncontrado
	}
	return p, nil
}

func (s *RutasService) BorrarPunto(id uint) error {
	if !s.repo.BorrarPunto(id) {
		return ErrPuntoNoEncontrado
	}
	return nil
}

func validarPunto(p models.Punto) error {
	if p.RutaID == 0 {
		return ErrRutaIDVacio
	}
	if strings.TrimSpace(p.Nombre) == "" {
		return ErrNombreVacio
	}
	if strings.TrimSpace(p.Direccion) == "" {
		return ErrDireccionVacia
	}
	return nil
}

// ═══════════════════════════ TRANSPORTISTAS ══════════════════════════════

func (s *RutasService) ListarTransportistas() []models.Transportista {
	return s.repo.ListarTransportistas()
}

func (s *RutasService) ObtenerTransportista(id uint) (models.Transportista, error) {
	t, ok := s.repo.BuscarTransportistaPorID(id)
	if !ok {
		return models.Transportista{}, ErrTransportistaNoEncontrado
	}
	return t, nil
}

func (s *RutasService) CrearTransportista(t models.Transportista) (models.Transportista, error) {
	if err := validarTransportista(t); err != nil {
		return models.Transportista{}, err
	}
	if s.placaYaExiste(t.PlacaVehiculo, 0) {
		return models.Transportista{}, ErrPlacaDuplicada
	}
	if t.Estado == "" {
		t.Estado = "activo"
	}
	return s.repo.CrearTransportista(t), nil
}

func (s *RutasService) ActualizarTransportista(id uint, datos models.Transportista) (models.Transportista, error) {
	if err := validarTransportista(datos); err != nil {
		return models.Transportista{}, err
	}
	if s.placaYaExiste(datos.PlacaVehiculo, id) {
		return models.Transportista{}, ErrPlacaDuplicada
	}
	t, ok := s.repo.ActualizarTransportista(id, datos)
	if !ok {
		return models.Transportista{}, ErrTransportistaNoEncontrado
	}
	return t, nil
}

func (s *RutasService) BorrarTransportista(id uint) error {
	if !s.repo.BorrarTransportista(id) {
		return ErrTransportistaNoEncontrado
	}
	return nil
}

// placaYaExiste recorre los transportistas existentes para detectar una placa
// duplicada. idAIgnorar permite que un Actualizar no choque contra sí mismo.
func (s *RutasService) placaYaExiste(placa string, idAIgnorar uint) bool {
	for _, t := range s.repo.ListarTransportistas() {
		if t.PlacaVehiculo == placa && t.ID != idAIgnorar {
			return true
		}
	}
	return false
}

func validarTransportista(t models.Transportista) error {
	if strings.TrimSpace(t.Nombre) == "" {
		return ErrNombreVacio
	}
	if strings.TrimSpace(t.Telefono) == "" {
		return ErrTelefonoVacio
	}
	if strings.TrimSpace(t.PlacaVehiculo) == "" {
		return ErrPlacaVacia
	}
	return nil
}

// ═══════════════════════════ ENTREGAS ════════════════════════════════════

func (s *RutasService) ListarEntregas() []models.EntregaPedido {
	return s.repo.ListarEntregas()
}

func (s *RutasService) ObtenerEntrega(id uint) (models.EntregaPedido, error) {
	e, ok := s.repo.BuscarEntregaPorID(id)
	if !ok {
		return models.EntregaPedido{}, ErrEntregaNoEncontrada
	}
	return e, nil
}

func (s *RutasService) CrearEntrega(e models.EntregaPedido) (models.EntregaPedido, error) {
	if err := validarEntrega(e); err != nil {
		return models.EntregaPedido{}, err
	}
	if e.Estado == "" {
		e.Estado = "pendiente"
	}
	return s.repo.CrearEntrega(e), nil
}

func (s *RutasService) ActualizarEntrega(id uint, datos models.EntregaPedido) (models.EntregaPedido, error) {
	if err := validarEntrega(datos); err != nil {
		return models.EntregaPedido{}, err
	}
	e, ok := s.repo.ActualizarEntrega(id, datos)
	if !ok {
		return models.EntregaPedido{}, ErrEntregaNoEncontrada
	}
	return e, nil
}

func (s *RutasService) BorrarEntrega(id uint) error {
	if !s.repo.BorrarEntrega(id) {
		return ErrEntregaNoEncontrada
	}
	return nil
}

func validarEntrega(e models.EntregaPedido) error {
	if e.PedidoID == 0 {
		return ErrPedidoIDVacio
	}
	if e.PuntoID == 0 {
		return ErrPuntoIDVacio
	}
	if e.TransportistaID == 0 {
		return ErrTransportistaIDVacio
	}
	return nil
}

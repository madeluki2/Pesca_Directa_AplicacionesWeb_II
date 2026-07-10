package gestion_pesca

import (
	"strings"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	storage "Pesca_Directa_AplicacionesWeb_II/internal/storage/gestion_pesca"
)

// PescaService concentra toda la lógica de negocio del módulo.
type PescaService struct {
	repo storage.AlmacenPesca
}

func NewPescaService(repo storage.AlmacenPesca) *PescaService {
	return &PescaService{repo: repo}
}

// ═══════════════════════════ PESCADORES ══════════════════════════════════════

func (s *PescaService) ListarPescadores() []models.Pescador {
	return s.repo.ListarPescadores()
}

func (s *PescaService) ObtenerPescador(id int) (models.Pescador, error) {
	p, ok := s.repo.BuscarPescadorPorID(id)
	if !ok {
		return models.Pescador{}, ErrPescadorNoEncontrado
	}
	return p, nil
}

func (s *PescaService) CrearPescador(p models.Pescador) (models.Pescador, error) {
	if err := validarPescador(p); err != nil {
		return models.Pescador{}, err
	}
	p.Estado = true
	return s.repo.CrearPescador(p), nil
}

func (s *PescaService) ActualizarPescador(id int, datos models.Pescador) (models.Pescador, error) {
	if err := validarPescador(datos); err != nil {
		return models.Pescador{}, err
	}
	p, ok := s.repo.ActualizarPescador(id, datos)
	if !ok {
		return models.Pescador{}, ErrPescadorNoEncontrado
	}
	return p, nil
}

func (s *PescaService) BorrarPescador(id int) error {
	if !s.repo.BorrarPescador(id) {
		return ErrPescadorNoEncontrado
	}
	return nil
}

func validarPescador(p models.Pescador) error {
	if p.UsuarioID == 0 {
		return ErrCredencialesInvalidas // reutilizamos: usuario_id es requerido
	}
	if strings.TrimSpace(p.Cedula) == "" {
		return ErrCedulaVacia
	}
	if strings.TrimSpace(p.PuertoBase) == "" {
		return ErrPuertoVacio
	}
	return nil
}

// ═══════════════════════════ EMBARCACIONES ════════════════════════════════════

func (s *PescaService) ListarEmbarcaciones() []models.Embarcacion {
	return s.repo.ListarEmbarcaciones()
}

func (s *PescaService) ObtenerEmbarcacion(id int) (models.Embarcacion, error) {
	e, ok := s.repo.BuscarEmbarcacionPorID(id)
	if !ok {
		return models.Embarcacion{}, ErrEmbarcacionNoEncontrada
	}
	return e, nil
}

func (s *PescaService) CrearEmbarcacion(e models.Embarcacion) (models.Embarcacion, error) {
	if err := validarEmbarcacion(e); err != nil {
		return models.Embarcacion{}, err
	}
	e.Estado = true
	return s.repo.CrearEmbarcacion(e), nil
}

func (s *PescaService) ActualizarEmbarcacion(id int, datos models.Embarcacion) (models.Embarcacion, error) {
	if err := validarEmbarcacion(datos); err != nil {
		return models.Embarcacion{}, err
	}
	e, ok := s.repo.ActualizarEmbarcacion(id, datos)
	if !ok {
		return models.Embarcacion{}, ErrEmbarcacionNoEncontrada
	}
	return e, nil
}

func (s *PescaService) BorrarEmbarcacion(id int) error {
	if !s.repo.BorrarEmbarcacion(id) {
		return ErrEmbarcacionNoEncontrada
	}
	return nil
}

func validarEmbarcacion(e models.Embarcacion) error {
	if e.PescadorID == 0 {
		return ErrPescadorNoEncontrado
	}
	if strings.TrimSpace(e.Nombre) == "" {
		return ErrNombreVacio
	}
	if strings.TrimSpace(e.Matricula) == "" {
		return ErrMatriculaVacia
	}
	return nil
}

// ═══════════════════════════ ESPECIES ════════════════════════════════════════

func (s *PescaService) ListarEspecies() []models.Especie {
	return s.repo.ListarEspecies()
}

func (s *PescaService) ObtenerEspecie(id int) (models.Especie, error) {
	e, ok := s.repo.BuscarEspeciePorID(id)
	if !ok {
		return models.Especie{}, ErrEspecieNoEncontrada
	}
	return e, nil
}

func (s *PescaService) CrearEspecie(e models.Especie) (models.Especie, error) {
	if err := validarEspecie(e); err != nil {
		return models.Especie{}, err
	}
	e.Estado = true
	return s.repo.CrearEspecie(e), nil
}

func (s *PescaService) ActualizarEspecie(id int, datos models.Especie) (models.Especie, error) {
	if err := validarEspecie(datos); err != nil {
		return models.Especie{}, err
	}
	e, ok := s.repo.ActualizarEspecie(id, datos)
	if !ok {
		return models.Especie{}, ErrEspecieNoEncontrada
	}
	return e, nil
}

func (s *PescaService) BorrarEspecie(id int) error {
	if !s.repo.BorrarEspecie(id) {
		return ErrEspecieNoEncontrada
	}
	return nil
}

func validarEspecie(e models.Especie) error {
	if strings.TrimSpace(e.NombreComun) == "" {
		return ErrNombreComunVacio
	}
	if strings.TrimSpace(e.UnidadMedida) == "" {
		return ErrUnidadMedidaVacia
	}
	return nil
}

// ═══════════════════════════ CAPTURAS ════════════════════════════════════════

func (s *PescaService) ListarCapturas() []models.Captura {
	return s.repo.ListarCapturas()
}

func (s *PescaService) ObtenerCaptura(id int) (models.Captura, error) {
	c, ok := s.repo.BuscarCapturaPorID(id)
	if !ok {
		return models.Captura{}, ErrCapturaNoEncontrada
	}
	return c, nil
}

func (s *PescaService) CrearCaptura(c models.Captura) (models.Captura, error) {
	if err := validarCaptura(c); err != nil {
		return models.Captura{}, err
	}
	return s.repo.CrearCaptura(c), nil
}

func (s *PescaService) ActualizarCaptura(id int, datos models.Captura) (models.Captura, error) {
	if err := validarCaptura(datos); err != nil {
		return models.Captura{}, err
	}
	c, ok := s.repo.ActualizarCaptura(id, datos)
	if !ok {
		return models.Captura{}, ErrCapturaNoEncontrada
	}
	return c, nil
}

func (s *PescaService) BorrarCaptura(id int) error {
	if !s.repo.BorrarCaptura(id) {
		return ErrCapturaNoEncontrada
	}
	return nil
}

func validarCaptura(c models.Captura) error {
	if c.EmbarcacionID == 0 {
		return ErrEmbarcacionNoEncontrada
	}
	if c.EspecieID == 0 {
		return ErrEspecieNoEncontrada
	}
	if c.CantidadKG <= 0 {
		return ErrCantidadInvalida
	}
	if strings.TrimSpace(c.Fecha) == "" {
		return ErrFechaVacia
	}
	if c.EstadoFrescura != "fresco" && c.EstadoFrescura != "refrigerado" && c.EstadoFrescura != "congelado" {
		return ErrFrescuraInvalida
	}
	return nil
}

// ═══════════════════════════ BODEGAS ═════════════════════════════════════════

func (s *PescaService) ListarBodegas() []models.Bodega {
	return s.repo.ListarBodegas()
}

func (s *PescaService) ObtenerBodega(id int) (models.Bodega, error) {
	b, ok := s.repo.BuscarBodegaPorID(id)
	if !ok {
		return models.Bodega{}, ErrBodegaNoEncontrada
	}
	return b, nil
}

func (s *PescaService) CrearBodega(b models.Bodega) (models.Bodega, error) {
	if err := validarBodega(b); err != nil {
		return models.Bodega{}, err
	}
	b.Estado = true
	return s.repo.CrearBodega(b), nil
}

func (s *PescaService) ActualizarBodega(id int, datos models.Bodega) (models.Bodega, error) {
	if err := validarBodega(datos); err != nil {
		return models.Bodega{}, err
	}
	b, ok := s.repo.ActualizarBodega(id, datos)
	if !ok {
		return models.Bodega{}, ErrBodegaNoEncontrada
	}
	return b, nil
}

func (s *PescaService) BorrarBodega(id int) error {
	if !s.repo.BorrarBodega(id) {
		return ErrBodegaNoEncontrada
	}
	return nil
}

func validarBodega(b models.Bodega) error {
	if strings.TrimSpace(b.Nombre) == "" {
		return ErrNombreVacio
	}
	if strings.TrimSpace(b.Ubicacion) == "" {
		return ErrUbicacionVacia
	}
	if b.CapacidadKG <= 0 {
		return ErrCapacidadInvalida
	}
	return nil
}

// ═══════════════════════════ STOCKS ══════════════════════════════════════════

func (s *PescaService) ListarStocks() []models.Stock {
	return s.repo.ListarStocks()
}

func (s *PescaService) ObtenerStock(id int) (models.Stock, error) {
	st, ok := s.repo.BuscarStockPorID(id)
	if !ok {
		return models.Stock{}, ErrStockNoEncontrado
	}
	return st, nil
}

func (s *PescaService) CrearStock(st models.Stock) (models.Stock, error) {
	if err := validarStock(st); err != nil {
		return models.Stock{}, err
	}
	st.Estado = st.CantidadKG > 0
	return s.repo.CrearStock(st), nil
}

func (s *PescaService) ActualizarStock(id int, datos models.Stock) (models.Stock, error) {
	if err := validarStock(datos); err != nil {
		return models.Stock{}, err
	}
	datos.Estado = datos.CantidadKG > 0
	st, ok := s.repo.ActualizarStock(id, datos)
	if !ok {
		return models.Stock{}, ErrStockNoEncontrado
	}
	return st, nil
}

func (s *PescaService) BorrarStock(id int) error {
	if !s.repo.BorrarStock(id) {
		return ErrStockNoEncontrado
	}
	return nil
}

func validarStock(st models.Stock) error {
	if st.BodegaID == 0 {
		return ErrBodegaNoEncontrada
	}
	if st.EspecieID == 0 {
		return ErrEspecieNoEncontrada
	}
	if st.CantidadKG < 0 {
		return ErrCantidadNegativa
	}
	if strings.TrimSpace(st.FechaIngreso) == "" {
		return ErrFechaIngresoVacia
	}
	return nil
}

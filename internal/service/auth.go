package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

// claveSecreta es la llave con la que se firman y verifican los tokens JWT.
const claveSecreta = "pesca-directa-tarqui-secret-2026"

// AuthService maneja el registro, login y verificación de tokens.
type AuthService struct {
	repo storage.UsuarioRepository
}

// NewAuthService crea un AuthService con el repositorio inyectado.
func NewAuthService(repo storage.UsuarioRepository) *AuthService {
	return &AuthService{repo: repo}
}

// Registrar crea un nuevo usuario con la contraseña hasheada.
func (s *AuthService) Registrar(email, password string) (models.Usuario, error) {
	// Verificamos que el email no esté en uso
	if _, existe := s.repo.BuscarUsuarioPorEmail(email); existe {
		return models.Usuario{}, ErrEmailEnUso
	}

	// Hasheamos la contraseña con bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.Usuario{}, err
	}

	u := models.Usuario{
		Email:        email,
		PasswordHash: string(hash),
	}

	return s.repo.CrearUsuario(u)
}

// Login verifica las credenciales y devuelve un token JWT si son correctas.
func (s *AuthService) Login(email, password string) (string, error) {
	// Buscamos el usuario por email
	u, existe := s.repo.BuscarUsuarioPorEmail(email)
	if !existe {
		return "", ErrCredencialesInvalidas
	}

	// Comparamos la contraseña con el hash guardado
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", ErrCredencialesInvalidas
	}

	// Generamos el token JWT con expiración de 24 horas
	claims := jwt.MapClaims{
		"usuario_id": u.ID,
		"email":      u.Email,
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(claveSecreta))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// VerificarToken valida el token JWT y devuelve los claims si es válido.
func (s *AuthService) VerificarToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		// Verificamos que el método de firma sea HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrCredencialesInvalidas
		}
		return []byte(claveSecreta), nil
	})

	if err != nil || !token.Valid {
		return nil, ErrCredencialesInvalidas
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrCredencialesInvalidas
	}

	return claims, nil
}

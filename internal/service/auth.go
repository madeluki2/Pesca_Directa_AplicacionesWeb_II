package service

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

const secretoPorDefecto = "pesca-directa-tarqui-secret-2026"

const SecretoPorDefecto = secretoPorDefecto

const duracionPorDefecto = 24 * time.Hour

type Claims struct {
	UsuarioID int `json:"usuario_id"`
	jwt.RegisteredClaims
}

type AuthService struct {
	repo     storage.UserRepository
	secreto  []byte
	duracion time.Duration
}

type Opcion func(*AuthService)

// WithSecreto inyecta la clave secreta para firmar tokens JWT.
func WithSecreto(s string) Opcion {
	return func(a *AuthService) {
		if s != "" {
			a.secreto = []byte(s)
		}
	}
}

func WithDuracionToken(d time.Duration) Opcion {
	return func(a *AuthService) {
		if d > 0 {
			a.duracion = d
		}
	}
}

// NewAuthService crea el servicio de autenticación.
func NewAuthService(repo storage.UserRepository, opts ...Opcion) *AuthService {
	a := &AuthService{
		repo:     repo,
		secreto:  []byte(secretoPorDefecto),
		duracion: duracionPorDefecto,
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// Registrar valida los datos, hashea la contraseña y crea el usuario.
func (s *AuthService) Registrar(email, password string) (models.Usuario, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return models.Usuario{}, ErrCredencialesInvalidas
	}
	if _, existe := s.repo.BuscarUsuarioPorEmail(email); existe {
		return models.Usuario{}, ErrEmailEnUso
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.Usuario{}, err
	}

	return s.repo.CrearUsuario(models.Usuario{
		Email:        email,
		PasswordHash: string(hash),
	})
}

// Login verifica las credenciales y devuelve un token JWT firmado.
func (s *AuthService) Login(email, password string) (string, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return "", ErrCredencialesInvalidas
	}

	u, existe := s.repo.BuscarUsuarioPorEmail(email)
	if !existe {
		return "", ErrCredencialesInvalidas
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", ErrCredencialesInvalidas
	}

	return s.GenerarToken(u)
}

func (s *AuthService) GenerarToken(u models.Usuario) (string, error) {
	claims := &Claims{
		UsuarioID: u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.duracion)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secreto)
}

func (s *AuthService) ValidarToken(tokenStr string) (int, error) {
	parsedToken, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrCredencialesInvalidas
		}
		return s.secreto, nil
	})
	if err != nil || !parsedToken.Valid {
		return 0, ErrCredencialesInvalidas
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok {
		return 0, ErrCredencialesInvalidas
	}
	return claims.UsuarioID, nil
}

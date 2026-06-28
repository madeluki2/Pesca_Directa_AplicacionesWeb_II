package storage

import (
	"errors"

	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// UsuarioGORM implementa UsuarioRepository usando GORM + SQLite.
type UsuarioGORM struct {
	db *gorm.DB
}

// NewUsuarioGORM crea un repositorio de usuarios con GORM inyectado.
func NewUsuarioGORM(db *gorm.DB) *UsuarioGORM {
	return &UsuarioGORM{db: db}
}

// CrearUsuario inserta un nuevo usuario en la base de datos.
func (r *UsuarioGORM) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	resultado := r.db.Create(&u)
	if resultado.Error != nil {
		return models.Usuario{}, resultado.Error
	}
	return u, nil
}

// BuscarUsuarioPorEmail busca un usuario por su email (patrón comma-ok).
func (r *UsuarioGORM) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	var u models.Usuario
	resultado := r.db.Where("email = ?", email).First(&u)
	if errors.Is(resultado.Error, gorm.ErrRecordNotFound) {
		return models.Usuario{}, false
	}
	if resultado.Error != nil {
		return models.Usuario{}, false
	}
	return u, true
}

package storage

import (
	"time"

	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

type UserRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

// UsuarioGORM implementa UserRepository usando GORM sobre SQLite.
type UsuarioGORM struct {
	db *gorm.DB
}

func NewUsuarioGORM(db *gorm.DB) *UsuarioGORM {
	return &UsuarioGORM{db: db}
}

func (r *UsuarioGORM) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	u.CreadoEn = time.Now()
	if err := r.db.Create(&u).Error; err != nil {
		return models.Usuario{}, err
	}
	return u, nil
}

func (r *UsuarioGORM) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	var u models.Usuario
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return models.Usuario{}, false
	}
	return u, true
}

var _ UserRepository = (*UsuarioGORM)(nil)

package storage

import (
	"errors"

	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

// UsuarioGORM implementa UserRepository usando GORM/SQLite.
// El registro y login SIEMPRE pasan por aquí, sin importar si el
// resto de la app (Rutas) está corriendo en memoria o en SQLite.
type UsuarioGORM struct {
	db *gorm.DB
}

func NewUsuarioGORM(db *gorm.DB) *UsuarioGORM {
	return &UsuarioGORM{db: db}
}

func (u *UsuarioGORM) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	var usuario models.Usuario
	if err := u.db.Where("email = ?", email).First(&usuario).Error; err != nil {
		return models.Usuario{}, false
	}
	return usuario, true
}

func (u *UsuarioGORM) CrearUsuario(usuario models.Usuario) (models.Usuario, error) {
	if err := u.db.Create(&usuario).Error; err != nil {
		return models.Usuario{}, errors.New("no se pudo crear el usuario: " + err.Error())
	}
	return usuario, nil
}

var _ UserRepository = (*UsuarioGORM)(nil)
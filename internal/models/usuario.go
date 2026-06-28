package models

import "time"

// Usuario representa a cualquier persona registrada en el sistema.
// Solo contiene los campos necesarios para autenticación JWT.
type Usuario struct {
	ID           uint      `json:"id"          gorm:"primaryKey;autoIncrement"`
	Email        string    `json:"email"       gorm:"uniqueIndex;not null"`
	PasswordHash string    `json:"password"    gorm:"not null"`
	CreadoEn     time.Time `json:"creado_en"   gorm:"autoCreateTime"`
}

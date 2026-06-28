package models

import "time"

// Usuario representa una cuenta para autenticación (register/login).
// PasswordHash nunca se serializa en las respuestas JSON (json:"-").
type Usuario struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Email        string    `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string    `json:"-" gorm:"not null"`
	CreadoEn     time.Time `json:"creado_en" gorm:"autoCreateTime"`
}
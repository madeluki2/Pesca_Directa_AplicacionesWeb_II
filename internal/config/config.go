// Package config carga la configuración de la aplicación.
package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config agrupa toda la configuración del servidor en un solo lugar.
type Config struct {
	Puerto      string        // puerto HTTP, ej "8080" o ":8080"
	DBDriver    string        // motor de base de datos: "sqlite" (default, local) o "postgres" (Docker)
	DBDsn       string        // DSN de PostgreSQL (solo se usa si DBDriver="postgres")
	RutaDB      string        // archivo SQLite, ej "pesca.db" (solo si DBDriver="sqlite")
	Storage     string        // "gorm" o "memoria"
	JWTSecreto  string        // Clave secreta para firmar tokens JWT
	JWTDuracion time.Duration // Duración del token, ej: 24h
}

// Cargar lee cada variable de entorno con un valor por defecto seguro.
func Cargar() (Config, error) {
	_ = godotenv.Load()

	duracion, err := parseDuracion(getEnv("JWT_DURACION", "24h"))
	if err != nil {
		return Config{}, fmt.Errorf("JWT_DURACION inválida: %w", err)
	}

	return Config{
		Puerto:      getEnv("PUERTO", "8080"),
		DBDriver:    getEnv("DB_DRIVER", "sqlite"),
		DBDsn:       getEnv("DB_DSN", ""),
		RutaDB:      getEnv("RUTA_DB", "pesca.db"),
		Storage:     getEnv("STORAGE", "gorm"),
		JWTSecreto:  getEnv("JWT_SECRETO", "pesca-directa-tarqui-secret-2026"),
		JWTDuracion: duracion,
	}, nil
}

// getEnv devuelve el valor de la variable de entorno o un fallback por defecto.
func getEnv(clave, porDefecto string) string {
	if v := os.Getenv(clave); v != "" {
		return v
	}
	return porDefecto
}

// parseDuracion convierte strings como "24h" en una estructura de tiempo nativa de Go
func parseDuracion(s string) (time.Duration, error) {
	d, err := time.ParseDuration(s)
	if err == nil {
		return d, nil
	}
	horas, errInt := strconv.Atoi(s)
	if errInt != nil {
		return 0, err
	}
	return time.Duration(horas) * time.Hour, nil
}

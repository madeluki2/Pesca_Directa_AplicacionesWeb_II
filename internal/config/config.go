package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config agrupa toda la configuración de la aplicación.
type Config struct {
	Puerto      string        // Puerto HTTP, ej: "8080"
	RutaDB      string        // Ruta del archivo SQLite, ej: "pesca.db"
	Storage     string        // "gorm" o "memoria"
	JWTSecreto  string        // Clave secreta para firmar tokens JWT
	JWTDuracion time.Duration // Duración del token, ej: 24h
}

func Cargar() (Config, error) {

	_ = godotenv.Load()

	duracion, err := parseDuracion(getEnv("JWT_DURACION", "24h"))
	if err != nil {
		return Config{}, fmt.Errorf("JWT_DURACION inválida: %w", err)
	}

	return Config{
		Puerto:      getEnv("PUERTO", "8080"),
		RutaDB:      getEnv("RUTA_DB", "pesca.db"),
		Storage:     getEnv("STORAGE", "gorm"),
		JWTSecreto:  getEnv("JWT_SECRETO", "pesca-directa-tarqui-secret-2026"),
		JWTDuracion: duracion,
	}, nil
}

// getEnv devuelve el valor de la variable de entorno o el valor por defecto.
func getEnv(clave, porDefecto string) string {
	if v := os.Getenv(clave); v != "" {
		return v
	}
	return porDefecto
}

func parseDuracion(s string) (time.Duration, error) {
	d, err := time.ParseDuration(s)
	if err == nil {
		return d, nil
	}
	// Intento fallback: número de horas sin unidad
	horas, errInt := strconv.Atoi(s)
	if errInt != nil {
		return 0, err
	}
	return time.Duration(horas) * time.Hour, nil
}

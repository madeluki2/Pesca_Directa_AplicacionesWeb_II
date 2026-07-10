<<<<<<< HEAD
// Package config carga la configuracion de la aplicacion desde variables de
// entorno (con soporte para un archivo .env opcional) y expone valores por
// defecto razonables para desarrollo.
package config

import "os"

// Config agrupa toda la configuracion del servidor en un solo lugar.
type Config struct {
	Puerto   string // puerto HTTP, ej ":8080"
	DBDriver string // motor de base de datos: "sqlite" (default, local) o "postgres" (Docker)
	DBDsn    string // DSN de PostgreSQL (solo se usa si DBDriver="postgres")
	RutaDB   string // archivo SQLite, ej "pesca.db" (solo si DBDriver="sqlite")
	Backend  string // "memoria" (datos volatiles) o cualquier otro valor -> GORM
}

// Cargar lee cada variable de entorno con un valor por defecto seguro para
// desarrollo local (sqlite + gorm), igual que antes de dockerizar.
func Cargar() Config {
	return Config{
		Puerto:   conTexto("PUERTO", ":8080"),
		DBDriver: conTexto("DB_DRIVER", "sqlite"),
		DBDsn:    conTexto("DB_DSN", ""),
		RutaDB:   conTexto("RUTA_DB", "pesca.db"),
		Backend:  conTexto("STORAGE", "gorm"),
	}
}

func conTexto(clave, porDefecto string) string {
=======
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

// Cargar inicializa y lee las variables de entorno del archivo .env
func Cargar() (Config, error) {
	// Intenta cargar el archivo .env, si no existe (como en producción con Docker) no pasa nada
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

// getEnv devuelve el valor de la variable de entorno o un fallback por defecto.
func getEnv(clave, porDefecto string) string {
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
	if v := os.Getenv(clave); v != "" {
		return v
	}
	return porDefecto
}
<<<<<<< HEAD
=======

// parseDuracion convierte strings como "24h" en una estructura de tiempo nativa de Go
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
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327

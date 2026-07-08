// Package config carga la configuración de Pesca-Directa Tarqui desde
// variables de entorno (con soporte para un archivo .env opcional) y
// expone valores por defecto razonables para desarrollo.
package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config agrupa TODA la configuración del servidor en un solo lugar.
//
// Antes: el secreto JWT vivía en una var global de service/auth.go, el puerto
// y la ruta de la DB eran literales en main.go, y el backend se leía con
// os.Getenv suelto. Ahora hay UNA sola fuente de verdad.
type Config struct {
	Puerto       string        // puerto HTTP, ej ":8080"
	DBDriver     string        // motor de BD: "sqlite" (default) o "postgres"
	DBDsn        string        // DSN de PostgreSQL (solo si DBDriver="postgres")
	RutaDB       string        // archivo SQLite, ej "pesca.db" (solo si DBDriver="sqlite")
	Backend      string        // "gorm" (default) o "memoria"
	JWTSecreto   []byte        // clave para firmar/verificar JWT
	JWTDuracion  time.Duration // validez del token
	ReadTimeout  time.Duration // timeout de lectura del servidor HTTP
	WriteTimeout time.Duration // timeout de escritura del servidor HTTP
}

// Cargar lee la configuración. Primero intenta cargar un archivo .env (si no
// existe, no es un error: en producción las variables vienen del entorno real).
// Luego lee cada variable con un valor por defecto seguro para desarrollo.
func Cargar() Config {
	// godotenv.Load NO sobreescribe variables ya presentes en el entorno.
	// Si no hay .env, devuelve error que ignoramos a propósito.
	_ = godotenv.Load()

	return Config{
		Puerto:       conTexto("PUERTO", ":8080"),
		DBDriver:     conTexto("DB_DRIVER", "sqlite"),
		DBDsn:        conTexto("DB_DSN", ""),
		RutaDB:       conTexto("RUTA_DB", "pesca.db"),
		Backend:      conTexto("STORAGE", "gorm"),
		JWTSecreto:   []byte(conTexto("JWT_SECRETO", "pesca-directa-tarqui-secreto-solo-dev")),
		JWTDuracion:  conDuracion("JWT_DURACION", 24*time.Hour),
		ReadTimeout:  conDuracion("HTTP_READ_TIMEOUT", 10*time.Second),
		WriteTimeout: conDuracion("HTTP_WRITE_TIMEOUT", 10*time.Second),
	}
}

func conTexto(clave, porDefecto string) string {
	if v := os.Getenv(clave); v != "" {
		return v
	}
	return porDefecto
}

func conDuracion(clave string, porDefecto time.Duration) time.Duration {
	v := os.Getenv(clave)
	if v == "" {
		return porDefecto
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return porDefecto
	}
	return d
}

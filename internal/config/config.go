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
	if v := os.Getenv(clave); v != "" {
		return v
	}
	return porDefecto
}

package gestion_pesca

import "gorm.io/gorm"

func NuevoAlmacenPesca(db *gorm.DB, backend string) AlmacenPesca {
	if backend == "memoria" {
		return NuevaMemoriaPesca()
	}
	return NuevoAlmacenSQLitePesca(db)
}

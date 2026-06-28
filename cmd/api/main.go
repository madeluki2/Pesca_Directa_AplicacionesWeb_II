package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"Pesca_Directa_AplicacionesWeb_II/internal/handlers"
	"Pesca_Directa_AplicacionesWeb_II/internal/middleware"
	"Pesca_Directa_AplicacionesWeb_II/internal/models"
	"Pesca_Directa_AplicacionesWeb_II/internal/service"
	"Pesca_Directa_AplicacionesWeb_II/internal/storage"
)

func main() {
	// 1. Abre la DB y migra las tablas automáticamente.
	gdb, err := gorm.Open(sqlite.Open("rutas.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}
	if err := gdb.AutoMigrate(
		&models.Usuario{},
		&models.Ruta{},
		&models.Punto{},
		&models.Transportista{},
		&models.EntregaPedido{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	// 2. Backend único: GORM + SQLite
	almacen := storage.NuevoAlmacenSQLiteRutas(gdb)
	log.Println("Backend: GORM + SQLite (rutas.db)")

	// 3. UserRepository siempre usa GORM (auth necesita persistencia real).
	usuarioRepo := storage.NewUsuarioGORM(gdb)

	// 4. Services con inyección de dependencias.
	authService := service.NewAuthService(usuarioRepo)
	rutasService := service.NewRutasService(almacen)

	// 5. Server agrupa los services para los handlers.
	servidor := handlers.NewServer(rutasService, authService)

	// 6. Router + middlewares globales.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 7. Rutas versionadas /api/v1/
	r.Route("/api/v1", func(r chi.Router) {

		// Públicas — sin token
		r.Post("/auth/register", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		// Protegidas — requieren: Authorization: Bearer <token>
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			// Rutas
			r.Get("/rutas", servidor.ListarRutas)
			r.Post("/rutas", servidor.CrearRuta)
			r.Get("/rutas/{id}", servidor.ObtenerRuta)
			r.Put("/rutas/{id}", servidor.ActualizarRuta)
			r.Delete("/rutas/{id}", servidor.BorrarRuta)

			// Puntos
			r.Get("/puntos", servidor.ListarPuntos)
			r.Post("/puntos", servidor.CrearPunto)
			r.Get("/puntos/{id}", servidor.ObtenerPunto)
			r.Put("/puntos/{id}", servidor.ActualizarPunto)
			r.Delete("/puntos/{id}", servidor.BorrarPunto)

			// Transportistas
			r.Get("/transportistas", servidor.ListarTransportistas)
			r.Post("/transportistas", servidor.CrearTransportista)
			r.Get("/transportistas/{id}", servidor.ObtenerTransportista)
			r.Put("/transportistas/{id}", servidor.ActualizarTransportista)
			r.Delete("/transportistas/{id}", servidor.BorrarTransportista)

			// Entregas
			r.Get("/entregas", servidor.ListarEntregas)
			r.Post("/entregas", servidor.CrearEntrega)
			r.Get("/entregas/{id}", servidor.ObtenerEntrega)
			r.Put("/entregas/{id}", servidor.ActualizarEntrega)
			r.Delete("/entregas/{id}", servidor.BorrarEntrega)
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

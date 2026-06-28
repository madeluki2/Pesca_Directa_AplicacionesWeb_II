package main

import (
	"log"
	"net/http"
	"os"

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
	// 1. GORM es el DUEÑO DEL ESQUEMA: abre la DB y migra las tablas.
	//    AutoMigrate lee los tags gorm: de cada struct y crea/actualiza
	//    las tablas en pesca.db sin borrar datos existentes.
	gdb, err := gorm.Open(sqlite.Open("pesca.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}
	if err := gdb.AutoMigrate(
		&models.Usuario{},
		&models.Pescador{},
		&models.Embarcacion{},
		&models.Especie{},
		&models.Captura{},
		&models.Bodega{},
		&models.Stock{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	// 2. Elegir el backend que SIRVE las peticiones según la variable STORAGE.
	//    STORAGE=gorm  → usa GORM sobre pesca.db  (por defecto)
	//    STORAGE=memoria → usa almacenamiento en RAM (útil para pruebas rápidas)
	//    >>> Esta es la ÚNICA decisión que cambia entre backends. <<<
	var almacen storage.AlmacenPesca
	switch os.Getenv("STORAGE") {
	case "memoria":
		almacen = storage.NuevaMemoriaPesca()
		log.Println("Backend de almacenamiento: Memoria (datos volátiles)")
	default:
		almacen = storage.NuevoAlmacenSQLitePesca(gdb)
		log.Println("Backend de almacenamiento: GORM + SQLite (pesca.db)")
	}

	// 3. UserRepository siempre usa GORM porque el registro/login
	//    necesita persistencia real (bcrypt + JWT no funcionan en memoria).
	usuarioRepo := storage.NewUsuarioGORM(gdb)

	// 4. Services con inyección de dependencias.
	//    No saben qué backend recibieron — solo conocen la interfaz.
	authService := service.NewAuthService(usuarioRepo)
	pescaService := service.NewPescaService(almacen)

	// 5. Server: punto único de entrada a todos los services para los handlers.
	servidor := handlers.NewServer(pescaService, authService)

	// 6. Router + middlewares globales.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 7. Rutas versionadas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {

		// Rutas públicas — sin token
		r.Post("/auth/register", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		// Rutas protegidas — requieren header: Authorization: Bearer <token>
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			// Pescadores
			r.Get("/pescadores", servidor.ListarPescadores)
			r.Post("/pescadores", servidor.CrearPescador)
			r.Get("/pescadores/{id}", servidor.ObtenerPescador)
			r.Put("/pescadores/{id}", servidor.ActualizarPescador)
			r.Delete("/pescadores/{id}", servidor.BorrarPescador)

			// Embarcaciones
			r.Get("/embarcaciones", servidor.ListarEmbarcaciones)
			r.Post("/embarcaciones", servidor.CrearEmbarcacion)
			r.Get("/embarcaciones/{id}", servidor.ObtenerEmbarcacion)
			r.Put("/embarcaciones/{id}", servidor.ActualizarEmbarcacion)
			r.Delete("/embarcaciones/{id}", servidor.BorrarEmbarcacion)

			// Especies
			r.Get("/especies", servidor.ListarEspecies)
			r.Post("/especies", servidor.CrearEspecie)
			r.Get("/especies/{id}", servidor.ObtenerEspecie)
			r.Put("/especies/{id}", servidor.ActualizarEspecie)
			r.Delete("/especies/{id}", servidor.BorrarEspecie)

			// Capturas
			r.Get("/capturas", servidor.ListarCapturas)
			r.Post("/capturas", servidor.CrearCaptura)
			r.Get("/capturas/{id}", servidor.ObtenerCaptura)
			r.Put("/capturas/{id}", servidor.ActualizarCaptura)
			r.Delete("/capturas/{id}", servidor.BorrarCaptura)

			// Bodegas
			r.Get("/bodegas", servidor.ListarBodegas)
			r.Post("/bodegas", servidor.CrearBodega)
			r.Get("/bodegas/{id}", servidor.ObtenerBodega)
			r.Put("/bodegas/{id}", servidor.ActualizarBodega)
			r.Delete("/bodegas/{id}", servidor.BorrarBodega)

			// Stocks
			r.Get("/stocks", servidor.ListarStocks)
			r.Post("/stocks", servidor.CrearStock)
			r.Get("/stocks/{id}", servidor.ObtenerStock)
			r.Put("/stocks/{id}", servidor.ActualizarStock)
			r.Delete("/stocks/{id}", servidor.BorrarStock)
		})
	})

	log.Println("Servidor Pesca-Directa Tarqui escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

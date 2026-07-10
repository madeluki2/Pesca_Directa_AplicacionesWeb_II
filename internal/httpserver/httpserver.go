package httpserver

import (
	"net/http"
<<<<<<< HEAD
=======
	"strings"
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d
	"time"
)

// Servidor representa un servidor HTTP.
type Servidor struct {
	srv *http.Server
}

type Opcion func(*Servidor)

func Nuevo(handler http.Handler, opts ...Opcion) *Servidor {
	s := &Servidor{
		srv: &http.Server{
			Handler:      handler,
			Addr:         ":8080",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

<<<<<<< HEAD
// ConPuerto cambia el puerto del servidor.
func ConPuerto(puerto string) Opcion {
	return func(s *Servidor) {
		s.srv.Addr = ":" + puerto
=======
// ConPuerto cambia el puerto del servidor de forma segura.
func ConPuerto(puerto string) Opcion {
	return func(s *Servidor) {
		// Si el puerto ya viene con ":", lo asigna directo; si viene plano "8080", se lo agrega.
		if strings.HasPrefix(puerto, ":") {
			s.srv.Addr = puerto
		} else {
			s.srv.Addr = ":" + puerto
		}
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d
	}
}

// ConTimeoutLectura sobreescribe el timeout de lectura.
func ConTimeoutLectura(d time.Duration) Opcion {
	return func(s *Servidor) {
		s.srv.ReadTimeout = d
	}
}

// ConTimeoutEscritura sobreescribe el timeout de escritura.
func ConTimeoutEscritura(d time.Duration) Opcion {
	return func(s *Servidor) {
		s.srv.WriteTimeout = d
	}
}

func (s *Servidor) Iniciar() error {
	return s.srv.ListenAndServe()
}

func (s *Servidor) Detener(ctx interface{ Done() <-chan struct{} }) error {
	if c, ok := ctx.(interface {
		Done() <-chan struct{}
		Err() error
	}); ok {
		_ = c
	}
	return nil
}

func (s *Servidor) Dirección() string {
	return s.srv.Addr
}

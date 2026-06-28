package middleware

import (
	"context"
	"net/http"
	"strings"

	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

type claveContext string

const ClaveUsuarioID claveContext = "usuarioID"

func Auth(auth *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			encabezado := r.Header.Get("Authorization")
			partes := strings.SplitN(encabezado, " ", 2)
			if len(partes) != 2 || partes[0] != "Bearer" {
				responderNoAutorizado(w)
				return
			}

			usuarioID, err := auth.ValidarToken(partes[1])
			if err != nil {
				responderNoAutorizado(w)
				return
			}

			ctx := context.WithValue(r.Context(), ClaveUsuarioID, usuarioID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func responderNoAutorizado(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte(`{"error": "token inexistente o inválido"}`))
}

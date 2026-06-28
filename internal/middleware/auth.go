package middleware

import (
	"net/http"
	"strings"

	"Pesca_Directa_AplicacionesWeb_II/internal/service"
)

// Auth protege las rutas verificando el token JWT en el header Authorization.
func Auth(auth *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// El token viene en el header: Authorization: Bearer <token>
			header := r.Header.Get("Authorization")
			if !strings.HasPrefix(header, "Bearer ") {
				http.Error(w, `{"error":"token inexistente o inválido"}`, http.StatusUnauthorized)
				return
			}

			// Extraemos el token quitando el prefijo "Bearer "
			tokenStr := strings.TrimPrefix(header, "Bearer ")

			// Verificamos que el token sea válido
			if _, err := auth.VerificarToken(tokenStr); err != nil {
				http.Error(w, `{"error":"token inexistente o inválido"}`, http.StatusUnauthorized)
				return
			}

			// Token válido, continuamos con la petición
			next.ServeHTTP(w, r)
		})
	}
}

# ============================================================================
# Build multi-stage: la etapa "builder" compila el binario y la etapa final
# solo lo copia. Resultado: imagen pequeña (~15MB) sin el toolchain de Go.
# ============================================================================

# ---- Etapa 1: builder ----
FROM golang:1.26-alpine AS builder
WORKDIR /src

# Cachear dependencias primero: si go.mod/go.sum no cambian, Docker reutiliza
# esta capa aunque el código fuente cambie → builds mucho más rápidos.
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código y compilar el binario.
# CGO_ENABLED=0: binario estático sin dependencias C.
# Usamos github.com/glebarez/sqlite (pure-Go), así que NO necesitamos CGO.
# GOOS=linux: necesario si estás compilando desde Windows o Mac.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/pesca-api ./cmd/api

# ---- Etapa 2: runner (imagen final mínima) ----
FROM alpine:3.20

# ca-certificates: por si en el futuro se conecta por TLS.
# tzdata: para manejar zonas horarias (Ecuador = America/Guayaquil).
RUN apk add --no-cache ca-certificates tzdata

# Usuario no-root por seguridad: la app no corre como root dentro del contenedor.
RUN adduser -D -u 10001 appuser

WORKDIR /app

# Copiar solo el binario compilado desde la etapa builder.
COPY --from=builder /bin/pesca-api /app/pesca-api

# Crear el directorio donde se guardará pesca.db con los permisos correctos.
RUN mkdir -p /app/data && chown appuser:appuser /app/data

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/pesca-api"]
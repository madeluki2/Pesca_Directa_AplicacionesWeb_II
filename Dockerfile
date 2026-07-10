# Build multi-stage: una etapa "builder" compila el binario y una etapa final
# mínima solo lo copia. Resultado: imagen pequeña y sin el toolchain de Go.

# ---- Etapa 1: builder ----
FROM golang:1.26-alpine AS builder
WORKDIR /src

# Cachear dependencias: copiar primero los módulos y descargar.
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código y compilar.
COPY . .
# CGO_ENABLED=0 produce un binario estático (los drivers de SQLite y Postgres
# que usamos son Go puro, así que no hace falta CGO). GOOS=linux para el runner.
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/pesca-api ./cmd/api

# ---- Etapa 2: runner (imagen final mínima) ----
FROM alpine:3.20
# ca-certificates por si en el futuro se conecta por TLS; tzdata para zonas horarias.
RUN apk add --no-cache ca-certificates tzdata
# Usuario no-root por seguridad.
RUN adduser -D -u 10001 appuser
WORKDIR /app
COPY --from=builder /bin/pesca-api /app/pesca-api
USER appuser
EXPOSE 8080
ENTRYPOINT ["/app/pesca-api"]
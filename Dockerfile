# ---- Etapa 1: builder (Compilación de la aplicación) ----
FROM golang:1.26.4-alpine AS builder

RUN apk add --no-cache git

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/pesca-api ./cmd/api

# ---- Etapa 2: runner (Entorno de ejecución ligero) ----
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata
RUN adduser -D -u 10001 appuser

WORKDIR /app
COPY --from=builder /bin/pesca-api /app/pesca-api

# Crear directorio de datos con permisos para persistencia
RUN mkdir -p /app/data && chown -R appuser:appuser /app/data

USER appuser
EXPOSE 8080

# Comando para arrancar el servidor de Pesca Directa
ENTRYPOINT ["/app/pesca-api"]
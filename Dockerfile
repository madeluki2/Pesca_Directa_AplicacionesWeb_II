# ---- Etapa 1: builder (Compilación de la aplicación) ----
FROM golang:1.26.4-alpine AS builder

# Instalar git por si go mod tidy necesita descargar dependencias de repositorios públicos
RUN apk add --no-cache git

WORKDIR /src

# 1. Copiar primero los archivos de dependencias para aprovechar la caché de Docker
COPY go.mod go.su* ./
RUN go mod download

# 2. Copiar todo el código fuente del proyecto
COPY . .

# 3. Compilar el binario estático de Go desactivando CGO y forzando target Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/pesca-api ./cmd/api

# ---- Etapa 2: runner (Imagen final mínima y segura) ----
FROM alpine:3.20

# ca-certificates: necesario si el backend hace peticiones HTTPS externas
# tzdata: configura la zona horaria de Ecuador (America/Guayaquil)
RUN apk add --no-cache ca-certificates tzdata

# Crear un usuario sin privilegios por seguridad (No-Root)
RUN adduser -D -u 10001 appuser

WORKDIR /app

# Copiar el binario compilado de manera aislada desde la etapa anterior
COPY --from=builder /bin/pesca-api /app/pesca-api

# Crear el directorio donde SQLite generará el archivo 'pesca.db'
# y darle la propiedad al usuario 'appuser' para evitar errores de permisos de escritura
RUN mkdir -p /app/data && chown -R appuser:appuser /app/data

# Cambiar al usuario seguro antes de arrancar
USER appuser

# Exponer el puerto por defecto de tu API
EXPOSE 8080

# Comando para arrancar el servidor de Pesca Directa
ENTRYPOINT ["/app/pesca-api"]
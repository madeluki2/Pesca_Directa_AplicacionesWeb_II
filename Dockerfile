<<<<<<< HEAD
<<<<<<< HEAD
FROM golang:1.26.4-alpine AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/pesca-api ./cmd/api

FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
RUN adduser -D -u 10001 appuser
WORKDIR /app
COPY --from=builder /bin/pesca-api /app/pesca-api
USER appuser
EXPOSE 8080
=======
# ---- Etapa 1: builder (Compilación de la aplicación) ----
=======
# ---- Etapa 1: builder ----
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d
FROM golang:1.26-alpine AS builder

RUN apk add --no-cache git

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/pesca-api ./cmd/api

# ---- Etapa 2: runner ----
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata
RUN adduser -D -u 10001 appuser

WORKDIR /app
COPY --from=builder /bin/pesca-api /app/pesca-api

RUN mkdir -p /app/data && chown -R appuser:appuser /app/data

USER appuser
EXPOSE 8080

<<<<<<< HEAD
# Comando para arrancar el servidor de Pesca Directa
>>>>>>> a7d7cf21cfe890d3e243c29e2cce8961e9021327
ENTRYPOINT ["/app/pesca-api"]
=======
ENTRYPOINT ["/app/pesca-api"]
>>>>>>> 5350001560abd8ae5ce9a208a676c9635fbff78d

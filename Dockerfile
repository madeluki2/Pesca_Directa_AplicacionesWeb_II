# ---- Etapa 1: builder ----
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

ENTRYPOINT ["/app/pesca-api"]

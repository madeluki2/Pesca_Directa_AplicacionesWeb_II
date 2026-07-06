# ============================================================================
# Pesca-Directa Tarqui — Makefile
# Atajos de desarrollo. Uso: make <objetivo>
# ============================================================================

.PHONY: tidy run test cover docker up down logs

tidy:            ## Resolver dependencias y generar go.sum
	go mod tidy

run:             ## Correr local con SQLite (backend por defecto)
	go run ./cmd/api

run-memoria:     ## Correr local con almacenamiento en memoria
	STORAGE=memoria go run ./cmd/api

test:            ## Correr la suite de tests
	go test ./...

cover:           ## Tests con cobertura
	go test ./internal/... -cover

docker:          ## Construir solo la imagen de la API
	docker build -t pesca-directa-api .

up:              ## Levantar la API con docker compose (reconstruye si hay cambios)
	docker compose up --build

down:            ## Bajar los contenedores (mantiene el volumen de datos)
	docker compose down

down-v:          ## Bajar los contenedores Y borrar el volumen de datos
	docker compose down -v

logs:            ## Ver logs en tiempo real de la API
	docker compose logs -f api
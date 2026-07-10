up:
	docker compose up --build

down:
	docker compose down -v

stop:
	docker compose down

logs:
	docker compose logs -f api
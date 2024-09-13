DSN := $(shell godotenv -f .env -- sh -c 'echo $$DB_DSN')

dev:
	@air

db-migrate-up:
	@goose -dir migrations postgres $(DSN) up

db-migrate-down:
	@goose -dir migrations postgres $(DSN) down

db-create-migration:
	@goose -dir migrations create $(name) sql
ifndef name
	$(error name is not set. Usage: make db-create-migration name=<migration_name>)
endif


start-docker:
	@docker-compose --env-file .env -f assets/docker-compose.yaml up
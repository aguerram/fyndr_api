DSN := $(shell godotenv -f .env -- sh -c 'echo $$DB_DSN')

.PHONY: dev debug build test lint clean db-migrate-up db-migrate-down db-create-migration start-docker start-docker-debug

dev:
	@air

debug:
	@RUN_MODE=debug air

build:
	@go build -o ./tmp/main ./src

test:
	@go test ./...

lint:
	@golangci-lint run

clean:
	@rm -rf ./tmp

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
	@docker-compose --env-file .env -f docker-compose.dev.yaml up

re-build-docker:
	@docker-compose --env-file .env -f docker-compose.dev.yaml build

stop-docker:
	@docker-compose --env-file .env -f docker-compose.dev.yaml down

start-docker-debug:
	@docker-compose --env-file .env -f docker-compose.dev.yaml up
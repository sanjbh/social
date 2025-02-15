include .envrc
MIGRATIONS_PATH=./cmd/migrate/migrations

.PHONY: migrate-create migrate-up migrate-down

migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_ADDR) up

migrate-down:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_ADDR) down
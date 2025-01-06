ifneq ("$(wildcard .env)", "")
	include .env
	export $(shell sed 's/=.*//' .env)
endif

# for tidying the import statements
tidy:
	go mod tidy

# Creating New Migrations for up and down
createNewMigration:
	migrate create -ext sql -dir db/migration -seq init-schema

# Apply migrations
migrateup:
	migrate -path ./db/migration -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" -verbose up

# Apply newest migrations
migrateupLast:
	migrate -path ./db/migration -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" -verbose up 1

# Apply Force migrations
migrateupForce:
	migrate -path ./db/migration -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" force 1

# Confirm migrations have been applied
confirmMigrateup:
	docker exec -it simple_banking_db_1 psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) -c '\dt'

# Remove all migrations applied
migratedown:
	migrate -path ./db/migration -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" -verbose down

# Remove Last migration applied
migratedownLast:
	migrate -path ./db/migration -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" -verbose down 1

# Force database version if dirty
forcedatabaseVersion:
	migrate -path ./db/migration -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" force 1

# Generate Go files with SQL queries...
sqlc:
	sqlc generate

# Running Unit test without caching...
test:
	go test -count=1 -v -cover ./...

# Running server command...
server:
	go run main.go

.PHONY: confirmMigrateup migrateup migratedown createNewMigration migrateupForce forcedatabaseVersion sqlc test server migratedownLast migrateupLast tidy

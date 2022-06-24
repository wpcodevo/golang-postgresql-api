# alias sqlc='$(go env GOPATH)/bin/sqlc'
# alias air='$(go env GOPATH)/bin/air'

dev:
	docker-compose up -d

dev-down:
	docker-compose down

go:
	air

migrate:
	migrate create -ext sql -dir db/migrations -seq init_schema

migrate-up:
	migrate -path db/migrations -database "postgresql://admin:password123@localhost:6500/golang_postgres?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migrations -database "postgresql://admin:password123@localhost:6500/golang_postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: sqlc
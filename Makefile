include .env
export

build:
	GOARCH=amd64 go build -o ./bin/main ./cmd/api/main.go

deps:
	go mod tidy

migrate-up:
	migrate -path ./db/migrations -database $(DATABASE_URL) up

migrate-down:
	migrate -path ./db/migrations -database $(DATABASE_URL) down

migrate-drop:
	migrate -path ./db/migrations -database $(DATABASE_URL) drop

include .env
export

deps:
	go mod tidy

run:
	go run ./cmd/api/main.go

build:
	GOARCH=amd64 go build -o ${BINARY_FILE} ${MAIN_GO_FILE}

migrate-up:
	migrate -path ./db/migrations -database $(DATABASE_URL) up

migrate-down:
	migrate -path ./db/migrations -database $(DATABASE_URL) down

migrate-drop:
	migrate -path ./db/migrations -database $(DATABASE_URL) drop

deploy:
	scp -i $(PEM_FILE) ${BINARY_FILE} $(DEPLOY_REMOTE_USER)@$(DEPLOY_REMOTE_HOST):$(DEPLOY_REMOTE_PATH)

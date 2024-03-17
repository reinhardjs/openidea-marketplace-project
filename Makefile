include .env
export

build:
	GOARCH=amd64 go build -o ${BINARY_FILE} ${MAIN_GO_FILE}

deps:
	go mod tidy

migrate-up:
	migrate -path ./db/migrations -database $(DATABASE_URL) up

migrate-down:
	migrate -path ./db/migrations -database $(DATABASE_URL) down

migrate-drop:
	migrate -path ./db/migrations -database $(DATABASE_URL) drop

deploy:
	scp -i $(PEM_FILE) ${BINARY_FILE} $(DEPLOY_REMOTE_USER)@$(DEPLOY_REMOTE_HOST):$(DEPLOY_REMOTE_PATH)

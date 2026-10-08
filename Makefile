.PHONY: up down logs ps test lint fmt migrate build

COMPOSE := docker compose

up:            ## поднять всё и пересобрать образы
	$(COMPOSE) up -d --build

down:          ## остановить (данные в volume сохраняются)
	$(COMPOSE) down

logs:          ## логи; make logs s=chat-server — логи одного сервиса
	$(COMPOSE) logs -f $(s)

ps:
	$(COMPOSE) ps

build:
	go build -o bin/server ./cmd/main.go

test:
	go test ./... -race -cover

fmt:
	go fmt ./...

lint:          ## brew install golangci-lint
	go vet ./...
	golangci-lint run ./...

migrate:       ## прогнать миграции ещё раз
	$(COMPOSE) run --rm migrator

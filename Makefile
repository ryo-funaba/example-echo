.PHONY: help setup build up down restart logs ps setEnv exec build-go clean tidy start boil migrate-create migrate-up migrate-up-n migrate-down migrate-down-n lint dlint lint-docker test

include local.env

BASE_BRANCH="main"
GO_VERSION=1.19.0
TARGET_FILE=./...

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

setup: ## Create a container and migrate db and start a local server
	make build && make up && make migrate-down && make migrate-up && make start

build: setEnv ## Build docker container
	docker compose build --ssh default

up: ## Do docker compose up in detached mode
	docker compose up -d

down: ## Do docker compose down
	docker compose down

restart: down up ## Do docker compose restart

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

setEnv: ## Set Env to use SSH in Docker container
	export COMPOSE_DOCKER_CLI_BUILD=1 export DOCKER_BUILDKIT=1

exec: up ## Execute a command in a running app container
	docker compose exec -it app zsh

build-go: clean ## Build go file
	docker compose exec -it app \
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/app cmd/main.go

clean: ## Remove binary files and cached files
	docker compose exec -it app go clean && rm -rf ./bin

tidy: ## Run go mod tidy
	docker compose exec -it app go mod tidy

start: ## Start a local server
	docker compose exec -it app air -c conf/.air.toml

boil: ## Run SQLBoiler to generate a Go ORM
	docker compose exec -it app sqlboiler mysql -c conf/sqlboiler.toml

migrate-create: ## Create a set of up/down migrations titled $(f)
	docker compose exec -it app migrate create -ext sql -dir db/migrations -seq $(f)

migrate-up: ## Apply all up migrations and run `make boil`
	docker compose exec -it app migrate -database "$(DB_DRIVER)://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)?multiStatements=true" -path=db/migrations/ up && make boil

migrate-up-n: ## Apply $(n) up migrations and run `make boil`
	docker compose exec -it app migrate -database "$(DB_DRIVER)://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)?multiStatements=true" -path=db/migrations/ up $(n) && make boil

migrate-down: ## Apply all down migrations
	docker compose exec -it app migrate -database="$(DB_DRIVER)://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)?multiStatements=true" -path=db/migrations/ down -all

migrate-down-n: ## Apply $(n) down migrations and run `make boil`
	docker compose exec -it app migrate -database="$(DB_DRIVER)://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)?multiStatements=true" -path=db/migrations/ down $(n) && make boil

lint: ## Lint all files
	docker compose exec -it app golangci-lint run --config=conf/.golangci.yml $(TARGET_FILE)

dlint: ## Lint difference files
	docker compose exec -it app golangci-lint run --config=conf/.golangci.yml `$(call diff)`

lint-docker: ## Lint dockerfile
	docker run --rm -i hadolint/hadolint hadolint - --ignore DL4006 --ignore DL3018 < Dockerfile

test: ## Run go test
	docker compose exec -it app zsh -c "go test ${TARGET_FILE}"

define diff ## Get directory with change differences
	git diff --name-only --diff-filter=ACMRT | grep .go$ | xargs -I{} dirname {} | sort | uniq
endef

.PHONY: help setup build build-go clean tidy up down restart exec logs ps setEnv lint dlint test

BASE_BRANCH="main"
GO_VERSION=1.19.0
TARGET_FILE=./...

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

setup: ## Create a container and start a local server
	make build && make up

build: setEnv ## Build docker container
	docker compose build --ssh default

build-go: clean ## Build go file
	docker compose exec -it app \
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/app cmd/main.go

clean: ## remove binary files and cached files
	docker compose exec -it app go clean && rm -rf ./bin

tidy: ## Run go mod tidy
	docker compose exec -it app go mod tidy

up: ## Do docker compose up in detached mode
	docker compose up -d

down: ## Do docker compose down
	docker compose down

restart: down up ## Do docker compose restart

exec: up ## Execute a command in a running app container
	docker compose exec -it app zsh

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

setEnv: ## Set Env to use SSH in Docker container
	export COMPOSE_DOCKER_CLI_BUILD=1 export DOCKER_BUILDKIT=1

lint: ## Lint all files
	docker compose exec -it app golangci-lint run --config=.golangci.yml $(TARGET_FILE)

dlint: ## Lint difference files
	docker compose exec -it app golangci-lint run --config=.golangci.yml `$(call diff)`

test: ## Run go test
	docker compose exec -it app zsh -c "go test ${TARGET_FILE}"

define diff ## Get directory with change differences
	git diff --name-only --diff-filter=ACMRT | grep .go$ | xargs -I{} dirname {} | sort | uniq
endef

.PHONY: help build up down restart exec logs ps setEnv

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: setEnv ## Build docker container
	docker compose build --ssh default

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

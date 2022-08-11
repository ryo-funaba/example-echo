.PHONY: exec build up down restart logs ps help

exec: up ## Execute a command in a running app container
	docker compose exec -it app zsh

build: ## Build docker image to local development
	docker compose up -d --build

up: ## Do docker compose up in detached mode
	docker compose up -d

down: ## Do docker compose down
	docker compose down

restart: down build ## Do docker compose restart

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

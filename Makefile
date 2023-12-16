DOCKER_COMPOSE_CMD := docker compose --env-file env/.env -f docker-compose.yml
ATTACH_APP_CONTAINER := exec -it app
LINT_CMD := golangci-lint run --config=config/.golangci.yml
BASE_BRANCH := "main"
TARGET_FILE := ./...

define diff ## 変更差分のあるファイルを取得する
	git diff --name-only --diff-filter=ACMRT | grep .go$ | xargs -I{} dirname {} | sort | uniq
endef

.PHONY: help ## make task の説明を表示する
help:
	@grep -E "^.PHONY:( *)" $(MAKEFILE_LIST) | sed -e 's/\.PHONY: *//g' | sed -e 's/ *## */\t/g' | awk 'BEGIN {FS = "\t"}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'

#
# Install Tasks
#
.PHONY: install ## 開発環境を構築する
install: env docker.compose.build

#
# Run Tasks
#
.PHONY: run ## 開発サーバーを起動する
run: docker.compose.up hot.reload

#
# Stop Tasks
#
.PHONY: stop ## 開発サーバーを停止する
stop: docker.compose.down

#
# Restart Tasks
#
.PHONY: restart ## 開発サーバーを再起動する
restart: stop run

#
# Refresh Tasks
#
.PHONY: refresh ## 開発サーバー初期化し、起動する
refresh: destroy install run

#
# Destroy Tasks
#
.PHONY: destroy ## 開発サーバーを初期化する
destroy: docker.compose.destroy

#
# Lint Tasks
#
.PHONY: lint ## 全ファイルを対象に Lint を実行する
lint:
	@printf "\033[1;33m[Lint] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) $(ATTACH_APP_CONTAINER) $(LINT_CMD) $(TARGET_FILE)
	@printf "\033[1;32m[Lint] Success\33[0m\n"

.PHONY: lint.diff ## 変更差分のあるファイルを対象に Lint を実行する
lint.diff:
	@printf "\033[1;33m[Lint Diff] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) $(ATTACH_APP_CONTAINER) $(LINT_CMD) `$(call diff)`
	@printf "\033[1;32m[Lint Diff] Success\33[0m\n"

.PHONY: lint.docker ## Dockerfile の Lint を実行する
lint.docker:
	@printf "\033[1;33m[Lint Docker] Start\33[0m\n"; \
	docker run --rm -i hadolint/hadolint hadolint - --ignore DL4006 --ignore DL3018 < Dockerfile
	@printf "\033[1;32m[Lint Docker] Success\33[0m\n"

#
# Test Tasks
#
.PHONY: test.go ## Go ファイルのテストを実行する
test.go:
	@printf "\033[1;33m[Test Go] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) $(ATTACH_APP_CONTAINER) ash -c "go test ${TARGET_FILE}"
	@printf "\033[1;32m[Test Go] Success\33[0m\n"

#
# Go Tasks
#
.PHONY: build.go ## Go ファイルのビルドを行う
build.go:
	@printf "\033[1;33m[Build Go] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) $(ATTACH_APP_CONTAINER) \
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/app cmd/main.go
	@printf "\033[1;32m[Build Go] Success\33[0m\n"

.PHONY: clean ## バイナリとキャッシュを削除する
clean:
	@printf "\033[1;33m[Clean] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) $(ATTACH_APP_CONTAINER) go clean && rm -rf ./bin
	@printf "\033[1;32m[Clean] Success\33[0m\n"

.PHONY: tidy ## go mod tidy を実行する
tidy:
	@printf "\033[1;33m[Tidy] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) $(ATTACH_APP_CONTAINER) go mod tidy
	@printf "\033[1;32m[Tidy] Success\33[0m\n"

#
# Migration Tasks
#
.PHONY: migrate.create ## DB スキーマを変更するための migration ファイルを作成する
migrate.create:
	@printf "\033[1;33m[Migrate Create] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) $(ATTACH_APP_CONTAINER) migrate create -ext sql -dir db/migrations -seq $(f)
	@printf "\033[1;32m[Migrate Create] Success\33[0m\n"

.PHONY: migrate.up ## DB スキーマを最新に更新する
migrate.up:
	@printf "\033[1;33m[Migrate Up] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) $(ATTACH_APP_CONTAINER) migrate -database "$(DB_DRIVER)://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)?multiStatements=true" -path=db/migrations/ up
	@printf "\033[1;32m[Migrate Up] Success\33[0m\n"

.PHONY: migrate.up.n ## DB スキーマを $(n) 件分更新する
migrate.up.n:
	@printf "\033[1;33m[Migrate Up $(n)] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) $(ATTACH_APP_CONTAINER) migrate -database "$(DB_DRIVER)://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)?multiStatements=true" -path=db/migrations/ up $(n)
	@printf "\033[1;32m[Migrate Up $(n)] Success\33[0m\n"

.PHONY: migrate.down ## DB スキーマを全て削除する
migrate.down:
	@printf "\033[1;33m[Migrate Down] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) $(ATTACH_APP_CONTAINER) migrate -database="$(DB_DRIVER)://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)?multiStatements=true" -path=db/migrations/ down -all
	@printf "\033[1;32m[Migrate Down] Success\33[0m\n"

.PHONY: migrate.down.n ## DB スキーマを $(n) 件分削除する
migrate.down.n:
	@printf "\033[1;33m[Migrate Down $(n)] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) $(ATTACH_APP_CONTAINER) migrate -database="$(DB_DRIVER)://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)?multiStatements=true" -path=db/migrations/ down $(n)
	@printf "\033[1;32m[Migrate Down $(n)] Success\33[0m\n"

#
# Other Tasks
#
.PHONY: env ## Docker container をビルドする際に必要な環境変数を設定する
env:
	@printf "\033[1;33m[Env] Start\33[0m\n"; \
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1
	cp env/.env.sample env/.env
	@printf "\033[1;32m[Env] Success\33[0m\n"

.PHONY: hot.reload ## ホットリロードを開始する
hot.reload:
	@printf "\033[1;33m[Hot Reload] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) $(ATTACH_APP_CONTAINER) air -c config/.air.toml
	@printf "\033[1;32m[Hot Reload] Finish\33[0m\n"

.PHONY: exec ## Docker container に入る
exec:
	@printf "\033[1;33m[Exec] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) $(ATTACH_APP_CONTAINER) ash
	@printf "\033[1;32m[Exec] Finish\33[0m\n"

#
# Docker Compose Tasks
#
.PHONY: docker.compose.build ## 開発サーバーの環境構築を行う
docker.compose.build:
	@printf "\033[1;33m[Docker Compose Build] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) build
	@printf "\033[1;32m[Docker Compose Build] Success\33[0m\n"

.PHONY: docker.compose.up ## 開発サーバーを起動する
docker.compose.up:
	@printf "\033[1;33m[Docker Compose Up] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) up -d
	@printf "\033[1;32m[Docker Compose Up] Success\33[0m\n"

.PHONY: docker.compose.down ## 開発サーバーを停止する
docker.compose.down:
	@printf "\033[1;33m[Docker Compose Down] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) down
	@printf "\033[1;32m[Docker Compose Down] Success\33[0m\n"

.PHONY: docker.compose.destroy ## 開発サーバーを初期化する
docker.compose.destroy:
	@printf "\033[1;33m[Docker Compose Destroy] Start\33[0m\n"; \
	$(DOCKER_COMPOSE_CMD) down --rmi all --volumes --remove-orphans
	@printf "\033[1;32m[Docker Compose Destroy] Success\33[0m\n"

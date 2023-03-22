# example_echo

[![CI](https://github.com/ryo-funaba/example-echo/actions/workflows/main.yml/badge.svg)](https://github.com/ryo-funaba/example-echo/actions/workflows/main.yml)

このリポジトリは Go の Web フレームワークである echo のサンプルコードリポジトリです。

## ⚙️技術スタック

- 言語
  - Go@1.19.0
- フレームワーク
  - echo@4.10.0
- パッケージ管理
  - Go Modules
- アーキテクチャ
  - クリーンアーキテクチャ
- Linter
  - golangci-lint
- ORM
  - SQLBoiler
- マイグレーションツール
  - golang-migrate
- 実行環境
  - ローカル環境
    - Docker Compose
  - CI 環境
    - GitHub Actions
  - 本番環境
    - ECS (想定)

## 🔨環境構築

- 前提条件
  - Docker(v20.10.17 以上) がインストールされていること
  - Docker compose(v2.7.0 以上) がインストールされていること
  - Docker デスクトップアプリが起動していること

1. リポジトリをクローン

   ```bash
   git clone github:ryo-funaba/example_echo.git
   ```

2. トップディレクトリに移動

   ```bash
   cd example_echo
   ```

3. API サーバーと DB を起動

   ```bash
   make setup
   ```

   - API サーバー
     - ポート 3000 で起動
     - 接続例：`http://localhost:3000`
   - DB
     - ポート 3306 で起動
     - 接続例：`mysql --host=0.0.0.0 --port=3306 --user=user --password=password`

4. 以上

## 🔍動作確認

|内容|URL|HTTPメソッド|レスポンス|
|----|----|----|----|
|ヘルスチェック|/health_check|GET|{"message":"server is healthy"}|

## 💅Lint の実行

- 全ファイル対象

  ```bash
  make lint
  ```

- 特定ファイル対象

  ```bash
  make lint TARGET_FILE='ファイルの相対パス'
  ```

- 差分対象

  ```bash
  make dlint
  ```

## 🧪テストの実行

- 全ファイル対象

  ```bash
  make test
  ```

- 特定ファイル対象

  ```bash
  make test TARGET_FILE='ファイルの相対パス'
  ```

## 🏗️マイグレーションの実行

- マイグレーションファイルの作成

  ```bash
  make migrate-create f={migration_title}
  ```

- マイグレーション（アップグレード）

  ```bash
  make migrate-up
  ```

- マイグレーション（ダウングレード）

  ```bash
  make migrate-down
  ```

## 💁‍♂️make コマンド一覧

```bash
$ make help

help                 Show options
setup                Create a container and migrate db and start a local server
build                Build docker container
up                   Do docker compose up in detached mode
down                 Do docker compose down
restart              Do docker compose restart
logs                 Tail docker compose logs
ps                   Check container status
setEnv               Set Env to use SSH in Docker container
exec                 Execute a command in a running app container
build-go             Build go file
clean                Remove binary files and cached files
tidy                 Run go mod tidy
start                Start a local server
boil                 Run SQLBoiler to generate a Go ORM
migrate-create       Create a set of timestamped up/down migrations titled $(f)
migrate-up           Apply all up migrations
migrate-up-n         Apply $(n) up migrations
migrate-down         Apply all down migrations
migrate-down-n       Apply $(n) down migrations
lint                 Lint all files
dlint                Lint difference files
test                 Run go test
```

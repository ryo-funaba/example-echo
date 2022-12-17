# example-serverless-go

このリポジトリは Go を使用した Serverless Framework のサンプルコードリポジトリです。

## ⚙️技術スタック

- 言語
  - Go@1.19.0
- フレームワーク
  - Serverless Framework@3.22.0
- パッケージ管理
  - Go Modules
- アーキテクチャ
  - クリーンアーキテクチャ
- Linter
  - golangci-lint
- CI
  - GitHub Actions
- 実行環境
  - ローカル環境
    - Docker Compose
  - 本番環境
    - Lambda

## 🔨環境構築

- 前提条件
  - Docker(v20.10.17 以上) がインストールされていること
  - Docker compose(v2.7.0 以上) がインストールされていること

1. リポジトリをクローン

   ```bash
   git clone github:ryo-funaba/example-serverless-go.git
   ```

2. トップディレクトリに移動

   ```bash
   cd example-serverless-go
   ```

3. Docker image を作成

   ```bash
   make build
   ```

4. Docker Compose を起動

   ```bash
   make up
   ```

5. 以上

## 🔍動作確認

（これから書く）

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

## 💁‍♂️make コマンド一覧

```bash
$ make help

help                 Show options
build                Build docker container
up                   Do docker compose up in detached mode
down                 Do docker compose down
restart              Do docker compose restart
exec                 Execute a command in a running app container
logs                 Tail docker compose logs
ps                   Check container status
setEnv               Set Env to use SSH in Docker container
lint                 Lint all files
dlint                Lint difference files
```

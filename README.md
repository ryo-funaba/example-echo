# example-serverless-go

このリポジトリは Go を使用した Serverless Framework のサンプルコードリポジトリです。

## 技術スタック

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

# 環境構築

- 前提条件

  - GitHub と SSH 接続する設定を済ませていること
    - `ssh-add -A`を実行済みであること
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
   - 初回はキャッシュがないため、1 分ほど時間がかかる想定
   - `go.mod`のライブラリのバージョンを変更した際も image の作成からやり直す
   ```bash
   make build
   ```
4. Docker Compose を起動
   ```bash
   make up
   ```
5. Docker コンテナに入る
   - 基本的にビルドや動作確認はコンテナ内で行う
   ```bash
   make exec
   ```
6. コンテナ内で Shell を再起動
   - szh は `source ~/.zshrc`のエイリアス
   - Shell を再起動すると、プロンプトに現在のブランチ名が表示されるようになる
   ```bash
   szh
   ```
7. 以上

# 動作確認方法

（これから書く）

## Lint の実行

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

# make コマンドについて

`make help`を実行すると、用意された make コマンド一覧を確認できます。

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

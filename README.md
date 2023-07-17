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

3. 開発環境を構築

   ```bash
   make install
   ```

4. API サーバーと DB を起動

   ```bash
   make run
   ```

   - API サーバー
     - ポート 3000 で起動
     - 接続例：`http://localhost:3000`
   - DB
     - ポート 3306 で起動
     - 接続例：`mysql --host=0.0.0.0 --port=3306 --user=user --password=password`

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
  make lint.diff
  ```

## 🧪テストの実行

- 全ファイル対象

  ```bash
  make test.go
  ```

- 特定ファイル対象

  ```bash
  make test.go TARGET_FILE='ファイルの相対パス'
  ```

## 🏗️マイグレーションの実行

- マイグレーションファイルの作成

  ```bash
  make migrate.create f={migration_title}
  ```

- マイグレーション（アップグレード）

  ```bash
  make migrate.up
  ```

- マイグレーション（ダウングレード）

  ```bash
  make migrate.down
  ```

## 💁‍♂️Help

```bash
$ make help

help                                     make task の説明を表示する
install                                  開発環境を構築する
run                                      開発サーバーを起動する
stop                                     開発サーバーを停止する
restart                                  開発サーバーを再起動する
refresh                                  開発サーバー初期化し、起動する
destroy                                  開発サーバーを初期化する
lint                                     全ファイルを対象に Lint を実行する
lint.diff                                変更差分のあるファイルを対象に Lint を実行する
lint.docker                              Dockerfile の Lint を実行する
test.go                                  Go ファイルのテストを実行する
build.go                                 Go ファイルのビルドを行う
clean                                    バイナリとキャッシュを削除する
tidy                                     go mod tidy を実行する
migrate.create                           DB スキーマを変更するための migration ファイルを作成する
migrate.up                               DB スキーマを最新に更新する
migrate.up.n                             DB スキーマを $(n) 件分更新する
migrate.down                             DB スキーマを全て削除する
migrate.down.n                           DB スキーマを $(n) 件分削除する
env                                      Docker container をビルドする際に必要な環境変数を設定する
hot.reload                               ホットリロードを開始する
exec                                     Docker container に入る
docker.compose.build                     開発サーバーの環境構築を行う
docker.compose.up                        開発サーバーを起動する
docker.compose.down                      開発サーバーを停止する
docker.compose.destroy                   開発サーバーを初期化する
```

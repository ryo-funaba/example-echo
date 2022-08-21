# example-serverless-go

このリポジトリは Go を使用した Serverless Framework のサンプルコードリポジトリです。

# 環境構築

- 前提条件

  - Docker(v20.10.17 以上) がインストールされていること
  - Docker compose(v2.7.0 以上) がインストールされていること

- Docker イメージが未作成の場合は作成します。

```bash
make build
```

1. Docker Compose を起動させます。

```bash
make up
```

2. 以上。

# 動作確認方法

（これから書く）

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
```

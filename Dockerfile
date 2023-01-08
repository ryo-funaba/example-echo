# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.19.0-alpine3.16 AS builder

#パッケージをインストール
RUN apk update && \
    apk --no-cache add \
    curl=7.83.1-r5 \
    gcc=11.2.1_git20220219-r2 \
    git=2.36.3-r0 \
    make=4.3-r0 \
    musl-dev=1.2.3-r2 \
    openssh=9.0_p1-r2 \
    tzdata=2022f-r1 \
    vim=8.2.5000-r0 \
    zsh=5.8.1-r4

# タイムゾーンの設定
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime


# 非rootユーザーを作成
ENV USER tempUser
ENV HOME /home/$USER
RUN addgroup -S $USER && \
    adduser -S -G $USER $USER && \
    chown -R $USER:$USER $HOME
USER $USER

WORKDIR $HOME

# .zshrcをコピー
COPY .zshrc ./

# プラグインをインストール
RUN git clone https://github.com/zsh-users/zsh-autosuggestions ~/.zsh/zsh-autosuggestions && \
    git clone https://github.com/zsh-users/zsh-completions.git ~/.zsh/zsh-completions && \
    git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ~/.zsh/zsh-syntax-highlighting && \
    curl -o ~/.zsh/git-prompt.sh https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh && \
    curl -o _git https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.zsh && \
    curl -o ~/.zsh/git-completion.bash https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.bash

# golangci-lintとAirをインストール
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1 && \
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# GitHubにSSH接続するための設定
RUN mkdir -m 700 ~/.ssh && \
    ssh-keyscan github.com > ~/.ssh/known_hosts && \
    git config --global url.git@github.com:.insteadOf https://github.com/

# モジュールをインストール
COPY --chown=$USER:$USER go.mod .
RUN go mod download

# goファイルをビルド
COPY --chown=$USER:$USER . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o bin/app cmd/main.go

##
## Deploy
##
FROM scratch as deploy

# 非rootユーザーを使用
USER $USER

# タイムゾーンの設定
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
ENV TZ=Asia/Tokyo

# バイナリファイルをコピー
COPY --from=builder /home/tempUser/bin .

CMD ["./app"]

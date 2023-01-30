# syntax=docker/dockerfile:1

##
## Builder
##
FROM golang:1.19.0-alpine3.16 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o bin/app cmd/main.go

##
## Developer
##
FROM golang:1.19.0-alpine3.16 AS developer

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

COPY .zshrc /root/

WORKDIR /app

# プラグインをインストール
RUN git clone https://github.com/zsh-users/zsh-autosuggestions ~/.zsh/zsh-autosuggestions && \
    git clone https://github.com/zsh-users/zsh-completions.git ~/.zsh/zsh-completions && \
    git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ~/.zsh/zsh-syntax-highlighting && \
    curl -o ~/.zsh/git-prompt.sh https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh && \
    curl -o _git https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.zsh && \
    curl -o ~/.zsh/git-completion.bash https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.bash

# golangci-lint・Airをインストール
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)"/bin v1.50.1 && \
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b "$(go env GOPATH)"/bin

# sqlboiler・golang-migrateをインストール
RUN go install github.com/volatiletech/sqlboiler/v4@latest && \
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest && \
go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest

##
## Deploy
##
FROM scratch AS deploy

# タイムゾーンの設定
COPY --from=developer /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=developer /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
ENV TZ=Asia/Tokyo

WORKDIR /app

COPY --from=builder /app/bin ./

CMD ["./app"]

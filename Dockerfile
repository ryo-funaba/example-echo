# syntax=docker/dockerfile:1

##
## Builder
##
FROM golang:1.20.1-alpine3.17 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o bin/app cmd/main.go

##
## Developer
##
FROM golang:1.20.1-alpine3.17 AS developer

ARG AIR_VERSION=v1.41.0
ARG GOLANG_CI_LINT_VERSION=v1.51.1
ARG GOLANG_MIGRATE_VERSION=v4.15.2
ARG SQLBOILER_VERSION=v4.14.1

RUN apk update && \
    apk --no-cache add \
    curl=7.88.1-r0 \
    gcc=12.2.1_git20220924-r4 \
    git=2.38.4-r1 \
    make=4.3-r1 \
    musl-dev=1.2.3-r4 \
    openssh=9.1_p1-r2 \
    tzdata=2022f-r1 \
    vim=9.0.0999-r0 \
    zsh=5.9-r0

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
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)"/bin ${GOLANG_CI_LINT_VERSION} && \
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b "$(go env GOPATH)"/bin ${AIR_VERSION}

# sqlboiler・golang-migrateをインストール
RUN go install github.com/volatiletech/sqlboiler/v4@${SQLBOILER_VERSION} && \
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@${SQLBOILER_VERSION} && \
go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@${GOLANG_MIGRATE_VERSION}

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

# syntax=docker/dockerfile:1@sha256:fe40cf4e92cd0c467be2cfc30657a680ae2398318afd50b0c80585784c604f28

#
# Developer
#
FROM golang:1.21.5-alpine3.17@sha256:92cb87af996ec6befc85f0aec27e12ead2fab396695fa8a7abff79e021e58195 AS developer

WORKDIR /app

RUN apk update && \
    apk --no-cache add \
    tzdata

# タイムゾーンを Asia/Tokyo に設定
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

# 開発用ツールをインストール
RUN go install github.com/cosmtrek/air@latest && \
    go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
    wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)"/bin latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o bin/app cmd/main.go

#
# Deploy
#
FROM scratch AS deploy

# タイムゾーンを設定
COPY --from=developer /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=developer /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
ENV TZ=Asia/Tokyo

WORKDIR /app

COPY --from=developer /app/bin ./

CMD ["./app"]

# syntax=docker/dockerfile:1

#
# Developer
#
FROM golang:1.20.1-alpine3.17 AS developer

WORKDIR /app

RUN apk update && \
    apk --no-cache add \
    tzdata

# タイムゾーンを Asia/Tokyo に設定
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

# 開発用ツールをインストール
RUN go install github.com/cosmtrek/air@latest && \
    go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
    wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest

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

COPY --from=builder /app/bin ./

CMD ["./app"]

# TG DUMB BOT

[![test](https://github.com/hackroid/tg-dumb-bot/actions/workflows/pull_request_closed.yml/badge.svg)](https://github.com/hackroid/tg-dumb-bot/actions/workflows/pull_request_closed.yml)

This is a funny tg dumb bot.

## Usage

### Build from source

#### Build with Go

Dependencies

```shell
go get github.com/go-telegram-bot-api/telegram-bot-api/v5
go get github.com/joho/godotenv
```

Paste your token into `.env` like:

```text
TELEGRAM_APITOKEN=123456:ABCDEFGHIJKLMN
DEBUG=1
```

Then

```shell
rm -rf bin
go mod tidy
go build -o bin/ ./app/main
nohup ./bin/main > ./test.log 2>&1 &
```

#### Build with Docker

```shell
docker build . -t test
# run on detach
docker run -d -v /etc/localtime:/etc/localtime:ro -e token="123456:ABCDEFGHIJKLMN" test
# run in foreground
docker run -it -v /etc/localtime:/etc/localtime:ro -e token="123456:ABCDEFGHIJKLMN" test
```

### Use DockerHub Image

```shell
docker pull hackroid/tg-dumb-bot:latest
docker run --name bot-one -d -v /etc/localtime:/etc/localtime:ro -e token="123456:ABCDEFGHIJKLMN" hackroid/tg-dumb-bot:latest
```


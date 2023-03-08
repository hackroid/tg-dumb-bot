# TG DUMB BOT

## Dependencies

```shell
go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
go get github.com/joho/godotenv
```

## Usage

Paste your token into `.env` like:

```text
TELEGRAM_APITOKEN=123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ
DEBUG=1
```

Then

### Build from source

```shell
rm -rf bin
go mod tidy
go build -o bin/ ./main
nohup ./bin/main > ./test.log 2>&1 &
```

### Use Docker

> CI/CD on constructing

```shell
docker compose up --build
```

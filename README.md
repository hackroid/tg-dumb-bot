# TG DUMB BOT

### Dependencies

```shell
go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
go get github.com/joho/godotenv
```

### Usage

Paste your token into `.env` like:

```text
TELEGRAM_APITOKEN=123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ
```

Then

```shell
go mod tidy
go run main/main.go
```


# TG DUMB BOT

## Dependencies

```shell
go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
go get github.com/joho/godotenv
```

## Usage

Paste your token into `.env` like:

```text
TELEGRAM_APITOKEN=123456:ABCDEFGHIJKLMN
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
docker build . -t test
# run on detach
docker run -d -v /etc/localtime:/etc/localtime:ro -e token="123456:ABCDEFGHIJKLMN" test
# run in foreground
docker run -it -v /etc/localtime:/etc/localtime:ro -e token="123456:ABCDEFGHIJKLMN" test
```

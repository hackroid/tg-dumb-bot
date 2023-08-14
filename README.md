# TG DUMB BOT

[![pull_request_closed](https://img.shields.io/github/actions/workflow/status/hackroid/tg-dumb-bot/release.yml)](https://github.com/hackroid/tg-dumb-bot/actions/workflows/pull_request_closed.yml) [![docker_pull](https://img.shields.io/docker/pulls/hackroid/tg-dumb-bot)](https://hub.docker.com/repository/docker/hackroid/tg-dumb-bot) [![issue](https://img.shields.io/github/issues/hackroid/tg-dumb-bot)](https://github.com/hackroid/tg-dumb-bot/issues) [![license](https://img.shields.io/github/license/hackroid/tg-dumb-bot)](https://github.com/hackroid/tg-dumb-bot/blob/main/LICENSE) ![last_commit](https://img.shields.io/github/last-commit/hackroid/tg-dumb-bot?color=red)

This is a funny tg dumb bot.

## Features

> **Warning**  
> All code only for reference.  
> Check CAREFULLY before deployment.

♻️ Using `weibo` command can feed you the TOP10 RUBBISH on Weibo platform RIGHT NOW!

🎲 The `choice` command gives you a random choice between serveral space-sperated options.

🤡 This bot will also randomly reply "Surprise Sunshine Boy!" to a new arrived message in the group with a probability of 0.15.

Propose your function now with a single PR.

## Usage

### Build from source

#### Build with Go

Dependencies

```shell
go get github.com/go-telegram-bot-api/telegram-bot-api/v5
go get github.com/joho/godotenv
go get github.com/gocolly/colly/v2
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

## Other

Test for signing commit

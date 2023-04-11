#!/bin/sh
rm -rf bin
rm -rf .env
go get github.com/go-telegram-bot-api/telegram-bot-api/v5
go get github.com/joho/godotenv
go get github.com/gocolly/colly/v2
go mod tidy
go build -o bin/ ./app/main/
touch .env
echo "TELEGRAM_APITOKEN=$1" >> .env
echo "DEBUG=$2" >> .env
./bin/main
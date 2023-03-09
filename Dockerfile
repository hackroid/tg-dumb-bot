FROM golang:1.20.0-alpine3.17

ENV token="123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ"
ENV debug="1"

LABEL maintainer="Takayama"

COPY . /usr/local/go/src/BOT

WORKDIR /usr/local/go/src/BOT

RUN rm -rf bin && \
    rm -rf .env && \
    go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5 && \
    go get github.com/joho/godotenv && \
    go mod tidy && \
    go build -o bin/ ./main/ && \
    touch .env && \
    chmod +x run.sh

CMD sh run.sh $token $debug

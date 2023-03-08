FROM houxianyj/go1.20.0-alpine3.17:1.0

LABEL maintainer="Takayama"

COPY . /usr/local/go/src/BOT

WORKDIR /usr/local/go/src/BOT

RUN go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5 && \
    go get github.com/joho/godotenv && \
    mkdir bin && \
    cp .env bin/ && \
    go mod tidy && \
    go build -o bin/ ./main/

CMD ["./bin/main"]
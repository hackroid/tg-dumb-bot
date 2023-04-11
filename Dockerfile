FROM golang:1.20.2-bullseye

ENV token="123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ"
ENV debug="1"

LABEL maintainer="Takayama"

COPY . /usr/local/go/src/BOT

WORKDIR /usr/local/go/src/BOT

RUN chmod +x run.sh

CMD sh run.sh $token $debug

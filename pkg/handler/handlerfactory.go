package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHandlerFactory interface {
	Extract()
	Generate()
	Pack()
	Send(chan tgbotapi.MessageConfig)
}

func GetMessageHandler(msg *tgbotapi.Message) (MessageHandlerFactory, error) {
	if len(msg.Text) == 0 {
		return nil, nil
	}
	if msg.IsCommand() {
		return &CommandMessageHandler{recvMsg: msg, reply: false, ok: false}, nil
	}

	return &TextMessageHandler{recvMsg: msg, reply: false, ok: false}, nil
}

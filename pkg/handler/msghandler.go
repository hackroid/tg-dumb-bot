package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hackroid/tg-dumb-bot/pkg/datatype"
	"github.com/hackroid/tg-dumb-bot/pkg/static"
	"log"
)

type TextMessageHandler struct {
	recvMsg     *tgbotapi.Message
	recvContent datatype.TextContentRecv
	reply       bool
	ok          bool
	respContent datatype.TextContentResp
	respMsg     tgbotapi.MessageConfig
}

func (h *TextMessageHandler) Extract() {
	h.recvContent.Text = h.recvMsg.Text
	h.reply = true
}

func (h *TextMessageHandler) Generate() {
	h.respContent.Text, h.ok, _ = static.AddGouBa(h.recvContent)
}

func (h *TextMessageHandler) Pack() {
	if !h.ok {
		return
	}
	// Start packing you msg
	h.respMsg = tgbotapi.NewMessage(h.recvMsg.Chat.ID, h.respContent.Text)

	// End packing
	if h.reply {
		h.respMsg.ReplyToMessageID = h.recvMsg.MessageID
	}
}

func (h *TextMessageHandler) Send(ch chan tgbotapi.MessageConfig) {
	if h.ok {
		ch <- h.respMsg
	}
}

type CommandMessageHandler struct {
	recvMsg     *tgbotapi.Message
	recvContent datatype.CommandContentRecv
	reply       bool
	ok          bool
	respContent datatype.CommandContentResp
	respMsg     tgbotapi.MessageConfig
}

func (h *CommandMessageHandler) Extract() {
	h.recvContent.Text = h.recvMsg.Text
	h.recvContent.Cmd = h.recvMsg.Command()
	h.reply = true
	h.ok = true
}

func (h *CommandMessageHandler) Generate() {
	h.respContent.Text, h.ok, _ = static.NormalCommandMessage(h.recvContent)
}

func (h *CommandMessageHandler) Pack() {
	if !h.ok {
		return
	}
	// Start packing you msg
	h.respMsg = tgbotapi.NewMessage(h.recvMsg.Chat.ID, h.respContent.Text)

	// End packing
	if h.reply {
		h.respMsg.ReplyToMessageID = h.recvMsg.MessageID
	}
}

func (h *CommandMessageHandler) Send(ch chan tgbotapi.MessageConfig) {
	if h.ok {
		log.Printf("[Replyed]\n")
		ch <- h.respMsg
	}
}

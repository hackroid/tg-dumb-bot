package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hackroid/tg-dumb-bot/pkg/datatype"
	"github.com/hackroid/tg-dumb-bot/pkg/static"
	"log"
)

// TextMessageHandler is for handling normal pure text message
type TextMessageHandler struct {
	recvMsg     *tgbotapi.Message        // recvMsg is reference of received tgbotapi.Message
	recvContent datatype.TextContentRecv // recvContent is all content extracted from recvMsg and will be used later
	reply       bool                     // reply is whether to use ReplyToMessageID
	ok          bool                     // ok is whether to send this message
	respContent datatype.TextContentResp // respContent is all content will be packed into respMsg
	respMsg     tgbotapi.MessageConfig   // is the tgbotapi.MessageConfig object to be replied
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

// CommandMessageHandler is for handling normal pure text message with command
type CommandMessageHandler struct {
	recvMsg     *tgbotapi.Message
	recvContent datatype.CommandContentRecv
	reply       bool
	ok          bool
	parseMode   string
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
	h.respContent.Text, h.parseMode, h.ok, _ = static.ParseCommandMessage(h.recvContent)
}

func (h *CommandMessageHandler) Pack() {
	if !h.ok {
		return
	}
	// Start packing you msg
	h.respMsg = tgbotapi.NewMessage(h.recvMsg.Chat.ID, h.respContent.Text)

	// End packing
	h.respMsg.ParseMode = h.parseMode
	if h.reply {
		h.respMsg.ReplyToMessageID = h.recvMsg.MessageID
	}
}

func (h *CommandMessageHandler) Send(ch chan tgbotapi.MessageConfig) {
	if h.ok {
		log.Println("[Replied]", h.respMsg.ParseMode)
		ch <- h.respMsg
	}
}

package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"strings"
	"unicode"
)

func NormalTextMessage(recvMsg *tgbotapi.Message) (string, bool, error) {
	log.Printf("[TEXT] %s", recvMsg.Text)
	p := rand.Intn(100)
	if p < 15 {
		return recvMsg.Text + "个几把", true, nil
	}
	return "", false, nil
}

func NormalCommandMessage(recvMsg *tgbotapi.Message) (string, bool, error) {
	splitter := func(r rune) bool { return unicode.IsSpace(r) }
	var msg string
	switch recvMsg.Command() {
	case "help":
		msg = "按 \"/\" 自己看"
	case "choice":
		log.Printf("[CMD] %s", recvMsg.Text)
		dices := strings.FieldsFunc(recvMsg.Text, splitter)
		if len(dices) == 1 {
			msg = "你选寄吧呢"
		} else if len(dices) == 2 {
			msg = "就一个你选寄吧呢"
		} else {
			dices = dices[1:]
			msg = dices[rand.Intn(len(dices))]
		}
	case "status":
		msg = "I'm 凹K."
	default:
		msg = "你说寄吧呢"
	}
	return msg, true, nil
}

package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

func main() {
	ff := func(r rune) bool { return unicode.IsSpace(r) }

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TELEGRAM_APITOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// ignore any non-Message updates
		if update.Message == nil {
			continue
		}

		// Then if we got a message
		recvMsg := update.Message

		// ignore non-text message
		if len(recvMsg.Text) == 0 {
			continue
		}

		// whether to reply
		replyWhat := false
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// process any non-command messages
		if !recvMsg.IsCommand() {
			log.Printf("[TEXT] %s", recvMsg.Text)
			p := rand.Intn(100)
			if p < 15 {
				replyWhat = true
				msg.Text = recvMsg.Text + "个几把"
			}
		} else {
			replyWhat = true
			switch recvMsg.Command() {
			case "help":
				msg.Text = "按 \"/\" 自己看"
			case "choice":
				log.Printf("[CMD] %s", recvMsg.Text)
				dices := strings.FieldsFunc(recvMsg.Text, ff)
				if len(dices) == 1 {
					msg.Text = "你选寄吧呢"
				} else if len(dices) == 2 {
					msg.Text = "就一个你选寄吧呢"
				} else {
					dices = dices[1:]
					msg.Text = dices[rand.Intn(len(dices))]
				}
			case "status":
				msg.Text = "I'm 凹K."
			default:
				msg.Text = "你说寄吧呢"
			}
		}

		// log.Printf("[%s@%s] %s", recvMsg.From.UserName, recvMsg.Chat.ID, recvMsg.Text)

		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		if replyWhat {
			msg.ReplyToMessageID = update.Message.MessageID
			_, err := bot.Send(msg)
			if err != nil {
				log.Fatalf("error sending msg: %s", err)
			}
		}
	}
}
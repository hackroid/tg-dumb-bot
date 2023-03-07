package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type msgHandlerFunc func(recvMsg *tgbotapi.Message) (string, bool, error)

type BotServer struct {
	bot           *tgbotapi.BotAPI
	sendCh        chan tgbotapi.MessageConfig
	handlerMapper map[uint8]msgHandlerFunc
	updates       tgbotapi.UpdatesChannel
}

var (
	MSG_TYPE_TEXT uint8 = 0
	MSG_TYPE_CMD  uint8 = 1
)

func (b *BotServer) initBot() {
	log.Println("Starting tg-DUMB-bot ...")
	var err error

	// Load global var
	b.handlerMapper = make(map[uint8]msgHandlerFunc)

	// Load env var
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TELEGRAM_APITOKEN")
	debug := os.Getenv("DEBUG") == "1"

	// New bot instance
	b.bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	b.bot.Debug = debug

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	b.updates = b.bot.GetUpdatesChan(u)
	b.updates.Clear()

	b.initSendQueue()

	log.Println("Bot Initialization Complete")
	log.Println("===============================")
}

func (b *BotServer) pollingChannelUpdates() {
	for update := range b.updates {
		// Ignore any non-Message updates
		if update.Message == nil {
			continue
		}

		go b.handleChannelUpdate(update)
	}
}

func (b *BotServer) initSendQueue() {
	// This func makes a sending queue, wait for msg in channel
	// and send it one by one
	ch := make(chan tgbotapi.MessageConfig, 100)
	go func() {
		for replyMsg := range ch {
			_, err := b.bot.Send(replyMsg)
			if err != nil {
				log.Fatalf("error sending msg: %s", err)
			}
		}
	}()
	b.sendCh = ch
}

func (b *BotServer) handleChannelUpdate(update tgbotapi.Update) {
	// Then if we got a message
	recvMsg := update.Message

	// ignore non-text message
	if len(recvMsg.Text) == 0 {
		return
	}

	// Get message type
	msgType := getMessageType(recvMsg)
	respMsgText, replyWhat, _ := b.handlerMapper[msgType](recvMsg)

	// log.Printf("[%s@%s] %s", recvMsg.From.UserName, recvMsg.Chat.ID, recvMsg.Text)
	// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	if replyWhat {
		log.Printf("[REPLY] %s", respMsgText)
		replyMsg := tgbotapi.NewMessage(update.Message.Chat.ID, respMsgText)
		replyMsg.ReplyToMessageID = update.Message.MessageID
		// Put the replyMsg into sending queue
		b.sendCh <- replyMsg
	}
}

func getMessageType(recvMsg *tgbotapi.Message) uint8 {
	if recvMsg.IsCommand() {
		return MSG_TYPE_CMD
	}
	return MSG_TYPE_TEXT
}

func (b *BotServer) addMessageHandler(handleType uint8, f msgHandlerFunc) {
	b.handlerMapper[handleType] = f
}

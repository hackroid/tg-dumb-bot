package botserver

import (
	"bufio"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hackroid/tg-dumb-bot/pkg/handler"
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

func New() *BotServer {
	return &BotServer{}
}

func (b *BotServer) InitBot() {
	var err error

	// Load env var
	token := os.Getenv("TELEGRAM_APITOKEN")
	debug := os.Getenv("DEBUG") == "1"

	// New bot instance
	b.bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	b.bot.Debug = debug

	log.Printf("Authorized on account %s\n", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Load member var
	b.sendCh = b.initSendQueue()
	b.handlerMapper = make(map[uint8]msgHandlerFunc)
	b.updates = b.bot.GetUpdatesChan(u)
	b.updates.Clear()

	log.Println("Bot Initialization Complete")
	log.Println("===============================")
}

func (b *BotServer) initSendQueue() chan tgbotapi.MessageConfig {
	// This func makes a sending queue, wait for msg in channel
	// and send it one by one
	ch := make(chan tgbotapi.MessageConfig, 100)
	go func() {
		for replyMsg := range ch {
			_, err := b.bot.Send(replyMsg)
			if err != nil {
				log.Printf("[ERROR] sending msg: %s\n", err)
			}
		}
	}()
	return ch
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

func (b *BotServer) handleChannelUpdate(update tgbotapi.Update) {
	// Then if we got a message
	recvMsg := update.Message

	// ignore non-text message
	if len(recvMsg.Text) == 0 {
		return
	}

	// Get message handler
	h, _ := handler.GetMessageHandler(recvMsg)
	if h == nil {
		return
	}

	// Handler execution
	h.Extract()
	h.Generate()
	h.Pack()
	h.Send(b.sendCh)
}

func (b *BotServer) Serve() {
	go b.pollingChannelUpdates()
	go b.cmd()
}

func (b *BotServer) cmd() {
	var s string
	reader := bufio.NewReader(os.Stdin)

	for {
		// fmt.Print("> ")
		s, _ = reader.ReadString('\n')
		fmt.Printf("Hello, %s", s)
	}
}

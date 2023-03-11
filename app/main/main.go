package main

import (
	"github.com/hackroid/tg-dumb-bot/pkg/botserver"
	"github.com/hackroid/tg-dumb-bot/pkg/constants"
	"github.com/hackroid/tg-dumb-bot/pkg/handler"
	"github.com/joho/godotenv"
	"log"
	"runtime"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	log.Printf("Start tg-DUMB-bot version #TODO; Go %s (%s/%s)\n",
		runtime.Version(), runtime.GOOS, runtime.GOARCH)

	bs := botserver.New()
	bs.InitBot()
	bs.AddMessageHandler(constants.MsgTypeCmd, handler.NormalCommandMessage)
	bs.AddMessageHandler(constants.MsgTypeText, handler.NormalTextMessage)
	bs.Serve()
}

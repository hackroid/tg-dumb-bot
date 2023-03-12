package main

import (
	"github.com/hackroid/tg-dumb-bot/pkg/botserver"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	log.Printf("Start tg-DUMB-bot version #TODO; Go %s (%s/%s)\n",
		runtime.Version(), runtime.GOOS, runtime.GOARCH)
	bs := botserver.New()
	bs.InitBot()
	bs.Serve()

	<-gracefulShutdown
	log.Println("Dumb Bot Stopped, Bye!")
}

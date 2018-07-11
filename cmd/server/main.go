package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/neoplatonist/botManager/cmd/server/bot"
	"github.com/neoplatonist/botManager/cmd/server/cli"
)

var (
	messages = make(chan string)
)

func config() {
	fmt.Println("Starting Bot Manager...")
	fmt.Println("------------------------------")

	// start our bot
	bot.Start()

	fmt.Println("------------------------------")
	fmt.Println("Bot Manager has been started. Press CTRL-C to exit.")
	messages <- "started"
}

func main() {
	// Config
	go config()

	<-messages
	// CLI
	go cli.Listen()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("\nWe are sad to see you leave. Until next time!")
}

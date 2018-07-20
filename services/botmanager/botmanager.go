package botmanager

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/neoplatonist/botManager/services/bot"
	"github.com/neoplatonist/botManager/services/cli/server"
)

// NOTES: Probably need to handle Registration and Connecting to modules differently

func config(messages chan<- string) {
	fmt.Println("Starting Bot Manager...")
	fmt.Println("------------------------------")

	// start our bot
	bot.Start()

	fmt.Println("------------------------------")
	fmt.Println("Bot Manager has been started. Press CTRL-C to exit.")
	messages <- "started"
}

// Start initializes the botManager and grpc cli listener
func Start() {
	var messages = make(chan string)
	// Config
	go config(messages)

	<-messages
	// CLI
	go cli.Listen()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("\nWe are sad to see you leave. Until next time!")
}

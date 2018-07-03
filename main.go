package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/neoplatonist/discord-bot/bot-v2/bot"
)

var (
	file     = os.Stdin
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

func cmdListener(input *os.File) []byte {
	// Listens to input from console
	buf := bufio.NewReader(input)
	fmt.Print("> ")

	// Grabs command on submit
	text, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	}

	return bytes.TrimSpace(text)
}

func cli() {
	<-messages

	for {
		input := cmdListener(file)

		if string(input) == "" {
			continue
		}

		fmt.Printf("You typed command: %s\n", input)

		// Meta Commands

		// Prepare Commands
		// ? potentially already done, maybe try to hook into them

		// Execute Commands
		// for _, cmd := range command.ActiveCommands {
		// 	if string(input) == cmd.Name {
		// 		cmd.Action()
		// 	}
		// }
	}
}

func main() {
	// Config
	go config()

	// CLI
	go cli()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("\nWe are sad to see you leave. Until next time!")
}

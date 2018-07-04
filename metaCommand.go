package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/neoplatonist/botManager/bot"
	command "github.com/neoplatonist/botManager/commands"
)

const (
	metaSuccess      = "Meta Success."
	metaUnrecognized = "Meta Unrecognized."
)

type metaCommands interface {
	Execute()
}

meta[command].Execute(input)

func doMetaCommand(input []byte) string {
	inputArr := strings.Fields(string(input))

	switch inputArr[0] {
	case ".exit":
		fmt.Println("We are sad to see you leave. Until next time!")
		os.Exit(3)

	case ".dc", ".disconnect":
		if len(inputArr) < 2 {
			fmt.Println("no module specified")
			return metaSuccess
		}

		bot.Disconnect(inputArr[1])
		return metaSuccess

	case ".c", ".connect":
		if len(inputArr) < 2 {
			fmt.Println("no module specified")
			return metaSuccess
		}

		bot.Connect(inputArr[1])
		return metaSuccess

	case ".help":
		if len(inputArr) < 2 {
			fmt.Println("no module specified")
			return metaSuccess
		}

		list := command.List(inputArr[1])
		if len(list) < 1 {
			fmt.Println("no commands found for module:", inputArr[1])
		}

		fmt.Println(list)

		return metaSuccess
	}

	return metaUnrecognized
}

func checkMeta(command []byte) string {
	if string(command[0]) != "." {
		return ""
	}

	if err := doMetaCommand(command); err == metaUnrecognized {
		fmt.Printf("unrecognized Meta Command %q. \n", command)
		return metaUnrecognized
	}

	return metaSuccess
}

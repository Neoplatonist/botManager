package cli

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/neoplatonist/botManager/cmd/server/bot"
	command "github.com/neoplatonist/botManager/cmd/server/commands"
)

const (
	metaSuccess      = "Meta Success."
	metaUnrecognized = "Meta Unrecognized."
)

func doMetaCommand(input []byte) (string, error) {
	inputArr := strings.Fields(string(input))

	switch inputArr[0] {
	case ".exit":
		// kills server after 5 seconds
		exit := func() {
			time.Sleep(5 * time.Second)
			os.Exit(3)
		}

		go exit()

		return "We are sad to see you leave. Until next time!", nil

	case ".dc", ".disconnect":
		if len(inputArr) < 2 {
			return "", errors.New("no module specified")
		}

		resp, err := bot.Disconnect(inputArr[1])
		if err != nil {
			return "", err
		}

		return resp, nil

	case ".c", ".connect":
		if len(inputArr) < 2 {
			return "no module specified", nil
		}

		resp, err := bot.Connect(inputArr[1])
		if err != nil {
			return "", err
		}

		return resp, nil

	case ".help":
		if len(inputArr) < 2 {
			return "no module specified", nil
		}

		list := command.List(inputArr[1])
		if len(list) < 1 {
			return "", errors.New("no commands found for module: " + inputArr[1])
		}

		return list, nil
	}

	return "", errors.New("no command found: " + string(input))
}

func checkMeta(command []byte) (string, error) {
	if string(command[0]) != "." {
		return "", errors.New("not a command: " + string(command))
	}

	resp, err := doMetaCommand(command)
	if err != nil {
		return "", err
	}

	return resp, nil
}

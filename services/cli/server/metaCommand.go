package cli

import (
	"os"
	"strings"
	"time"

	"github.com/neoplatonist/botManager/services/bot"
	"github.com/neoplatonist/botManager/services/modules"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	metaSuccess      = "Meta Success."
	metaUnrecognized = "Meta Unrecognized."
)

func stringComb(list []string) string {
	return strings.Join(list, "\n")
}

func doMetaCommand(input []byte) (string, error) {
	inputArr := strings.Fields(string(input))
	cmd := []string{
		".exit - closes the server",
		".dc / .disconnect <module> - module connection closed",
		".c / .connect <module> - module connection started",
		".modules - lists all active modules",
		".help - lists all meta commands",
		"<module> -help - lists help for a particular module",
	}

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
			return "", status.Error(codes.InvalidArgument, "no module specified")
		}

		resp, err := bot.Disconnect(inputArr[1])
		if err != nil {
			return "", err
		}

		return resp, nil

	case ".c", ".connect":
		if len(inputArr) < 2 {
			return "", status.Error(codes.InvalidArgument, "no module specified")
		}

		resp, err := bot.Connect(inputArr[1])
		if err != nil {
			return "", err
		}

		return resp, nil

	case ".modules":
		return modules.RegisteredModules(), nil

	case ".help":
		if len(inputArr) < 2 {
			return stringComb(cmd), nil
		}
	}

	return "", status.Error(codes.InvalidArgument, "no command found: "+string(input))
}

func checkMeta(command []byte) (string, error) {
	if string(command[0]) != "." {
		return "", status.Error(codes.InvalidArgument, "not a command: "+string(command))
	}

	resp, err := doMetaCommand(command)
	if err != nil {
		return "", err
	}

	return resp, nil
}

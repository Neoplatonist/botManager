package cli

import (
	"strings"

	"github.com/neoplatonist/botManager/services/bot"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func checkCmd(input []byte) (string, error) {
	inputArr := strings.Fields(string(input))

	mod, ok := bot.ModuleList[inputArr[0]]
	if ok {
		resp, err := mod.Command(inputArr)
		if err == nil {
			return resp, nil
		}
	}

	return "", status.Error(codes.InvalidArgument, "no command found")
}

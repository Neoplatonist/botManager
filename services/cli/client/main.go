package client

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"

	pb "github.com/neoplatonist/botManager/services/cli/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const port = ":6000"

var (
	file = os.Stdin
)

type env struct {
	pb.CLIClient
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

func (e env) cli() {
	for {
		input := cmdListener(file)

		if string(input) == "" {
			continue
		}

		resp, err := e.CliCommand(
			context.Background(),
			&pb.CliCommandReq{Command: input},
		)
		if err != nil {
			stat := status.Convert(err)
			fmt.Println(stat.Message())
			continue
		}

		fmt.Println(resp.Message)
	}
}

// Start the grpc client connection to the server
// then runs the repl watcher
func Start() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}

	e := env{
		pb.NewCLIClient(conn),
	}

	e.cli()
}

package cli

import (
	"context"
	"log"
	"net"

	pb "github.com/neoplatonist/botManager/proto"
	"google.golang.org/grpc"
)

const port = ":6000"

type cliService struct {
	CLI Service
}

type Service interface {
	CliCommand(command []byte) error
}

// CliCommand checks the commands and executes them
func (c *cliService) CliCommand(ctx context.Context, in *pb.CliCommandReq) (*pb.CliCommandRes, error) {
	// Handles meta (.commands) commands
	resp, err := checkMeta(in.Command)
	if err != nil {
		return &pb.CliCommandRes{
			Message: "",
		}, err
	}

	return &pb.CliCommandRes{
		Message: resp,
	}, nil
}

// Listen starts the grpc server listener
func Listen() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	service := cliService{}
	pb.RegisterCLIServer(grpcServer, &service)
	grpcServer.Serve(lis)
}

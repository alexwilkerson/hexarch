package rpc

import (
	"log"
	"net"

	"github.com/alexwilkerson/hexarch/internal/adapters/left/grpc/pb"
	"github.com/alexwilkerson/hexarch/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (grpca Adapter) Run() {
	var err error

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to server gRPC server over port 9000: %v", err)
	}
}

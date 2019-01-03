package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/maddevsio/grpc-rest-api-example/pb"
	"google.golang.org/grpc"
)

// Greeter ...
type Greeter struct {
}

// New creates new server greeter
func New() *Greeter {
	return &Greeter{}
}

// Start starts server
func (g *Greeter) Start() error {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, g)
	grpcServer.Serve(lis)
	return nil
}

// SayHello says hello
func (g *Greeter) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello, %s!", r.Name),
	}, nil
}

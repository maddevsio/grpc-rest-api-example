package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/maddevsio/grpc-rest-api-example/pb"
	"google.golang.org/grpc"
)

// Greeter ...
type Greeter struct {
	wg sync.WaitGroup
}

// New creates new server greeter
func New() *Greeter {
	return &Greeter{}
}

// Start starts server
func (g *Greeter) Start() {
	g.wg.Add(1)
	go func() {
		log.Fatal(g.startGRPC())
		g.wg.Done()
	}()
	g.wg.Add(1)
	go func() {
		log.Fatal(g.startREST())
		g.wg.Done()
	}()
}
func (g *Greeter) WaitStop() {
	g.wg.Wait()
}
func (g *Greeter) startGRPC() error {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, g)
	grpcServer.Serve(lis)
	return nil
}
func (g *Greeter) startREST() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, ":8080", opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8090", mux)
}

// SayHello says hello
func (g *Greeter) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello, %s!", r.Name),
	}, nil
}

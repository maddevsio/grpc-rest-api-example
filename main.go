package main

import "github.com/maddevsio/grpc-rest-api-example/server"

func main() {
	g := server.New()
	g.Start()
	g.WaitStop()
}

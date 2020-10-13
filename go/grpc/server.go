package hello

import (
	"context"
	"log"
)

const (
	name     = "server"
	greeting = "World, Hello!"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	log.Printf("From %s: %s", in.GetName(), in.GetMessage())
	return &HelloResponse{Name: name, Message: greeting}, nil
}

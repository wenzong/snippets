package hello

import (
	"context"
	"log"
)

const (
	name     = "server"
	greeting = "World, Hello!"
	N        = 5
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	log.Printf("From %s: %s", in.GetName(), in.GetMessage())
	return &HelloResponse{Name: name, Message: greeting}, nil
}

func (s *server) SayNTimesHello(in *HelloRequest, stream HelloService_SayNTimesHelloServer) error {
	log.Printf("From %s: %s", in.GetName(), in.GetMessage())
	for i := 0; i < N; i++ {
		if err := stream.Send(&HelloResponse{Name: name, Message: greeting}); err != nil {
			return err
		}
	}

	return nil
}

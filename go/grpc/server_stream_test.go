package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type mockHelloService_SayNTimesHelloServer struct {
	grpc.ServerStream
	Results []*HelloResponse
}

func (m *mockHelloService_SayNTimesHelloServer) Send(out *HelloResponse) error {
	m.Results = append(m.Results, out)
	return nil
}

func TestSayNTimesHello(t *testing.T) {
	s := &server{}
	req := &HelloRequest{Name: "client", Message: "Hello, World!"}

	mock := &mockHelloService_SayNTimesHelloServer{}
	s.SayNTimesHello(req, mock)

	assert.Equal(t, N, len(mock.Results), "Result expected to contain %s item", N)
}

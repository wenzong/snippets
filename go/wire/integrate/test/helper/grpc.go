package helper

import (
	"context"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const (
	bufSize = 1024 * 1024
)

var (
	listen *bufconn.Listener
	once   sync.Once
)

func NewTestListener() *bufconn.Listener {
	once.Do(func() {
		listen = bufconn.Listen(bufSize)
	})

	return listen
}

func NewTestServerOption() []grpc.ServerOption {
	return []grpc.ServerOption{}
}

func NewTestClientConn(l *bufconn.Listener) *grpc.ClientConn {
	conn, _ := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return l.Dial() }),
		grpc.WithInsecure())

	return conn
}

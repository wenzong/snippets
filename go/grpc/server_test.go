package hello

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var (
	lis    *bufconn.Listener
	client HelloServiceClient
	svr    HelloServiceServer
)

func startServer() *grpc.Server {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	svr = &server{}
	RegisterHelloServiceServer(s, svr)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	return s
}

func startClient() *grpc.ClientConn {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return lis.Dial() }),
		grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
	}
	client = NewHelloServiceClient(conn)

	return conn
}

func TestMain(m *testing.M) {
	s := startServer()
	defer s.GracefulStop()

	c := startClient()
	defer c.Close()

	os.Exit(m.Run())
}

func TestSayHelloAsync(t *testing.T) {
	resp, err := client.SayHello(context.Background(), &HelloRequest{Name: "client", Message: "Hello, World!"})
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}

	if resp.Message != greeting {
		t.Errorf("SayHello failed: %s should be %s", resp.Message, greeting)
	}
}

func TestSayHelloSync(t *testing.T) {
	resp, err := svr.SayHello(context.Background(), &HelloRequest{Name: "client", Message: "Hello, World!"})
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}

	if resp.Message != greeting {
		t.Errorf("SayHello failed: %s should be %s", resp.Message, greeting)
	}
}

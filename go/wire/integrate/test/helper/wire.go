// +build wireinject

package helper

import (
	"net"
	"testing"

	"github.com/google/wire"
	"github.com/wenzong/demo/api/pb"
	"github.com/wenzong/demo/biz/user"
	"github.com/wenzong/demo/infra/app"
	"github.com/wenzong/demo/infra/db"
	infragrpc "github.com/wenzong/demo/infra/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func gRPCRegisterServiceFn(
	userServer *user.Server,
) infragrpc.RegisterServiceFunc {
	return func(s *grpc.Server) *grpc.Server {
		pb.RegisterUserServiceServer(s, userServer)
		return s
	}
}

var ProviderSet = wire.NewSet(
	NewTestConfig,
	infragrpc.NewServer,
	NewTestListener,
	NewTestServerOption,
	NewTestClientConn,
	NewNilHTTPServer,
	wire.Bind(new(net.Listener), new(*bufconn.Listener)),
	gRPCRegisterServiceFn,
)

var Set = wire.NewSet(
	// config.ProviderSet,
	app.ProviderSet,
	db.ProviderSet,
	user.ProviderSet,
	// http.ProviderSet,
	// grpc.ProviderSet,
	ProviderSet,
)

func App(t *testing.T) (*app.App, func()) {
	panic(wire.Build(Set))
}

func ClientConn() *grpc.ClientConn {
	panic(wire.Build(Set))
}

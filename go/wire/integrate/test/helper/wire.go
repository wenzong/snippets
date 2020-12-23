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
	"github.com/wenzong/demo/infra/log"
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
	app.ProviderSet,
	// config.ProviderSet,
	db.ProviderSet,
	log.ProviderSet,
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

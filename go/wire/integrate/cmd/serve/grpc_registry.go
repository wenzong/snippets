package serve

import (
	"github.com/wenzong/demo/api/pb"
	"github.com/wenzong/demo/biz/user"
	infragrpc "github.com/wenzong/demo/infra/grpc"
	"google.golang.org/grpc"
)

// 所有 gRPC Service 统一在此注册
func gRPCRegisterServiceFn(
	userServer *user.Server,
) infragrpc.RegisterServiceFunc {
	return func(s *grpc.Server) *grpc.Server {
		pb.RegisterUserServiceServer(s, userServer)
		return s
	}
}

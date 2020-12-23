package user

import (
	"context"

	"github.com/wenzong/demo/api/pb"
	"github.com/wenzong/demo/infra/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	service *Service
	logger  func(context.Context) log.Logger
}

// 获取用户信息
func (s *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	s.logger(ctx).Debug("获取用户信息")

	u, err := s.service.Get(ctx, in.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "获取用户信息失败")
	}

	return &pb.GetResponse{User: u}, nil
}

func NewServer(s *Service, fn func(context.Context) log.Logger) *Server {
	return &Server{service: s, logger: fn}
}

// make sure Server is pb.UserServiceServer
var _ pb.UserServiceServer = (*Server)(nil)

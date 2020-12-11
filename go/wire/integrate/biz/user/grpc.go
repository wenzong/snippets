package user

import (
	"context"

	"github.com/wenzong/demo/api/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	service *Service
}

// 获取用户信息
func (s *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	u, err := s.service.Get(ctx, in.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "获取用户信息失败")
	}

	return &pb.GetResponse{User: u}, nil
}

func NewServer(s *Service) *Server {
	return &Server{service: s}
}

// make sure Server is pb.UserServiceServer
var _ pb.UserServiceServer = (*Server)(nil)

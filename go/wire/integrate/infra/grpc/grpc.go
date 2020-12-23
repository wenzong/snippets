package grpc

import (
	"google.golang.org/grpc"
)

type RegisterServiceFunc func(s *grpc.Server) *grpc.Server

// cmd 需要提供:
// + []grpc.ServerOption: e.g. Interceptors
// + RegisterServiceFunc: e.g. proto.RegisterXXXServer(s, xxxServer)
func NewServer(options []grpc.ServerOption, fn RegisterServiceFunc) *grpc.Server {
	return fn(grpc.NewServer(options...))
}

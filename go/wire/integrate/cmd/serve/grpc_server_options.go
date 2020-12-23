package serve

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"

	// grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	grpc_ctxtags "github.com/wenzong/demo/infra/grpc/interceptors/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	traceHeaderName = "uber-trace-id"
)

// 可选配置
// + ctxtag
// + opentracing
// + prometheus
// + zap
// + recovery
// + auth
// + ratelimit
//
// 有些 interceptor 需要依赖其他组件，在 infra/grpc/interceptors 中提供
func gRPCServerOptions(logger *zap.Logger) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.ExtractUserId,
			),
			// grpc_opentracing.UnaryServerInterceptor(
			// 	grpc_opentracing.WithTraceHeaderName(traceHeaderName),
			// ),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	}
}

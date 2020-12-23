/* grpc_ctxtags

Reuse `github.com/grpc-ecosystem/go-grpc-middleware/tags`'s context **Tags**

Use user provided `TagExtractFunc`, which takes ctx, req methodName for later usage
+ metadata(from context)
+ request
+ gRPC ServerInfo(only UnaryServerInfo for now)
*/
package grpc_ctxtags

import (
	"context"

	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
)

type TagExtractFunc func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo) map[string]interface{}

// UnaryServerInterceptor returns a new unary server interceptors that sets the values for request tags.
func UnaryServerInterceptor(fns ...TagExtractFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx := grpc_ctxtags.SetInContext(ctx, grpc_ctxtags.NewTags())

		t := grpc_ctxtags.Extract(newCtx)
		for _, fn := range fns {
			if valMap := fn(ctx, req, info); valMap != nil {
				for k, v := range valMap {
					t.Set(k, v)
				}
			}
		}

		return handler(newCtx, req)
	}
}

func ExtractUserId(_ context.Context, req interface{}, _ *grpc.UnaryServerInfo) map[string]interface{} {
	if msg, ok := req.(interface {
		GetUserId() int64
	}); ok {
		return map[string]interface{}{
			"grpc.request.user_id": msg.GetUserId(),
		}
	}

	return nil
}

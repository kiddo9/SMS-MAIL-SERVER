package middleware

import (
	"context"

	"google.golang.org/grpc"
)


func Interceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.ServerOption{
	return grpc.ChainUnaryInterceptor(
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			// recursively apply all interceptors
			var h grpc.UnaryHandler = handler
			for i := len(interceptors) - 1; i >= 0; i-- {
				interceptor := interceptors[i]
				next := h
				h = func(ctx context.Context, req interface{}) (interface{}, error) {
					return interceptor(ctx, req, info, next)
				}
			}
			return h(ctx, req)
		},
	)
}
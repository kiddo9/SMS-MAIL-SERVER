package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func RecaptchaMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)(interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "context not found")
	}

	recaptchaToken := md.Get("recaptcha-token")
	if len(recaptchaToken) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid arguments")
	}

	return handler(ctx, req)
}
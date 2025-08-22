package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func RecaptchaMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)(interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "server responsed with a 404 error")
	}

	recaptchaToken := md.Get("recaptcha-token")
	fmt.Println(recaptchaToken, recaptchaToken[0])
	
	if len(recaptchaToken) == 0 || recaptchaToken[0] == "" {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid arguments")
	}

	recaptchaValidateUrl := os.Getenv("RECAPTCHA_VALIDATE_URL")
	secret := os.Getenv("RECAPTCHA_SECRET_KEY")

	requestString := fmt.Sprintf("%s?secret=%s&response=%s", recaptchaValidateUrl, secret, recaptchaToken[0])
	
	resp, err := http.Post(requestString, "application/json", nil)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Server rejected your request")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "server responsed with a 400 error")
	}

	if data["success"] != true {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized request",)
	}

	return handler(ctx, req)
}
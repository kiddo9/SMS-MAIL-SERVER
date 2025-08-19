package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kiddo9/SMS-MAIL-SERVER/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	skipAuth := map[string]bool{
		"/message.proto.AdminService/LoginAdmin":    true,
		"/message.proto.AdminService/ValidateToken": true,
	}

	if skipAuth[info.FullMethod] {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata not found in context")
	}

	authToken := md.Get("auth-token")
	if len(authToken) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "auth token not provided")
	}
	if len(authToken) > 1 {
		return nil, status.Errorf(codes.Unauthenticated, "multiple auth tokens provided")
	}

	decerptedToken, err := utils.ValidateToken(authToken[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	details, ok := decerptedToken.Claims.(jwt.MapClaims)
	if !ok || !decerptedToken.Valid {
		return nil, status.Errorf(codes.Unauthenticated, "invalid request")
	}

	fileName := "storage/admin.json"

	body, err := os.Open(fileName)
	fmt.Print(body)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not open file: %v", err)
	}

	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not read file: %v", err)
	}

	var admins []map[string]interface{}
	err = json.Unmarshal(data, &admins)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not read file: %v", err)
	}

	for _, admin := range admins {
		if details["uuid"] == admin["uuid"] {
			if admin["ApiKey"] == details["APIKey"] {
				jwtExpired := details["exp"].(float64)

				if time.Now().Unix() > int64(jwtExpired) {
					GetAdminLongTermToken := admin["Jwt"].(string)

					validateToken, err := utils.ValidateToken(GetAdminLongTermToken)
					if err != nil {
						return nil, status.Errorf(codes.Unauthenticated, "invalid long term token: %v", err)
					}

					infoData, ok := validateToken.Claims.(jwt.MapClaims)
					if !ok || !validateToken.Valid {
						return nil, status.Errorf(codes.Unauthenticated, "invalid long term token")
					}

					if time.Now().Unix() < int64(infoData["exp"].(float64)) {
						details["uuid"] = infoData["uuid"]
						details["APIKey"] = infoData["APIKey"]

						newToken, err := utils.GenerateRequestJWTToken(
							infoData["uuid"].(string),
							infoData["APIKey"].(string),
						)

						if err != nil {
							return nil, status.Errorf(codes.Internal, "could not generate new token: %v", err)
						}

						md.Set("auth-token", newToken)
						ctx = metadata.NewIncomingContext(ctx, md)
					} else {
						return nil, status.Errorf(codes.Unauthenticated, "long term token has expired")
					}
				}
			}
			break
		}
	}

	return handler(ctx, req)

}

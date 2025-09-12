package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kiddo9/SMS-MAIL-SERVER/structures"
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
		"/message.proto.AdminService/SendOtp": true,
		"/message.proto.AdminService/VerifyOtp": true,
	}

	if _, ok := skipAuth[info.FullMethod]; ok {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorizied request. request denied")
	}

	authToken := md.Get("x-auth-token")
	if len(authToken) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "requested denied")
	}
	if len(authToken) > 1 {
		return nil, status.Errorf(codes.Unauthenticated, "invalid request. request terminated")
	}

	decerptedToken, err := utils.ValidateToken(authToken[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid aurgument")
	}

	details, ok := decerptedToken.Claims.(jwt.MapClaims)
	if !ok || !decerptedToken.Valid {
		return nil, status.Errorf(codes.Unauthenticated, "invalid request")
	}

	fileName := "storage/admin.json"

	_, err = os.Open(fileName)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	var admins []structures.AdminStructs
	err = json.Unmarshal(data, &admins)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	for _, admin := range admins {
		if details["uuid"] == admin.Uuid {
			if admin.APIKey == details["APIKey"] {
				jwtExpired := details["exp"].(float64)

				if time.Now().Unix() > int64(jwtExpired) {
					GetAdminLongTermToken := admin.Jwt

					validateToken, err := utils.ValidateToken(GetAdminLongTermToken)
					if err != nil {
						return nil, status.Errorf(codes.Unauthenticated, "unkown user")
					}

					infoData, ok := validateToken.Claims.(jwt.MapClaims)
					if !ok || !validateToken.Valid {
						return nil, status.Errorf(codes.Unauthenticated, "unkown")
					}

					if time.Now().Unix() < int64(infoData["exp"].(float64)) {
						details["uuid"] = infoData["uuid"]
						details["APIKey"] = infoData["APIKey"]

						newToken, err := utils.GenerateRequestJWTToken(
							infoData["uuid"].(string),
							infoData["APIKey"].(string),
						)

						if err != nil {
							return nil, status.Errorf(codes.Internal, "internal server error")
						}

						md.Set("auth-token", newToken)
						ctx = metadata.NewIncomingContext(ctx, md)
						fmt.Println(md, ctx)
					} else {
						return nil, status.Errorf(codes.Unauthenticated, "expired")
					}
				}
			}
			break
		}
	}

	return handler(ctx, req)

}

package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/kiddo9/SMS-MAIL-SERVER/message/proto"
	"github.com/kiddo9/SMS-MAIL-SERVER/structures"
	"github.com/kiddo9/SMS-MAIL-SERVER/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AdminHandler struct {
	pb.UnimplementedAdminServiceServer
}

func (h *AdminHandler) LoginAdmin(ctx context.Context, req *pb.OtpRequest)(*pb.OtpResponse, error){
	_, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		log.Fatalf("error occoured %v", ok)
		return nil, status.Errorf(codes.Unauthenticated, "context absent")
	}

	fileName := "storage/admin.json"

	_, err := os.Open(fileName)


	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not open file: %v", err)
	}

	date, err := os.ReadFile(fileName)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not read file: %v", err)
	}

	var admins []structures.AdminStructs

	err = json.Unmarshal(date, &admins)

	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "could not unmarshal data: %v", err)
	}

	email := req.GetEmail()

	if email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "system got an empty data")
	}

	var jwtToken string

	for _, emails := range admins {
		if email != emails.Email {
			return nil, status.Errorf(codes.NotFound, "request returned a 404 response")
		}
		tokenExpiry := time.Now().Add(time.Minute * 5).Unix()
		// Generate JWT token
		jwtToken, err = utils.GenerateJWTToken(emails.Email, emails.Uuid, emails.APIKey, tokenExpiry)
		if err != nil {
			log.Fatalf("error generating JWT token: %v", err)
			return nil, status.Errorf(codes.Internal, "could not generate JWT token")
		}
		token := utils.GenerateCode(8)
		
		
		fmt.Printf("Token sent %v. token %v", token, jwtToken)
	}

	
	return &pb.OtpResponse{
		Message: jwtToken,
		OtpSent: true,
	}, nil
}

// func (h *AdminHandler) VerifyOtp(ctx context.Context, req *pb.OtpVerificationRequest) (*pb.OtpVerificationResponse, error) {
	
// }
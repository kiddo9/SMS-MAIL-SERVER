package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

var data []byte
var fileName string = "storage/admin.json"

func LoadFile() error {
	_, err := os.Open(fileName)

	if err != nil {
		return status.Errorf(codes.Internal, "could not open file: %v", err)
	}

	data, err = os.ReadFile(fileName)

	if err != nil {
		return status.Errorf(codes.Internal, "could not read file: %v", err)
	}

	return nil
}

// login admin handler
func (h *AdminHandler) LoginAdmin(ctx context.Context, req *pb.OtpRequest) (*pb.OtpResponse, error) {
	_, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		log.Fatalf("error occoured %v", ok)
		return nil, status.Errorf(codes.Unauthenticated, "context absent")
	}

	LoadFile()

	var admins []structures.AdminStructs

	err := json.Unmarshal(data, &admins)

	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "could not unmarshal data: %v", err)
	}

	email := req.GetEmail()

	if email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "system got an empty data")
	}

	var jwtToken string

	for idx, emails := range admins {
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

		emails.OTP = token
		emails.OTPExpiry = fmt.Sprintf("%v", tokenExpiry)

		admins[idx] = emails

		updateDate, err := json.MarshalIndent(admins, "", "")

		if err != nil {
			return nil, status.Errorf(codes.Internal, "unable to convert back to json")
		}

		if err := os.WriteFile(fileName, updateDate, 0644); err != nil {
			return nil, status.Errorf(codes.Internal, "unable to write into file")
		}
		//logic to send to email
	}

	return &pb.OtpResponse{
		Message: jwtToken,
		OtpSent: true,
	}, nil
}

// handler to resend otp
func (h *AdminHandler) SendOtp(ctx context.Context, req *pb.OtpRequest) (*pb.OtpResponse, error) {
	_, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.Aborted, "server returned a 422 status response")
	}

	LoadFile()

	var admins []map[string]interface{}

	err := json.Unmarshal(data, &admins)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not unmarshal data: %v", err)
	}

	email := req.GetEmail()

	for idx, admin := range admins {
		if email != admin["email"] {
			return nil, status.Errorf(codes.Canceled, "request was cancelled with status code 422")
		}
		tokenExpiry := time.Now().Add(time.Minute * 5).Unix()
		token := utils.GenerateCode(8)

		admin["otp"] = token
		admin["otpExpiry"] = fmt.Sprintf("%v", tokenExpiry)

		admins[idx] = admin

		update, err := json.MarshalIndent(admins, "", "")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "unable to complete request write")
		}

		err = os.WriteFile(fileName, update, 0644)

		if err != nil {
			return nil, status.Errorf(codes.Internal, "error occured while processing your request")
		}

		//logic to send to email
	}

	return &pb.OtpResponse{
		Message: "otp sent",
		OtpSent: true,
	}, nil
}

// token validation handler
func (h *AdminHandler) ValidateToken(ctx context.Context, req *pb.TokenValidationRequest) (*pb.TokenValidationResponse, error) {
	_, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.Aborted, "you request was aborted")
	}

	LoadFile()

	tokenToBeVerified := req.GetToken()

	if tokenToBeVerified == "" {
		return nil, status.Errorf(codes.NotFound, "request returned a 404 response")
	}

	tokenResponse, err := utils.ValidateToken(tokenToBeVerified)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	claims, ok := tokenResponse.Claims.(jwt.MapClaims)

	if !ok || !tokenResponse.Valid {
		return nil, status.Errorf(codes.Unauthenticated, "token is invalid: %v", err)
	}

	if exp, ok := claims["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
		return nil, status.Errorf(codes.Unauthenticated, "token has expired")
	}

	jwtemail, emailExists := claims["email"].(string)
	Uuid, uuidExists := claims["uuid"].(string)
	apiKey, exists := claims["APIKey"].(string)

	if !emailExists || !uuidExists || !exists {
		return nil, status.Errorf(codes.Unauthenticated, "token was dismissed: %v", err)
	}

	var admins []structures.AdminStructs

	err = json.Unmarshal(data, &admins)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not unmarshal data: %v", err)
	}

	for _, admin := range admins {
		if admin.Uuid != Uuid && admin.APIKey != apiKey && admin.Email != jwtemail {
			return nil, status.Errorf(codes.NotFound, "server returned a 404 response")
		}
	}

	return &pb.TokenValidationResponse{
		IsValid: true,
		Email:   jwtemail,
	}, nil
}

func (h *AdminHandler) VerifyOtp(ctx context.Context, req *pb.OtpVerificationRequest) (*pb.OtpVerificationResponse, error) {
	LoadFile()

	email := req.GetEmail()
	otp := req.GetOtp()

	if email == "" || otp == "" {
		return nil, status.Errorf(codes.InvalidArgument, "server returned a 400 response")
	}

	var admins []structures.AdminStructs

	err := json.Unmarshal(data, &admins)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not unmarshal data: %v", err)
	}

	var loginRequestToken string

	for _, admin := range admins {
		if email != admin.Email {
			return nil, status.Errorf(codes.NotFound, "server returned a 404 response")
		}

		user := structures.AdminStructs{
			Email:     admin.Email,
			Uuid:      admin.Uuid,
			APIKey:    admin.APIKey,
			OTP:       admin.OTP,
			OTPExpiry: admin.OTPExpiry,
			Jwt:       admin.Jwt,
		}

		if otp != user.OTP {
			return nil, status.Errorf(codes.PermissionDenied, "invalid otp provided")
		}

		otpExpiry, err := strconv.Atoi(user.OTPExpiry)

		if time.Now().Unix() > int64(otpExpiry) {
			return nil, status.Errorf(codes.PermissionDenied, "otp has expired")
		}

		validateLongTermToken, err := utils.ValidateToken(user.Jwt)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid long term token: %v", err)
		}

		infoData, ok := validateLongTermToken.Claims.(jwt.MapClaims)
		if !ok || !validateLongTermToken.Valid {
			return nil, status.Errorf(codes.Unauthenticated, "invalid long term token")
		}
		if time.Now().Unix() > int64(infoData["exp"].(float64)) {
			// Generate a new long-term token if the existing one has expired
			newLongTermToken, err := utils.GenerateJWTTokenLongTerm(
				infoData["email"].(string),
				infoData["uuid"].(string),
				infoData["APIKey"].(string),
			)

			// if err != nil {
			// 	return nil, status.Errorf(codes.Internal, "could not generate new long term token: %v", err)
			// }

			user.Jwt = newLongTermToken

			// Update the admin data with the new long-term token
			admin.Jwt = newLongTermToken

			// Save the updated admin data back to the file
			updateData, err := json.MarshalIndent(admins, "", "")
			if err != nil {
				return nil, status.Errorf(codes.Internal, "could not marshal updated admin data: %v", err)
			}

			if err := os.WriteFile(fileName, updateData, 0644); err != nil {
				return nil, status.Errorf(codes.Internal, "could not write updated admin data to file: %v", err)
			}
		}
		// generate login request token
		loginRequestToken, err = utils.GenerateRequestJWTToken(user.Uuid, user.APIKey)

		if err != nil {
			return nil, status.Errorf(codes.Internal, "could not generate login request token: %v", err)
		}
	}

	fmt.Print(loginRequestToken)

	return &pb.OtpVerificationResponse{
		IsVerified: true,
		Message:    loginRequestToken,
	}, nil
}

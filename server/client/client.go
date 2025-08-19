package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kiddo9/SMS-MAIL-SERVER/message/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	client := pb.NewAdminServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := client.VerifyOtp(ctx, &pb.OtpVerificationRequest{
		Otp:   "GNUAEJzP",
		Email: "dkido913@gmail.com",
	})

	if err != nil {
		log.Fatalf("error occoured %v", err)
	}
	log.Println("Response:", resp)
}

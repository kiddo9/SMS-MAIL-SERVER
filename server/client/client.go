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

	resp, err := client.ValidateToken(ctx, &pb.TokenValidationRequest{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBUElLZXkiOiIrS2VySSgtW14_dDUhb1YoZygwL1JPYnl3X1EvR0kiLCJlbWFpbCI6ImRraWRvOTEzQGdtYWlsLmNvbSIsImV4cCI6MTc1NTI2NjU5NywiaWF0IjoxNzU1MjY1OTk3LCJvdHBfZXhwaXJlc19hdCI6MTc1NTI2NjI5NywidXVpZCI6Ik9xQWw0TUJ2V3FKR0wxVyJ9.NYHN-V5N3lYX_QaekSZ0IroWHP5t8g3KxQChriHnaOM",
	})

	if err != nil {
		log.Fatalf("error occoured %v", err)
	}
	log.Println("Response:", resp)
}

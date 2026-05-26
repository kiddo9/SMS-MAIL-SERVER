package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/kiddo9/SMS-MAIL-SERVER/message/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, _ := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	client := pb.NewFileUploadServicesClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	md := metadata.Pairs(
		"x-send-using", "email",
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	fileContent, err := os.ReadFile("bulk_messaging_template.xlsx")
	if err != nil {
		log.Fatalf("error occoured %v", err)
	}

	resp, err := client.FileUpload(ctx, &pb.FileUploadRequest{Date: time.Now().Format("2006-01-02"), Content: fileContent})

	if err != nil {
		log.Fatalf("error occoured %v", err)
	}
	log.Println("Response:", resp)
}

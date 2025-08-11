// package client

package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kiddo9/SMS-MAIL-SERVER/github.com/kiddo9/SMS-MAIL-SERVER/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	client := pb.NewAdminServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, _ := client.GetAdmin(ctx, &pb.GetAndValidateAdminRequest{Uuid: 1})
	log.Println("Response:", resp)
}

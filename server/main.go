package main

import (
	"log"
	"net"

	"github.com/kiddo9/SMS-MAIL-SERVER/handlers"
	pb "github.com/kiddo9/SMS-MAIL-SERVER/message/proto"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", ":9001")

	if err != nil {
		panic(err)
	}

	pb.RegisterAdminServiceServer(grpcServer, &handlers.AdminHandler{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
	log.Println("Server is running on port 50051")
}

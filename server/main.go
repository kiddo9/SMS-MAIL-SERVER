package main

import (
	"log"
	"net"

	"github.com/kiddo9/SMS-MAIL-SERVER/handlers"
	pb "github.com/kiddo9/SMS-MAIL-SERVER/message/proto"
	"github.com/kiddo9/SMS-MAIL-SERVER/middleware"
	"google.golang.org/grpc"
)


func main() {
	grpcServer := grpc.NewServer(middleware.Interceptors(middleware.RecaptchaMiddleware, middleware.AuthMiddleware))

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

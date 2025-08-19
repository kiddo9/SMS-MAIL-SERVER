package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/kiddo9/SMS-MAIL-SERVER/handlers"
	pb "github.com/kiddo9/SMS-MAIL-SERVER/message/proto"
	"github.com/kiddo9/SMS-MAIL-SERVER/middleware"
	"google.golang.org/grpc"
)

func init(){
	godotenv.Load(".env")
}


func main() {
	grpcServer := grpc.NewServer(middleware.Interceptors(middleware.RecaptchaMiddleware, middleware.AuthMiddleware))
	port := os.Getenv("PORT")

	lis, err := net.Listen("tcp",":"+port)

	runningMessage := fmt.Sprintf("Server is running on port %s", port)
	fmt.Println(runningMessage)

	if err != nil {
		panic(err)
	}

	pb.RegisterAdminServiceServer(grpcServer, &handlers.AdminHandler{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

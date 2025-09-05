package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/kiddo9/SMS-MAIL-SERVER/handlers"
	pb "github.com/kiddo9/SMS-MAIL-SERVER/message/proto"
	"google.golang.org/grpc"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
}

func main() {
	grpcServer := grpc.NewServer()
	port := os.Getenv("PORT")

	lis, err := net.Listen("tcp", ":"+port)

	runningMessage := fmt.Sprintf("Server is running on port %s, address includes %s", port, lis.Addr())
	fmt.Println(runningMessage)

	if err != nil {
		panic(err)
	}

	pb.RegisterAdminServiceServer(grpcServer, &handlers.AdminHandler{})
	pb.RegisterFileUploadServicesServer(grpcServer, &handlers.FileUploadStruct{})
	pb.RegisterTemplateServicesServer(grpcServer, &handlers.Temp{})
	pb.RegisterSmsServicesServer(grpcServer, &handlers.Wallet{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

package main

import (
	"log"
	"net"

	pb "github.com/kiddo9/SMS-MAIL-SERVER/github.com/kiddo9/SMS-MAIL-SERVER/proto"
	"google.golang.org/grpc"
)


type TestStructs struct {
	pb.UnimplementedAdminServiceServer
}

func (s *TestStructs) AllAdmin(*pb.AllAdmins, pb.AdminService_AllAdminServer) error {
	// Implementation of AllAdmin method
	return nil
}
func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAdminServiceServer(grpcServer, &TestStructs{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
	log.Println("Server is running on port 50051")
	// defer grpcServer.Stop()
	// defer lis.Close()
	// log.Println("Server stopped gracefully")
	// defer grpcServer.GracefulStop()
}

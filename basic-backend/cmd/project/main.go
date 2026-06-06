package main

import (
	"log"
	"net"
	
	"google.golang.org/grpc"


	pb "github.com/DmitriiS-dev/basic-backend/proto"
	"github.com/DmitriiS-dev/basic-backend/internal/service"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil{
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTaskServiceServer(
		grpcServer,
		&service.TaskServer{},
	)

	log.Println("gRPC Server Running on :50051")

	grpcServer.Serve(lis)
}

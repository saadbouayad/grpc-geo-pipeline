package main

import (
	"grpc-geo-pipeline/pb"
	"grpc-geo-pipeline/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSensorServiceServer(grpcServer, &server.SensorService{})

	log.Println("Server is running on localhost:50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

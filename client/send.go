package client

import (
	"context"
	"grpc-geo-pipeline/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SendSensorData(address string, data *pb.SensorData) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewSensorServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SendSensorData(ctx, data)
	if err != nil {
		log.Fatalf("Error sensing sensor data: %v", err)
	}
	log.Printf("Server response: %s\n", res.Message)
}

package server

import (
	"context"
	"grpc-geo-pipeline/pb"
	"log"
	"time"
)

// SensorService implements the gRPC SensorServiceServer interface
type SensorService struct {
	pb.UnimplementedSensorServiceServer
}

// SendSensorData handles incoming sensor data from the client
func (s *SensorService) SendSensorData(ctx context.Context, data *pb.SensorData) (*pb.SensorResponse, error) {
	log.Println("SendSensorData method called") // Debug log

	log.Printf("Received data from sensor %s: Latitude=%f, Longitude=%f, Temperature=%f, Timestamp=%s",
		data.SensorId, data.Latitude, data.Longitude, data.Temperature, data.Timestamp)

	// Simulate processing the data
	time.Sleep(500 * time.Millisecond)

	return &pb.SensorResponse{Message: "Data received successfully"}, nil
}

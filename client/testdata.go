package client

import (
	"grpc-geo-pipeline/pb"
	"time"
)

// GetTestData returns simulated sensor data for the client to send to the server
func GetTestData() []*pb.SensorData {
	return []*pb.SensorData{
		{
			SensorId:    "sensor-1",
			Latitude:    37.7749,
			Longitude:   -122.4194,
			Temperature: 22.5,
			Timestamp:   time.Now().Format(time.RFC3339),
		},
		{
			SensorId:    "sensor-2",
			Latitude:    34.0522,
			Longitude:   -118.2437,
			Temperature: 25.0,
			Timestamp:   time.Now().Format(time.RFC3339),
		},
	}
}

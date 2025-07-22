package main

import (
	"grpc-geo-pipeline/client"
	"log"
)

func main() {
	testData := client.GetTestData()

	for _, data := range testData {
		client.SendSensorData("localhost:50051", data)
	}

	log.Println("Client finished sending data.")
}

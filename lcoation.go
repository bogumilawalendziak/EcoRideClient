package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

func StartLocationStreaming(orderNumber string, startTime time.Time, locationUpdates <-chan Location) {

	for location := range locationUpdates {
		locationUpdate := LocationUpdate{
			OrderNumber: orderNumber,
			StartTime:   startTime,
			Latitude:    location.Latitude,
			Longitude:   location.Longitude,
			Timestamp:   time.Now().Unix(),
		}
		locationUpdateJson, err := json.Marshal(locationUpdate)
		if err != nil {
			fmt.Printf("Failed to marshal location update: %s\n", err)
			continue
		}
		fmt.Printf("Sending location update: %s\n", locationUpdateJson)

		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &config.UpdateLocationTopic, Partition: kafka.PartitionAny},
			Value:          locationUpdateJson,
		}, nil)

		fmt.Printf("Location update sent to \n" + config.UpdateLocationTopic)
		if err != nil {
			fmt.Printf("Failed to produce message: %s\n", err)
			continue
		}
	}
}

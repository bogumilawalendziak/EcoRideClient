package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

func sendReserveRequest(request ReserveRequest) error {
	fmt.Printf("Sending reservation request: %s\n", request)
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.BootstrapServers})
	if err != nil {
		fmt.Printf("Failed to create Kafka producer: %s\n", err)
		return err
	}
	requestJson, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("Failed to marshal reservation request: %s\n", err)
		return err
	}

	fmt.Printf("Sending reservation request: %s\n", requestJson)
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &config.ReservationRequestTopic, Partition: kafka.PartitionAny},
		Value:          requestJson,
	}, nil)

	if err != nil {
		fmt.Printf("Failed to produce message: %s\n", err)
		return err
	}

	fmt.Printf(config.ReservationRequestTopic)
	fmt.Printf("Reservation request sent to \n" + config.ReservationRequestTopic)
	producer.Flush(15 * 1000)

	return nil
}

func ListenForReservationResponses() {
	fmt.Printf("Listening for reservation responses\n")
	var err = consumer.SubscribeTopics([]string{config.ReservationResponseTopic}, nil)
	if err != nil {
		fmt.Printf("Failed to subscribe to topics: %s", err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		fmt.Printf("Received message: %s\n", msg)
		if err != nil {
			fmt.Printf("Consumer error: %s %s", err, msg)
			continue
		}

		var response ReserveResponse
		if err := json.Unmarshal(msg.Value, &response); err != nil {
			fmt.Printf("Failed to unmarshal ReserveResponse: %s", err)
			continue
		}

		startStreamingData(response)
	}
}
func startStreamingData(response ReserveResponse) {
	duration := 1 * time.Hour
	locationUpdates := make(chan Location)

	go SimulateLocationChanges(duration, locationUpdates)

	go StartLocationStreaming(response.OrderNumber, response.StartTime, locationUpdates)
}

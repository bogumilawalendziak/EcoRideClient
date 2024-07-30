package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"os"
)

var producer *kafka.Producer
var consumer *kafka.Consumer
var config = loadConfigFromEnv()

func initKafkaProducer(config Config) {
	var err error
	producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.BootstrapServers})
	if err != nil {
		fmt.Printf("Failed to create Kafka producer: %s\n", err)
		return
	}
}

func initKafkaConsumer(config Config) {
	var err error
	consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.BootstrapServers,
		"group.id":          config.GroupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Printf("Failed to create Kafka consumer: %s", err)
		return
	}
}

func loadConfigFromEnv() Config {

	godotenv.Load()

	bootstrapServers := os.Getenv("BOOTSTRAP_SERVERS")
	groupId := os.Getenv("GROUP_ID")
	reserveReqTopic := os.Getenv("TOPIC_RESERVATION_REQUEST")
	reserveRespTopic := os.Getenv("TOPIC_RESERVATION_RESPONSE")
	updateLocationTopic := os.Getenv("TOPIC_LOCATION_UPDATE")
	return Config{
		BootstrapServers:         bootstrapServers,
		GroupId:                  groupId,
		ReservationRequestTopic:  reserveReqTopic,
		ReservationResponseTopic: reserveRespTopic,
		UpdateLocationTopic:      updateLocationTopic,
	}
}

func initKafka() {
	config = loadConfigFromEnv()

	initKafkaProducer(config)
	initKafkaConsumer(config)
}

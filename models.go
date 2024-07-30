package main

import "time"

type ReserveRequest struct {
	UserId    string `json:"userId"`
	VehicleId string `json:"vehicleId"`
}

type LocationUpdate struct {
	OrderNumber string    `json:"order_number"`
	StartTime   time.Time `json:"start_time"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Timestamp   int64     `json:"timestamp"`
}

type Location struct {
	Latitude  float64
	Longitude float64
	Timestamp time.Time
}

type ReserveResponse struct {
	OrderNumber string    `json:"order_number"`
	StartTime   time.Time `json:"start_time"`
}

type Config struct {
	BootstrapServers         string
	GroupId                  string
	ReservationRequestTopic  string
	ReservationResponseTopic string
	UpdateLocationTopic      string
}

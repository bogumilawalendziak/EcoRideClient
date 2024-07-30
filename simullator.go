package main

import (
	"math/rand"
	"time"
)

const (
	initialLatitude  = 52.520008
	initialLongitude = 13.404954
	latitudeChange   = 0.0001
	longitudeChange  = 0.0001
)

func SimulateLocationChanges(duration time.Duration, locationUpdates chan<- Location) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	endTime := time.Now().Add(duration)
	for now := range ticker.C {
		if now.After(endTime) {
			return
		}

		location := Location{
			Latitude:  52.5200 + rand.Float64()*0.0001,
			Longitude: 13.4050 + rand.Float64()*0.0001,
			Timestamp: now,
		}

		locationUpdates <- location
	}
}

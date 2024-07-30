package main

import (
	"testing"
	"time"
)

func TestSimulateLocationChanges(t *testing.T) {
	duration := 2 * time.Second
	locationUpdates := make(chan Location, 10)

	go SimulateLocationChanges(duration, locationUpdates)

	time.Sleep(duration + time.Second)

	close(locationUpdates)

	var updatesCount int
	for range locationUpdates {
		updatesCount++
	}

	if updatesCount < 1 {
		t.Errorf("Expected at least 1 location update, got %d", updatesCount)
	}
}

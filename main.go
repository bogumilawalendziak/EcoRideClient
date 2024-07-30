package main

func main() {

	initKafka()

	go ListenForReservationResponses()
	initAPI()
	select {}
}

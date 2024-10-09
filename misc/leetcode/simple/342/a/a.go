package main

func findDelayedArrivalTime(arrivalTime, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}

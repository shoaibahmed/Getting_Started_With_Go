package main

import "fmt"

// Only accepts a channel for sending values
func ping(pings chan <- string, msg string) {
	pings <- msg
}

// Accepts one channel to retrieval (pings) and second one for sends (pongs)
func pong(pings <- chan string, pongs chan <- string) {
	msg := <- pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "Passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

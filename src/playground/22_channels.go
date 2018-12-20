package main

import "fmt"

func main() {
	messages := make(chan string) // Chan "value type"

	// Send ping message to channel
	go func() { messages <- "ping" }()

	// Recieve message from channel
	msg := <- messages

	// Both reading and writing from channel are blocking calls
	fmt.Println(msg)
}

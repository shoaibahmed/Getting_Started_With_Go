package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// The default clause results in non-blocking ops
	select {
	case msg := <- messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Sent message:", msg)
	default:
		fmt.Println("No message Sent")
	}

	select {
	case msg := <- messages:
		fmt.Println("Received message:", msg)
	case sig := <- signals:
		fmt.Println("Received signal:", sig)
	default:
		fmt.Println("No activity")
	}
}

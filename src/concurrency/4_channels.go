package main

import "fmt"
import "time"
import "math/rand"

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // Send value to the channel
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	// Initialize a channel
	var c chan string
	c = make(chan string)

	// Execute the go routine (Async)
	go boring("boring!", c)
	
	for i := 0; i < 5; i++ {
		fmt.Printf("Received from channel: %q\n", <- c) // Retrieve the sent value from the channel
	}

	fmt.Println("You're boring. I'm leaving.")
}

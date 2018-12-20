package main

import "fmt"
import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "One"
	} ()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "Two"
	} ()

	for i := 0; i < 2; i++ {
		// Selects the first one that gets ready or selects one at random if multiple are ready
		select {
		case msg1 := <- c1:
			fmt.Println("Received:", msg1)
		case msg2 := <- c2:
			fmt.Println("Received:", msg2)
		}
	}
}

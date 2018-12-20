package main

import "fmt"

func main() {
	// Two specfies the size of the buffer (can hold a maximum of two elements)
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

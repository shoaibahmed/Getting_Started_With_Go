package main

import "fmt"

func main() {
	queue := make(chan string, 2)

	queue <- "One"
	queue <- "Two"
	close(queue)

	// It's possible to retrieve the remaining values from a channel even after is it closed
	for elem := range queue {
		fmt.Println(elem)
	}

	// Retrives nothing this time since all the elements in the queue have been consumed
	for elem := range queue {
		fmt.Println("Second attempt:", elem)
	}
}

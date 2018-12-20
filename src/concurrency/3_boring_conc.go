package main

import "fmt"
import "time"
import "math/rand"

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	// Execute the go routine (Async)
	go boring("boring!")
	
	fmt.Println("I'm listening.")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring. I'm leaving.")
}


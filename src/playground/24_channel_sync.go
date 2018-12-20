package main

import "fmt"
import "time"

// Done specifies that the goroutine will notify another goroutine
func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<- done
}

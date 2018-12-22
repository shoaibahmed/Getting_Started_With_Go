package main

import "fmt"
import "time"

// First receiving channel and then sending channel
func worker(id int, jobs <- chan int, results chan <- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job #", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job #", j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Add a barrier forcing all results to be computed
	for a := 1; a <= 5; a++ {
		<- results
	}
}

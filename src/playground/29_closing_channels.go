package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// More will return true if all the jobs have been rec
			j, more := <- jobs
			if more {
				fmt.Println("Received job:", j)
			} else {
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	} ()

	for j := 1; j <= 10; j++ {
		jobs <- j
		fmt.Println("Sent job:", j)
	}
	close(jobs)
	fmt.Println("Sent all jobs")

	<- done
}

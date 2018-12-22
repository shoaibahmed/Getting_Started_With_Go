package main

import "fmt"
import "time"

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<- timer1.C // Wait for the internal channel from the timer
	fmt.Println("Timer # 01 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<- timer2.C
		fmt.Println("Timer # 02 expired")
	} ()

	// Stop the timer before it expires
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer # 02 stopped")
	}
}

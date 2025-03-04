package main

import (
	"fmt"
	"time"
)

// example 1
// func main() {
// 	ch1, ch2 := make(chan int), make(chan int)

// 	go func() {
// 		ch1 <- 42
// 	}()

// 	select {
// 	case val := <-ch1:
// 		fmt.Printf("Got %d from ch1\n", val)
// 	case val := <-ch2:
// 		fmt.Printf("Got %d from ch2\n", val)
// 	}
// }

// example 2
func main() {
	out := make(chan float64)

	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()
	select {
	case val := <-out:
		fmt.Println("Got value from out:", val)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout")

	}
}

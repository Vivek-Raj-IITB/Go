package main

import (
	"fmt"
	"time"
)

// example 1
// func main() {
// 	ch := make(chan string)
// 	// ch <- "Hello" // This will cause a deadlock
// 	go func() {
// 		ch <- "Hello" // This will not cause a deadlock
// 	}()

// 	val := <-ch
// 	println(val)
// }

// example 2
// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		for i := 0; i < 3; i++ {

// 			time.Sleep(2 * time.Second)
// 			fmt.Println("Sending:", i)
// 			ch <- i

// 			// time.Sleep(time.Second)
// 		}
// 	}()
// 	for i := 0; i < 3; i++ {
// 		t2 := time.Now()
// 		val := <-ch
// 		fmt.Println("Time taken for val:", val, time.Since(t2))
// 	}
// }

// example 3
func main() {
	ch := make(chan int, 3) // buffered channel

	go func() {
		for i := 0; i < 3; i++ {

			time.Sleep(2 * time.Second)
			fmt.Println("Sending:", i)
			ch <- i

			// time.Sleep(time.Second)
		}
	}()
	time.Sleep(7 * time.Second)
	for i := 0; i < 3; i++ {

		t2 := time.Now()
		val := <-ch
		fmt.Println("Time taken for val:", val, time.Since(t2))
	}
}

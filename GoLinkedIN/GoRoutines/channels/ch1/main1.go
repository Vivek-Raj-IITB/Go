package main1

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendMessage(ch1)
	go receiveMessage(ch2)

	time.Sleep(3 * time.Second) // Wait for goroutines to finish
	select {
	case <-ch1:
		fmt.Println("Message received from ch1")
	case <-ch2:
		fmt.Println("Message received from ch2")

	}
	// select {
	// case <-ch2:
	// 	fmt.Println("Message received from ch2", <-ch2)
	// 	// case <-ch2:
	// 	// 	fmt.Println("Message received from ch2", <-ch2)
	// }

}

func sendMessage(ch chan string) {
	ch <- "Hello from goroutine send msg"
}

func receiveMessage(ch chan string) {
	ch <- "Hello from goroutine receive msg"
}

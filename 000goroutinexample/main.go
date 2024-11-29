package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond) // Simulate some work
	}
}

func main() {
	// Run printNumbers function in a separate Goroutine
	go printNumbers()

	// Main Goroutine continues execution
	for i := 'A'; i <= 'E'; i++ {
		fmt.Printf("%c\n", i)
		time.Sleep(500 * time.Millisecond) // Simulate some work
	}

	// Give the Goroutine time to finish
	// time.Sleep(3 * time.Second)
}

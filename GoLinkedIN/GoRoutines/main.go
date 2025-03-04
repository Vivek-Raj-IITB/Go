package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%c\n", i)
	}
}

func main() {
	go printNumbers()
	go printLetters()

	// Wait for goroutines to finish
	time.Sleep(6 * time.Second)
	fmt.Println("Done")
}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")
	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	mynum := 23
	var ptr *int = &mynum

	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value at pointer is ", *ptr)

}

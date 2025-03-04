package main

import "fmt"

// Hello returns a greeting for the named person.
func main() string {
	// Return a greeting that embeds the name in a message.
	println("hre in greetings functions cheking")
	message := fmt.Sprintf("Hi, %v. Welcome!")

	v := 42
	switch {
	case v == 42:
		fmt.Println(42)
		fallthrough
	case v < 100:
		fmt.Println(100)
		fallthrough
	case v > 42:
		fmt.Println(42)
		fallthrough
	case v == 42:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("default")
	}
	// Output:
	// 42
	// 1
	// default
	return message
}

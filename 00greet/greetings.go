package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	println("hre in greetings functions")
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

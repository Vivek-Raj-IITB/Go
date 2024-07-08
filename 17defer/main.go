package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer")

	defer fmt.Println("This is defer1")
	defer fmt.Println("This is defer2") // follow lifo using stack
	fmt.Println("This is not defer and will be printed first")
	myfun()

}

// defer is used to execute a function at the end of the functionS

func myfun() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, " ")
	}
}

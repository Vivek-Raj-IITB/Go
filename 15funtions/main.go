package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	// fmt.Println(strings.Join(os.Args[1:], " "))
	// fmt.Println(strings.Join(os.Args[1:], " "))()  {
	sayHello()

	result := adder(3, 5)
	fmt.Println("Result is ", result)
	val, mess := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Pro result is ", val, mess)

}

func sayHello() {
	fmt.Println("Hello")
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(numbers ...int) (int, string) {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total, "Hi"
}

// func main() {
// 	// anonymous function
// 	func() {
// 		fmt.Println("Hello")
// 	}()
// }

// anonymous functions exist in go lang as well
// func() {
// 	fmt.Println("Hello")
// }

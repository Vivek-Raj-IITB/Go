package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array")
	var arr [4]string
	arr[0] = "Hello"
	arr[1] = "World"
	arr[3] = "World"
	fmt.Println(arr)
	var veg = [3]string{"a", "b", "c"}
	fmt.Println(veg)
}

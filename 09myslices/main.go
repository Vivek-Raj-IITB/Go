package main

import (
	"fmt"
)

type person struct {
	name []interface{}
}

func main() {
	// fmt.Println("Welcome to Slices")
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
		goto default
	case "Wednesday":
		fmt.Println("Today is Wednesday")
	default:
		fmt.Println("Another day")
	}
	// abc := []interface{}{}
	// p := person{name: nil}
	// if p.name == nil {
	// 	fmt.Println("p.name is nil")
	// }
	// abc = p.name
	// fmt.Println(abc)

}
func updateSlice(s []string) {
	s[0] = "Bye"
}

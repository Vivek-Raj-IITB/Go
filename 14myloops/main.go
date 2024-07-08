package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ { //normal for loop
	// 	fmt.Println(days[d])
	// }

	// for i := range days {  // similar to python
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days { // similar to python
	// 	fmt.Printf("Index is %v and value is %v\n", index, day)
	// }

	var index int = 1
	for index < 7 { // while loop
		if index == 2 {
			index++
			goto label
		}
		fmt.Println(days[index])
		index++
	}
label:
	fmt.Println("Hi Gaurav")
}

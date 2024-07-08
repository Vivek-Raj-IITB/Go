package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var arr_list = []string{"a", "b", "c"}

	fmt.Printf("type of slices is %T\n ", arr_list)
	arr_list = append(arr_list, "d")
	fmt.Println(arr_list)
	arr_list = append(arr_list[1:]) // similar to pyhton
	fmt.Println(arr_list)
	updateSlice(arr_list)
	fmt.Println(arr_list)

	highScore := make([]int, 4)
	highScore[0] = 11
	highScore[1] = 20
	highScore[2] = 378
	highScore[3] = 41
	fmt.Println(highScore)
	highScore = append(highScore, 5)
	fmt.Println(highScore)
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))
	// fmt.Println(sort.IsSorted(highScore)) //wrong

	// how to remove values based on index

	var courses = []string{"math", "science", "english", "hindi", "social"}
	fmt.Println(courses)
	index := 2 // var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
func updateSlice(s []string) {
	s[0] = "Bye"
}

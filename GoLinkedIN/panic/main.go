package main

import "fmt"

func main() {
	vals := []int{1, 2, 3}
	v, err1 := safeValue(vals, 10)
	if err1 != nil {
		fmt.Println("Error:", err1)
	}
	fmt.Println("Value:", v)

}

func safeValue(vals []int, index int) (n int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		} else {
			fmt.Println("No panic")
		}
	}()

	return vals[index], nil
}

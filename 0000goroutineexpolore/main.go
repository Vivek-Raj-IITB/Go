package main

import (
	"fmt"
)

func main() {
	v := 42
	switch {
	case v == 42:
		fmt.Println(42)
		fallthrough
	case v < 100:
		fmt.Println(100)
		fallthrough
	case v > 4777:
		fmt.Println(426666)
		fallthrough
	case v == 43:
		fmt.Println(43)
		fallthrough
	default:
		fmt.Println("default")
	}
}

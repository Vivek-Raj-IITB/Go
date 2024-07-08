package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza app")
	fmt.Println("please rate bewteen 1 and 5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // comma ok syntax
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Updated rating is ", numRating+1)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("How are you")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give a a rating")
	input, _ := reader.ReadString('\n')
	fmt.Println("thank for your rating - ", input)
	fmt.Printf("typr of input %T \n", input)

}

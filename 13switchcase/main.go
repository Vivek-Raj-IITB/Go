package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to switch cases")

	num := rand.Intn(6) + 1

	fmt.Println("Valur of dice ", num)
	// in go we explicitly need to mention fallthrough for all the case below a case to run
	switch num {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spot")
		fallthrough
	case 6:
		fmt.Println("You can move 6 spot and roll dice again")
	default:
		fmt.Println("What was that!")

	}
}

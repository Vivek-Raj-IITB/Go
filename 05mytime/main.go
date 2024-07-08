package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time module")
	timeNow := time.Now()
	fmt.Println(timeNow)
	fmt.Printf("%T \n", timeNow)
	fmt.Println(timeNow.Format("02-01-2006 03:04:05 Monday"))
	createdDate := time.Date(2020, time.May, 12, 12, 12, 12, 12, time.UTC)
	fmt.Println(createdDate.Format("2006-02-01"))
}

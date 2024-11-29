package main

import (
	"errors"
	"fmt"
	"greetings"
)

const (
	PI = 3.14
)

func Foo() (string, error) {
	return "", errors.New("hi")
}
func Foo2() (string, error) {
	return "Here is the name", nil
}
func Foo3() (string, error) {
	return "", errors.New("From Foo3")
}
func main() {
	fmt.Println("Hello world")
	println(greetings.Hello("Vivek"))
	println(PI)

	name, err := Foo3()
	fmt.Println("for Foo3()", name, err)
	name, err = Foo()
	fmt.Println("for Foo3______@2()", name, err)
	if err != nil {
		name, err = Foo2()
		fmt.Println("printherrrrrrrr ", name)
		fmt.Println(err)
	}
	fmt.Println("blalking", name)
}

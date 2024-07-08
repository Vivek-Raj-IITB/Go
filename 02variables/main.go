package main

import "fmt"

const LoginCre = "ksdnkjsn" // Public data

func main() {

	var name string = "Vivek"
	fmt.Printf("%s\n", name)
	fmt.Printf("%T\n", name)

	var islog bool = true
	fmt.Println(islog)
	fmt.Printf("%T\n", islog)

	var smallVal uint = 12
	fmt.Println(smallVal)
	fmt.Printf("%T\n", smallVal)

	var smallFloat float32 = 12.65416576546465465446546516
	fmt.Println(smallFloat)
	fmt.Printf("%T\n", smallFloat)

	var smallFloat1 float32
	fmt.Println(smallFloat1)
	fmt.Printf("%T\n", smallFloat1)

	var name1 string
	fmt.Printf("%s\n", name1)
	fmt.Printf("%T\n", name1)

	varr := 123
	fmt.Println(varr)
	fmt.Printf("%T\n", varr)

	fmt.Println(LoginCre)
	fmt.Printf("%T\n", LoginCre)
}

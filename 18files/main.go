package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcome to myfiles")
	content := "This needs to go in files"
	file, err := os.Create("./myfile.txt")

	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is ", length)
	defer file.Close()
	readFile()

}

func readFile() {
	databyte, err := os.ReadFile("./myfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("The file is : ", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}


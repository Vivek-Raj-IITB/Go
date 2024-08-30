package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Welcoem to web verb ")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "www.google.com"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}

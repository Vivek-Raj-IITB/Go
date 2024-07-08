package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://en.wikipedia.org/wiki/Software_engineering"

func main() {
	fmt.Println("Welcome to Web Requests")
	response, err := http.Get(url) //Get method is used to make a get request to the url
	//This method returns a response and an error
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T \n", response) //We can use %T to get the type of the response
	defer response.Body.Close()
	//We can use the response to get the status code
	//The response is a pointer to the http response
	//We need to close the response after we are done with it

	databyte, err := io.ReadAll(response.Body) //We can use io.ReadAll to read the response body
	//This method returns a byte slice and an error
	if err != nil {
		panic(err)
	}
	content := string(databyte)

	fmt.Println(content)
}

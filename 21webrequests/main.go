package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcoem to web verb ")
	// PerformGetRequest()
	// PerformJsonRequest()
	PerformFormRequest()
}

func PerformGetRequest() {
	const myurl = "https://dummy.restapiexample.com/api/v1/employee/1"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(string(content))
	fmt.Println("Bytecount is: ", byteCount)
	fmt.Println(responseString.String())

}

func PerformJsonRequest() {
	const myurl = "https://dummy.restapiexample.com/api/v1/create"
	requestBody := strings.NewReader(`
		{
			"name":"test",
			"salary":"123",
			"age":"23"
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformFormRequest() {
	fmt.Println("Perform Form Request")
	const myurl = "https://dummy.restapiexample.com/api/v1/create"
	data := url.Values{}
	data.Add("name", "test")
	data.Add("salary", "123")
	data.Add("age", "23")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

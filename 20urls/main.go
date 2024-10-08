package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=dskjnskj"

func main() {
	fmt.Println("Welcome to Urls")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	fmt.Println("jhbdhb", result)

	fmt.Println("Scheme ", result.Scheme)
	fmt.Println("Host ", result.Host)
	fmt.Println("Path ", result.Path)
	fmt.Println("Port ", result.Port())
	fmt.Println("Query ", result.RawQuery)

	qparams := result.Query()

	fmt.Printf("Type of Query params are: %T\n", qparams)
	// fmt.Println("The params are: ", qparams)

	for key, val := range qparams {
		fmt.Println("Key is: ", key, "Value is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawPath:  "user/sanket",
		RawQuery: "coursename=reactjs&paymentid=dskjnskj",
	}
	fmt.Println(partsOfUrl.String())
	fmt.Println(partsOfUrl)
}

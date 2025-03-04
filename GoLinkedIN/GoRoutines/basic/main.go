package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func siteSerial(urls []string) {
	for _, url := range urls {
		returnType(url)
	}
}
func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	fmt.Println("Content Type:", ctype)

}

func sitesConcurrent(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main() {

	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
	}
	start := time.Now()
	sitesConcurrent(urls)
	fmt.Println("Serial:", time.Since(start))

}

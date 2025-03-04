package main

import (
	"context"
	"fmt"
	"time"
)

type Bid struct {
	AdURL string
	Price float64
}

func bestBid(url string) Bid {
	time.Sleep(20 * time.Millisecond)
	return Bid{
		AdURL: "http://www.google1.com",
		Price: 0.05,
	}
}

var defaultBid = Bid{
	AdURL: "http://www.default.com",
	Price: 0.02,
}

func findBid(ctx context.Context, url string) Bid {

	ch := make(chan Bid, 1)
	go func() {
		ch <- bestBid(url)
	}()
	select {
	case <-ctx.Done():
		return defaultBid
	case bid := <-ch:
		return bid
	}
}
func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	url := "http://www.google.com"
	bid := findBid(ctx, url)
	fmt.Println("Bid:", bid)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	url = "http://www.yahoo.com"
	bid = findBid(ctx, url)
	fmt.Println("Bid:", bid)

}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type url string
type httpGet struct{}

func (httpGet) Get(url url) {
	a := rand.Intn(100)
	time.Sleep(time.Duration(a) * time.Millisecond)
	fmt.Println(url)
}

var http httpGet

func main() {
	var wg sync.WaitGroup
	var urls = []url{
		"http://www.github.com/alamin-mahamud",
		"http://www.google.com",
		"http://www.somebalchal.com",
	}

	for _, ur := range urls {
		// Increment the WaitGroup Counter
		wg.Add(1)

		// Launch a Goroutine to fetch the URL
		go func(u url) {
			defer wg.Done()
			http.Get(u)
		}(ur)

	}

	// wait for all http fetches to complete
	wg.Wait()
}

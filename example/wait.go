package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{"http://retailmenot.com", "http://retailmeow.com"}

func main() {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			http.Get(url)
			fmt.Println("Fetched", url)
			wg.Done()
		}(url)
	}
	//wg.Wait()
	fmt.Println("All urls fetched")
}

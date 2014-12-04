package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	ch <- "ping"
	fmt.Println(<-ch)
}

package main

import "fmt"

func main() {
	ch := make(chan string)
	ch <- "ping"
	fmt.Println(<-ch)
}

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	fmt.Println(from)
}

func main() {

	go f("goroutine")
	go f("goroutine2")
	f("direct")

	time.Sleep(100 * time.Millisecond)
}

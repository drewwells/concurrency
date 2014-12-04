package main

import "fmt"

func f(ch chan string, msg string) {
	for {
		ch <- msg
	}
}

func main() {
	p, o := make(chan string), make(chan string)
	go f(p, "ping")
	go f(o, "pong")
	for i := 0; i < 8; i++ {
		select { // Non-blocking select across multiple channels
		case msg := <-p:
			fmt.Println(msg)
		case msg := <-o:
			fmt.Println(msg)
		}
	}
	fmt.Println("finished")
}

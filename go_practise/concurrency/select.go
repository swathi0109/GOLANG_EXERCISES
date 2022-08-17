package main

import (
	"fmt"
	"time"
)

func readdata1(c1 chan string) {
	c1 <- "one"
	time.Sleep(3 * time.Millisecond)
}

func readdata2(c2 chan string) {
	c2 <- "two"
	time.Sleep(3 * time.Millisecond)
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go readdata1(c1)
	go readdata2(c2)

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			time.Sleep(3 * time.Millisecond)
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		default:
			fmt.Println("no msg")
		}
	}
}

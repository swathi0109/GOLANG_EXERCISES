package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; i < 4; i++ {
		c <- "ping" //write "ping to channel"
	}
}

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	time.Sleep(2 * time.Millisecond)
	fmt.Println(<-c)
}

/*package main

import (
	"fmt"
)

func main() {
	n := 3

	out := make(chan int)
	go multiplyByTwo(n, out)

	fmt.Println(<-out)
}

// This function now accepts a channel as its second argument...
func multiplyByTwo(num int, out chan<- int) {
	result := num * 2

	//... and pipes the result into it
	out <- result
}*/

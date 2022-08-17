package main

import (
	"fmt"
	"time"
)

func start() {
	go start2()
	fmt.Println("Started")
}

func start2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
}

func main() {
	go start()
	fmt.Println("Main started")
	time.Sleep(2 * time.Millisecond)
	go start2()
	fmt.Println()
	fmt.Println("Start2() called again")
	time.Sleep(2 * time.Millisecond)
	fmt.Println("Main ended")
}

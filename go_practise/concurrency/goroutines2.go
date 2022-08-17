package main

import (
	"fmt"
	"time"
)

func printstring(n string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(n)
	}
}

func main() {
	fmt.Println("Started")
	go printstring("swathi")
	printstring("shekar")
	go greeter("hello")
	greeter("world")
	fmt.Println("Stopped main")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

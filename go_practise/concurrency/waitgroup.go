package main

import (
	"fmt"
	"sync"
)

func rt(wg *sync.WaitGroup) {
	fmt.Println("First go routine")
	wg.Done()
}

func rt2(wg *sync.WaitGroup) {
	fmt.Println("Second goroutine")
	wg.Done()
}

func rt3(wg *sync.WaitGroup) {
	fmt.Println("Third goroutine")
	wg.Done()
}

func main() {
	//creating waitgroup
	wg := new(sync.WaitGroup)
	//adding the goroutines to waitgroup
	wg.Add(2)
	go rt2(wg)
	go rt(wg)
	go rt3(wg)
	//block the execution until the counter value is 0
	wg.Wait()
}

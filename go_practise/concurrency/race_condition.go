package main

import (
	"fmt"
	"sync"
)

var a = 0

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	a = a + 1
	m.Unlock()

	wg.Done()
}

func main() {
	fmt.Println("Trying to get race condition")
	var wg sync.WaitGroup

	var m sync.Mutex
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	}
	wg.Wait()
	fmt.Println("Value of a is ", a)

}

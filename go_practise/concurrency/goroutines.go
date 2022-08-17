package main

import (
	"fmt"
)

// notice we've not changed anything in this function
// when compared to our previous sequential program
func compute(value int) {
	for i := 0; i < value; i++ {
		//
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Goroutine Tutorial")

	// notice how we've added the 'go' keyword
	// in front of both our compute function calls
	go compute(10)
	go compute(10)
	//time.Sleep(2 * time.Millisecond)

	var input string
	fmt.Scanln(&input)
}

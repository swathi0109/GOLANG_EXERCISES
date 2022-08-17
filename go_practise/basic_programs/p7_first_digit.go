package main

import (
	"fmt"
)

func main() {
	n := 236
	for n >= 10 {
		n = n / 10
	}
	fmt.Println("first digit is ", n)
}

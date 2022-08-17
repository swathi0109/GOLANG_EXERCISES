package main

import "fmt"

func factorial(num int32) int32 {
	if num == 0 || num == 1 {
		return 1
	} else if num > 1 {
		return num * factorial(num-1)
	} else {
		return -1
	}
}

func main() {
	fact := factorial(4)
	fmt.Println(fact)
}

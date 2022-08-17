package main

import "fmt"

func main() {
	fmt.Println("Calling the function")
	defer fmt.Println("The function is called successfully")
	sum := calcsum(5, 6)
	fmt.Println("The sum of 2 numbers is ", sum)
}

func calcsum(a, b int) int {
	return a + b
}

package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number")
	fmt.Scanln(&num)

	cube := num * num * num
	fmt.Printf("Cube of a number is %d", cube)
}

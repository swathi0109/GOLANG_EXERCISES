package main

import "fmt"

func main() {
	var num, count int
	fmt.Println("Enter any number: ")
	fmt.Scanln(&num)

	for num > 0 {
		num = num / 10
		count = count + 1
	}
	fmt.Println("number of digits is ", count)
}

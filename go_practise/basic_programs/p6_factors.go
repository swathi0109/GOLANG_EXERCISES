package main

import "fmt"

func main() {
	var a []int
	var n int
	fmt.Println("Enter the number to find its factors:")
	fmt.Scanln(&n)
	for i := 1; i <= 24; i++ {
		if n%i == 0 {
			a = append(a, i)
		}
	}
	fmt.Println(a)
}

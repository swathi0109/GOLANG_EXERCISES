//A function that is called with varying number of arguments is called variadic function
package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, "")
	total := 0
	for _, i := range nums {
		total += i
	}
	fmt.Println(" :", total)
}

func main() {
	sum(1, 2, 3)
	sum(10, 11, 12, 13)
	arr := []int{3, 5, 7, 4, 7, 8, 9}
	sum(arr...)
}

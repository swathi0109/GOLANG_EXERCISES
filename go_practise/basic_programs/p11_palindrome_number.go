package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number")
	fmt.Scanln(&n)
	rev := 0
	for temp := n; temp > 0; {
		dig := temp % 10
		rev = rev*10 + dig
		temp = temp / 10
	}
	if n == rev {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not palindrome")
	}
}

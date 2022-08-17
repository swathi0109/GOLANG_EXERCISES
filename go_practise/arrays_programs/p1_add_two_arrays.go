package main

import "fmt"

func main() {
	var size int
	fmt.Println("Enter the array size")
	fmt.Scanln(&size)

	//array_1 := make([]int, size)
	array_2 := make([]int, size)
	/*add_array := make([]int, size)
	sub_array := make([]int, size)
	mul_array := make([]int, size)
	div_array := make([]int, size)

	fmt.Println("Enter the array_1 elements")
	for i := 0; i < size; i++ {
		fmt.Scan(&array_1[i])
	}*/

	fmt.Println("Enter the array_2 elements")
	for i := 0; i < size; i++ {
		fmt.Scan(&array_2[i])
	}

	/*for i := 0; i < size; i++ {
		add_array[i] = array_1[i] + array_2[i]
		sub_array[i] = array_1[i] - array_2[i]
		mul_array[i] = array_1[i] * array_2[i]
		div_array[i] = array_1[i] / array_2[i]
	}
	fmt.Println("Addition of arrays is ", add_array)
	fmt.Println("Subtraction of arrays is ", sub_array)
	fmt.Println("Multiplicaion of arrays is ", mul_array)
	fmt.Println("Division of arrays is ", div_array)
	fmt.Println()

	//find average
	sum := 0
	for i := 0; i < size; i++ {
		sum = sum + array_1[i]
	}
	average := sum / size
	fmt.Println("Average is ", average)
	fmt.Println()*/

	//count even and odd numbers
	/*count_even := []int{}
	count_odd := []int{}
	for i := range array_2 {
		if array_2[i]%2 == 0 {
			count_even = append(count_even, array_2[i])
		} else {
			count_odd = append(count_odd, array_2[i])
		}
	}
	fmt.Printf("No of even numbers = %d\n and No of odd numbers = %d", count_even, count_odd)*/
	fmt.Println()

	//Program to find largest and smallest

	/*largest := array_2[0]
	smallest := array_2[0]

	for i := range array_2 {
		if array_2[i] > largest {
			largest = array_2[i]
		}
		if array_2[i] < smallest {
			smallest = array_2[i]
		}
	}
	fmt.Printf("Largest is %d and smallest is %d", largest, smallest)
	fmt.Println()*/

	//print elements in even and odd index position
	for i := 0; i < size; i = i + 2 {
		fmt.Println("Even index array elements are ", array_2[i])
	}
	for i := 1; i <= size; i = i + 2 {
		fmt.Println("Odd index array elements are ", array_2[i])
	}
}

package main

import "fmt"

func get_avg(a []int, size int) float32 {
	var i, sum int
	var avg float32
	sum = 0

	for i = 0; i < size; i++ {
		sum += a[i]
	}
	avg = float32(sum / size)
	return avg
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	avg := get_avg(a, 5)
	fmt.Println("avg =", avg)
}

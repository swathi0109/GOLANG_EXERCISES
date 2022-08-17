package main

import "fmt"

func main() {
	//Declaration of array
	var a [4]int
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	fmt.Println(a)
	b := []int{1, 2, 3, 4}
	fmt.Println(b)
	fmt.Println()

	//Slicing
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%T", a)
	fmt.Println()
	var myslice []int = arr[2:8]
	fmt.Printf("%T", myslice)
	fmt.Println()
	fmt.Println(len(myslice), cap(myslice), myslice)

	new_slice := myslice[0:4]
	fmt.Println(len(new_slice), cap(new_slice), new_slice)

	new_slice1 := new_slice[0:2]
	fmt.Println(len(new_slice1), cap(new_slice1), new_slice1)
	new_slice1 = append(new_slice1, 9)
	fmt.Println(new_slice1)
	var c []int = []int{1, 2, 3, 4}
	if b == c {
		fmt.Println("true")
	}
	/*arr = arr[0:6]
	fmt.Println("len(arr) =", len(arr), "cap(arr) =", cap(arr), "arr =", arr)*/
	//b := a[1:4]
	//fmt.Println("len(a) = %d and cap(a) = %d", len(b), cap(b), b)//

}

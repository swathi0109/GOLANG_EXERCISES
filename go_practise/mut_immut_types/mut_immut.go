package main

import "fmt"

func main() {
	var x []int = []int{2, 3, 4}
	y := x
	y[2] = 6
	fmt.Println(x, y)
	/*Output: [2 3 6] [2 3 6] - slice is specified without size []int. It is mutable datatype.
	If any value in x changes, then in y also it changes*/

	var a map[string]int = map[string]int{"a": 10, "b": 11, "c": 12}
	b := a
	a["b"] = 15
	fmt.Println(a, b)
	/*Output: map[a:10 b:15 c:12] map[a:10 b:15 c:12] - map is mutable datatype.
	If any value in a changes, then in b also it changes*/

	var p [3]int = [3]int{7, 8, 9}
	q := p
	q[0] = 6
	fmt.Println(p, q)
	/*Output: [7 8 9] [6 8 9] - array is specified with size [size]int. It is an immutable datatype.
	If any value in p changes, then value of q does not change*/
}

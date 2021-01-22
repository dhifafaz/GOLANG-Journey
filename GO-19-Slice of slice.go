package main

import "fmt"

func main() {
	// slice of slice
	//matrix
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	a = append(a, 5)
	fmt.Println(a)

	b := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(b)

}

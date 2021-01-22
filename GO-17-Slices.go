package main

import "fmt"

func main() {
	// first way
	// var n []int
	// fmt.Println(n)
	// n = append(n, 1, 2, 3, 4, 5, 6, 7, 8)
	// fmt.Println(n)

	//second way to create a slices
	n := [5]int{1, 2, 3, 4, 5}
	var a []int = n[0:3] //assign value of array a with  value of array n from index 0 to 2 (3 index)
	fmt.Println(a)
	a[0] = 100
	fmt.Println(a)
}

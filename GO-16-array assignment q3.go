package main

import "fmt"

func main() {
	//calculate array length
	a := []int{1, 2, 4, 5, 6, 7}
	fmt.Println(len(a))

	// find sum of all no in array
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	fmt.Println(sum)
}

package main

import "fmt"

func main() {
	// Array find the largest number
	var n int
	var largest int = 0
	fmt.Scan(&n)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println(array)
	for i := 0; i < n; i++ {
		if array[i] >= largest {
			largest = array[i]
		}
	}
	fmt.Println(largest)

}

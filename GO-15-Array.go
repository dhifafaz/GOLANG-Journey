package main

import "fmt"

func main() {
	// Array
	var n [5]int
	for i := 0; i < 5; i++ {
		n[i] = i + 10
	}
	fmt.Println(n[0], n[4])

	z := [10]int{2, 3, 4, 5}
	fmt.Println(z)

}

package main

import "fmt"

func main() {
	//copy a slice to another slice
	n := []int{1, 2, 3, 4, 5}
	var n1 []int
	for i := 0; i < len(n); i++ {
		n1 = append(n1, n[i])
	}
	fmt.Println(n1)

}

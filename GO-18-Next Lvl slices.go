// nil slices
// slice default
// capacity and length

package main

import "fmt"

func main() {
	// null slice
	var a []int
	a = append(a, 2)
	if a == nil {
		fmt.Println("slice is null")
	} else {
		fmt.Println("slice not null")
	}

	// slice default

	b := [5]int{1, 2, 3, 4, 5}
	var c []int = b[:5] // slices==> array variable [lower bound : higher bound]
	// the default[0:len of array]
	fmt.Println(c)

	// capacity and length
	//to calculate the capacity is len array - lower bound of slice
	//the length it selfs is the length of the slice wich is the higehr bound
}

package main

import "fmt"

func main() {
	//print slice using loop
	n := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}

}

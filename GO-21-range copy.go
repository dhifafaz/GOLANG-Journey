package main

import "fmt"

func main() {
	// this is how you can write down the range in GO
	a := []int{12, 34, 21, 32}
	for i, v := range a {
		fmt.Println("Index", i, "Value", v)
	}
	b := []int{12, 34, 21, 32}
	for _, v := range b {
		fmt.Println("Value", v)
	}
	c := []int{12, 34, 21, 32}
	for i := range c {
		fmt.Println("Index", i)
	}
}

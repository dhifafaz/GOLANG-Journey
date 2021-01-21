package main

import "fmt"

func main() {
	// Go only support for loop that can be implemented widely
	for counter := 1; counter <= 100; counter++ {
		fmt.Println(counter)
	}
	// you can use for loop as this
	counter2 := 0
	for counter2 <= 10 {
		fmt.Println(counter2)
		counter2++
	}
	// you can use for loop as this
	counter3 := 10
	for ; counter3 >= 0; counter3-- {
		fmt.Println(counter3)
	}
}

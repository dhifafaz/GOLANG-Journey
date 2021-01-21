package main

import "fmt"

func main() {
	// to count how many digits in given number
	var input int
	fmt.Scanln(&input)
	var counter int = 0
	for input != 0 {

		counter++
		input /= 10
	}
	fmt.Print(counter)
}

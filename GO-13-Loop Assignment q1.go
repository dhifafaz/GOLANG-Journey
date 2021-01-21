package main

import "fmt"

func main() {
	// calculate factorial
	var number int
	var factorial int = 1
	fmt.Scanln(&number)
	for i := 1; i <= number; i++ {
		factorial *= i
	}
	fmt.Println(factorial)
}

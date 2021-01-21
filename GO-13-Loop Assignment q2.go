package main

import "fmt"

func main() {
	// to reverse a ngiven number
	var reversedNumber, input int
	fmt.Scanln(&input)
	for input != 0 {
		reversedNumber = input % 10
		fmt.Print(reversedNumber)
		input /= 10
	}
}

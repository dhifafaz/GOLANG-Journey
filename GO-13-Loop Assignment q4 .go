package main

import "fmt"

func main() {
	// to check if given number is palindrome
	var lastDigit, input, note int
	fmt.Scanln(&input)
	var newNumber int = 0
	note = input
	for input != 0 {
		lastDigit = input % 10
		newNumber = newNumber*10 + lastDigit
		input /= 10

	}
	if newNumber == note {
		fmt.Println("Yes it is a palindrome")
	} else {
		fmt.Println("No it is not a palindrome")
	}
}

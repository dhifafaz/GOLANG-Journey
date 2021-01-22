package main

import "fmt"

func main() {
	//reverse every number in slice

	n := []int{121, 432, 233, 324, 125}
	var reversedNumber, lastDigit int

	for i := 0; i < len(n); i++ {
		for n[i] != 0 {
			lastDigit = n[i] % 10
			reversedNumber = reversedNumber*10 + lastDigit
			n[i] /= 10
		}
		n[i] = reversedNumber
		reversedNumber = 0
	}
	fmt.Println(n)
}

package main

import "fmt"

func main() {
	var number1, number2, number3 int
	fmt.Scanln(&number1, &number2, &number3)
	//fmt.Scanln(&number2)
	//fmt.Scanln(&number3)
	if number1 > number2 && number1 > number3 {
		fmt.Println(number1)
	} else if number2 > number1 && number2 > number3 {
		fmt.Println(number2)
	} else if number3 > number1 && number3 > number2 {
		fmt.Println(number3)
	}
}

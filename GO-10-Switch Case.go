package main

import "fmt"

func main() {
	var input int
	fmt.Println("Enter a number")
	fmt.Scanln(&input)
	switch input {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Enter between 1 to 3")
	}

	var inputString string
	fmt.Println("Enter a color")
	fmt.Scanln(&inputString)
	switch inputString {
	case "red":
		fmt.Println("Red")
	case "green":
		fmt.Println("Green")
	case "blue":
		fmt.Println("Blue")
	default:
		fmt.Println("Enter only RGB color")
	}
}

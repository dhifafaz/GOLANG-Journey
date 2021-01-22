package main

import "fmt"

func main() {
	//store name only when the first letter is capital
	var a string
	a = "Dhifaf"
	fmt.Println(a[0]) //concept of string as an array of char in ASCII number

	var name []string
	var input string
	for {
		fmt.Println("Enter name with first",
			"Letter capital")
		fmt.Scanln(&input)
		if input[0] >= 65 && input[0] <= 90 {
			name = append(name, input)
		} else {
			break
		}
	}
	fmt.Println(name)
}

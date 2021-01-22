package main

import "fmt"

func main() {
	/*create a slices and store values in it, within range 1 to 5 and ask user to enter index
	position and than ask him to guess value at that location
	User have 2 chances to guess*/
	arr := []int{1, 4, 2, 3}
	var choice, value, try int = 0, 0, 0
	fmt.Println("Enter position")
	fmt.Scanln(&choice)
	for i, v := range arr {
		if i == (choice - 1) {
			for try != 2 {
				fmt.Println("Please enter your guess number :")
				fmt.Scanln(&value)
				if value == v {
					fmt.Println("You are right!")
					break
				} else {
					try++
				}
			}
			if try == 2 {
				fmt.Println("Please try again")
			}
		}
	}
}

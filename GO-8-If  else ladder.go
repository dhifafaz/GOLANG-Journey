package main

import "fmt"

func main() {
	var choice string
	fmt.Println("How would you like to go somewhere ?")
	fmt.Scanln(&choice)
	if choice == "car" {
		fmt.Println("wear your seat belts")
	} else if choice == "bike" {
		fmt.Println("wear your helmet")
	} else if choice == "walking" {
		fmt.Println("cross road carefully")
	}
}

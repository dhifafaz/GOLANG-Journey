package main

import "fmt"

func main() {
	var a string
	fmt.Println("Are my friends coming ?")
	fmt.Scanln(&a)
	if a == "Yes" {
		fmt.Println("Horrey!!")
	} else {
		fmt.Println("Oh nooo")
	}
}

package main

import "fmt"

func main() {
	var a int
	var b string
	fmt.Println("Insert here : ")
	fmt.Scanln(&a)
	fmt.Println("Insert here : ")
	fmt.Scanln(&b)
	fmt.Println(a, b)
}

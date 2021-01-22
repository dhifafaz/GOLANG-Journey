package main

import "fmt"

func main() {
	//break
	for i := 1; i <= 100; i++ {
		if i == 10 {
			break
		}
		fmt.Println(i)
	}
	//continue
	for i := 1; i <= 12; i++ {
		if i == 10 {
			continue
		}
		fmt.Println(i)
	}

}

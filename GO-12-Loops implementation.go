package main

import "fmt"

func main() {
	var ranges int
	fmt.Scanln(&ranges)
	counter3 := ranges
	for ; counter3 >= 0; counter3-- {
		fmt.Println(counter3)
	}
	//print alphabet with loop
	//ASCII A=65 Z=90 a=97 z=122
	// Using print definition
	fmt.Println("\n")
	for i := 65; i <= 90; i++ {
		// using printf (print function) and using %c for char format
		// %d use to print numbers
		// %f use to print float
		// %s use to print string
		fmt.Printf("%c", i)
	}
	fmt.Println("\n")
	for j := 122; j >= 97; j-- {
		// using printf (print function) and using %c for char format
		// %d use to print numbers
		// %f use to print float
		// %s use to print string
		fmt.Printf("%c", j)
	}
	fmt.Println("\n")
	//printing all odd number from 1 to 100
	for g := 1; g <= 100; g++ {
		if g%2 == 0 {
			fmt.Println(g)
		}
	}

	var input, sum int
	fmt.Scanln(&input)
	count := input
	for count >= 1 {
		if count%2 != 0 {
			sum += count
		}
		count--
	}
	fmt.Println(sum)
}

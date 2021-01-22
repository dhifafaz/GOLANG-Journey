package main

import "fmt"

func main() {
	// checking how manu times a value is reapeated in slice using range
	// and also print the location
	a := []int{1, 2, 3, 1, 2, 5, 3}
	var ch []int // to store the number that already checked
	flag := 0
	count := 0

	for i, v := range a {
		for q := 0; q < len(ch); q++ {
			if v == ch[q] {
				flag = 1
			}
		}
		if flag != 1 {
			for j := i + 1; j < len(a); j++ {
				if a[i] == a[j] {
					fmt.Println("Position", j)
					count++
				}
			}
			fmt.Println("Value", v)
			if count != 0 {
				fmt.Println("Times", count, "\n")
				ch = append(ch, a[i])
			}
		}
		flag = 0
		count = 0
	}

}

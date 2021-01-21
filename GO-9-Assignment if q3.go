package main

import "fmt"

func main() {
	var math, science, history, total, percentage float64
	fmt.Scanln(&math, &science, &history)
	total = math + science + history
	percentage = (total / 300) * 100
	if percentage >= 90 {
		fmt.Println("A")
	} else if percentage >= 80 && percentage < 90 {
		fmt.Println("B")
	} else if percentage >= 70 && percentage < 80 {
		fmt.Println("C")
	} else if percentage >= 60 && percentage < 70 {
		fmt.Println("D")
	} else {
		fmt.Println("Fail")
	}

}

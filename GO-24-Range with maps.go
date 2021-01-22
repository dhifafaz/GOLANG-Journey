package main

import "fmt"

func main() {
	//maps
	m := map[string]int{
		"Dhifaf": 1919,
		"rob":    1919,
		"peter":  1919,
	}

	for key, value := range m {
		fmt.Println(key, value)
	}
}

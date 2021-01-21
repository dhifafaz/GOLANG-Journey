package main

import "fmt"

func main() {
	const PI float64 = 3.141592653589793
	//it's look like var PI float64
	//PI = 3.23
	fmt.Println(PI)

	// type casting
	var a int = 10
	var b float64 = 20
	b = float64(a)
	fmt.Println(b)

}

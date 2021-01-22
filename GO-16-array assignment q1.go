// find name in array
package main

import "fmt"

func main() {
	// Array finding name in array
	name := [5]string{"Dhifaf", "adit", "fahmi", "willy", "ojan"}
	var findName string
	fmt.Scanln(&findName)
	var i int = 0
	for ; i < 5; i++ {
		if name[i] == findName {
			fmt.Println("Name found in array")
			break
		}
	}
	if i == 5 {
		fmt.Println("Name not found in array")
	}

}

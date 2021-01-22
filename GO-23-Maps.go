package main

import "fmt"

func main() {
	//maps
	// in maps we dont have index but has key value
	x := map[string]int{
		"Dhifaf": 119140047,
		"Adit":   1191140052,
		"Arip":   119140149,
		//remember to always put the coma
	}
	fmt.Println(x["Dhifaf"], x["Adit"], x["Arip"])

	// another way to write map
	y := make(map[string]int)
	// to assign value to the map
	y["Dhifaf"] = 119140047
	fmt.Println(y["Dhifaf"])

	//maps will return to value
	z := make(map[string]int)
	z["Dhifaf"] = 119140047
	v, t := z["Dhifaf"]
	// it will return the value and true or false
	fmt.Println(v, t)

	// to delete any valu in maps
	// using delete(name of the maps, key value)
	delete(z, "Dhifaf")
	fmt.Println(z["Dhifaf"])
}

package main

import (
	"fmt"
)

func main() {
	//Aritmathic
	var a int = 10
	b := 20
	c := a + b
	fmt.Println(c)
	fmt.Println(b - a)
	fmt.Println(b * a)
	fmt.Println(b / a)
	fmt.Println(b % a)
	var d int = 1
	d++
	fmt.Println(d)
	d--
	fmt.Println(d)

	//Relational Operators
	fmt.Println('H' == 'H')
	fmt.Println('H' == 'z')
	fmt.Println(2 != 1)
	fmt.Println('H' != 'H')
	fmt.Println(3 > 2) //the same with greater equal to
	fmt.Println(1 < 2) //the same with smaller equal to

	//Logical Operators
	//&& and
	fmt.Println(1 == 1 && 2 == 2)
	//|| or
	fmt.Println(1 == 1 || 2 == 2)
	// ! not
	fmt.Println(!(1 == 1))

	//Assignmetn Operators
	//= assign value right to the left
	var a1 int = 110
	var b1 int = 20
	a1 = b1
	fmt.Println(a1)
	//+= variable = variable + x
	a1 += 5
	fmt.Println(a1)
	//-= variable = variable - x
	a1 -= 5
	fmt.Println(a1)
	// also ther are /= and *=, %=
	a1 %= 5
	fmt.Println(a1)
}

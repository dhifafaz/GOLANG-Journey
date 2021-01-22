package main

import "fmt"

func main() {
	//create a simple contact directory and user can search the number of whos name is typed in by user
	contactApp := make(map[string]int)
	var name string
	var phoneNumber, n int
	var counter int = 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&name)
		fmt.Scanf("%d", &phoneNumber)
		contactApp[name] = phoneNumber
	}
	//fmt.Println(contactApp)
	var choice int
	fmt.Println("Choose the way you want to search a contact : 1. By Name 2. By Number")
	fmt.Scanln(&choice)
	if choice == 1 {
		fmt.Scanln(&name)
		for key, value := range contactApp {
			if key == name {
				fmt.Println(key, value)
				counter = 1
				break
			}
		}
	}
	if choice == 2 {
		fmt.Scanln(&phoneNumber)
		for key, value := range contactApp {
			if value == phoneNumber {
				fmt.Println(key, value)
				counter = 1
				break
			}
		}
	}
	if counter == 0 {
		fmt.Println("There is no such a name in your contact")
	}

}

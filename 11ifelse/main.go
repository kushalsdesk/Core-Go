package main

import (
	"fmt"
)

func main() {
	fmt.Println("If Else in Golang")
	var result string

	loginCount := 10

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Non Regular user"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	// on the go checking
	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// structure like handling web request

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Numver is not less than 10")
	}
}

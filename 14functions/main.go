package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to fucntions in golang")
	greeter()

	// result := adder(3, 5)
	result := proAdder(2, 2, 2)
	fmt.Println("Result is: ", result)

	anoRes, anoString := twoTypes(5, 6, 7, 8) // also need to get them in different values

	fmt.Println(anoRes, anoString)
}

// custom function
func greeter() {
	fmt.Println("Hello from Golang")
}

// custom function with parameter
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// standard syntax for handling common values
func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

// returning multiple types
func twoTypes(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "This is the another return type"
}

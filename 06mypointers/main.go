package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to a class on pointers")

	//declaring a pointer that holds an integer
	var ptr *int

	fmt.Println("Value of pointer is, " , ptr)

	myNumber := 23

	//creating a pointer to point myNumber by using "&" to include the reference

	var ptr2 = &myNumber

	fmt.Println("The value of the reference pointer is, ", ptr2)
	// "*" is used to check the value inside
	fmt.Println("The value of the reference pointer is, ", *ptr2)

	*ptr2 = *ptr2 * 2 
	fmt.Println("New value is : " , myNumber)
}

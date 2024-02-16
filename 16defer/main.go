package main

import (
	"fmt"
)

func main() {
	// Defer printing "World" after all other statements in the surrounding function (main) execute
	defer fmt.Println("World")

	// Loop 1: Print numbers 1 to 5, with defered messages at start and end
	defer fmt.Println("Loop 1 starting...")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	defer fmt.Println("Loop 1 finished.")

	// Loop 2: Print letters A to E, with nested defer for a message in the middle
	defer fmt.Println("Loop 2 starting...")
	for letter := 'A'; letter <= 'E'; letter++ {
		defer fmt.Println("Processing letter", string(letter)) // Nested defer will execute before outer loop's defer
		fmt.Println(string(letter))
	}
	defer fmt.Println("Loop 2 finished.")

	looper := 1
	for looper < 5 {
		fmt.Print(looper)
		looper++
	}
}

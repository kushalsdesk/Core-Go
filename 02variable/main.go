package main

import (
	"fmt"
)

// To make any variable public. Go doesnot use a keyword "public", it gets it through the variable name as it will have the first char in Uppercase

var Language string = "Go programming language" // public variable

func main() {
	var username string = "kushal"

	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var intValue int = 23444
	fmt.Println("The integer is :", intValue)
	fmt.Println("int can be used in natural work, OS specific work may need more specific  type..")

	var floatValue float32 = 245.5657657
	var floatValue2 float64 = 245.5657657

	fmt.Println("Now there is the difference between float32 and float64 is that first one will be approximate.. second one will be accurate ", floatValue, floatValue2)

	// implicit way to declare & initialize

	// We can also declare variables implicitly without giving any type...
	// But for this the variable must be initiated at the same time so that lexer can know which type to include it with the var name
	// Modern Go syntax also use a declaration named "no var" where it does not use "var" keyword but this is only allowed in any method scope or local scope....not in global scope, that remains like the above declaration

	name := "Kushal" // here as i have given the var a value of string, lexer will set the type as string as well
	fmt.Printf("The name value is of type %T ", name)

	// Declaring Constant, one-time, can not be modified later

	const loginToken string = "@lastEnteredPass"
}

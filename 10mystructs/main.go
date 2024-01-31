package main

import (
	"fmt"
)

// Thank go we don't have any inheritence (so, no super,class,parent,child)
func main() {
	fmt.Println("Welcome to Structures in Golang")

	kushal := User{"Kushal", "kushal@go.dev", true, 24}
	fmt.Println("The new user is: ", kushal)
	// more detail with placeholder
	fmt.Printf("Details are : %+v\n", kushal)
	fmt.Printf("Name is %v and email  is %v", kushal.Name, kushal.Email)
}

// creating a structure
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	// Syntax for Declaring and initiating the slice at one

	fruitList := []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// To add elements, we use "append" built-in function followed by the slice name and the elements

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// to slice more an existing slice we can have this syntax
	fruitList = append(fruitList[1:]) // this will leave the first element or start the slice from the first index to end
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // This is give it a range in -between{index 1 to index 2}
	fmt.Println(fruitList)

	fruitList = append(fruitList[:5]) // this gets the element from first to 3rd index
	fmt.Println(fruitList)

	// Memory management 1.with "make" keyword, 2.with "new" keyword
	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 465
	highscores[3] = 777

	// highscores[0] = 234 //this will not work as we set the size of 4

	// but when we append new value go actually re-allocates the memory for the data structure, it is fast and memory safe
	highscores = append(highscores, 555, 666, 321) // this will surely add the latest values with the slice

	fmt.Println(highscores)

	// slices also have the "sort" functionality that array lacks

	fmt.Println("Is the array sorted: ", sort.IntsAreSorted(highscores))
	sort.Ints(highscores) // sorts the structrure's value with the integer order
	fmt.Println(highscores)
	fmt.Println("Is the array sorted: ", sort.IntsAreSorted(highscores))

	// remove a value from slices based on index

	courses := []string{"reactjs", "typescript", "java", "ruby", "golang"}
	fmt.Println("All the courses are", courses)
	index := 2
	// the syntax is pretty weird, the last arguement is getting injected in a explicit manner with a "tripple dot" to modify the fucntion defeinition to accept more arguements than the function is  actually designed
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("The New courses are ", courses)
}

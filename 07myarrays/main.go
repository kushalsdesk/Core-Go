package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitlist [4]string // this is how arrays are declared

	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[3] = "Peach"

	fmt.Println("Fruit list is : ", fruitlist)
	fmt.Println("Fruit list is: ", len(fruitlist))
	// Array declaration and initializing at same time
	// Arrays will always show the declared size as the length, no matter how less values it has...
	veglist := [3]string{"potato", "beans", "Mushroom"}

	fmt.Println("Veg list is : ", veglist)
	fmt.Println("Veg list is : ", len(veglist))
}

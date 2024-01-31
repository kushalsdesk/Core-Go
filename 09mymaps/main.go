package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["JV"] = "Java"

	fmt.Println("List of the languages are: ", languages)
	fmt.Println("JS is: ", languages["JS"])

	// Deleting a value
	delete(languages, "RB")

	fmt.Println("List of the languages are: ", languages)

	// Iterate over a map

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// if dont need any of key or value
	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}
}

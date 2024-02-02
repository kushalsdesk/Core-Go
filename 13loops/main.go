package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for loop

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// shorter syntax

	for i := range days {
		fmt.Println(days[i])
	}

	// for with comma-ok syntax

	for _, days := range days {
		fmt.Printf("The valus is %v\n", days)
	}

	// while syntax

	target, looper := 2, 1
	for looper < 10 {
		if looper == target {
			looper++
			continue // skips the target value
		}
		fmt.Println("The Number is : ", looper)
		looper++
	}
}

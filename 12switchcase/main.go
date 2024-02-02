package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("SwitchCase in golang")

	rand.Seed(time.Now().UnixNano())

	diceNum := rand.Intn(6) // generate number between 0 and 5
	fmt.Println("Value of dice is ", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Dice value is 1, Game started")
	case 2:
		fmt.Println("Dice Value is 2 , move two stops")
	case 3:
		fmt.Println("Dice Value is 3 , move three stops")
	case 4:
		fmt.Println("Dice Value is 4 , move four stops")
	case 5:
		fmt.Println("Dice Value is 5 , move five stops")
	case 6:
		fmt.Println("Dice Value is 6 , move six stops")
	}
}

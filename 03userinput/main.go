package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name:")

	input, _ := reader.ReadString('\n')

	fmt.Println(input)

	// with Scanln
	var name string

	fmt.Println("Enter your username ")
	fmt.Scanln(&name)
	fmt.Println("Your username is ", name)

	// for single integer
	var age int
	fmt.Println("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Println("Your age is", age)

	// taking inputs with more Formatting & Typecasted
	var num1, num2 int
	fmt.Print("Enter two numbers separated by space: ")
	fmt.Scanf("%d %d", &num1, &num2)
	fmt.Println("Sum:", num1+num2)
}

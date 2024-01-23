package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for out Pizza:")

	// comma ok syntax || comma err  syntax(like try-catch block)
	// to check the if recieving any error while handling packages
	// the first var before comma is for the result, and the after part is the error,
	// it can be name anything, but we leave the var as underscore if that var is not needed

	input, _ := reader.ReadString('\n') // as we are taking the number as string, so we can only have string operations over it, need to learn typecasting

	fmt.Println("Thanks for you rating, ", input)
	fmt.Printf("Type of the rating is %T", input)
}

package main

import (
	"fmt"
)

/*
*  Format
* 	type TypeName struct {
    // ... fields ...
}
		func (receiver TypeName) methodName(parameters...) returnType {
    		// ... method body ...
		}
*/

// to make any block exportable it need to have first Letter in Capital
type Point struct {
	X int
	Y int
}

// getter with no return type
func (pArg Point) getPoints() {
	fmt.Printf("The current Point X is %v and Point Y is %v", pArg.X, pArg.Y)
}

// setter method (always need to Pointer Receiver when using setter)
// as Pointer will actually modify the value with the memory reference
func (pArg *Point) setPoints(dx int, dy int) {
	pArg.X = dx
	pArg.Y = dy
}

func main() {
	fmt.Println("Welcome to Methods in Go")

	var pointX, pointY int
	fmt.Println("Enter the X Point value")
	fmt.Scanln(&pointX)
	fmt.Println("Enter the Y Point value")
	fmt.Scanln(&pointY)

	userPoint := Point{}
	userPoint.setPoints(pointX, pointY)
	userPoint.getPoints()
}

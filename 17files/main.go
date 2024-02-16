package main

import (
	"fmt"
	"io"
	"os"

	check "files.in/utils/service"
)

func main() {
	content := "#This is will be in the Content"
	file, err := os.Create("./content.md")
	check.IsErr(err)

	length, err := io.WriteString(file, content)
	check.IsErr(err)

	fmt.Println("the length of the file is", length)

	defer file.Close()

	// read the file

	readFile("./content.md")
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	check.IsErr(err)
	fmt.Println("The Content of the File is: \n ", string(data))
}

package main

import (
	"fmt"
	"io"
	"net/http"

	check "files.in/utils/service"
)

// targetting a local php page, hosting in local php server

func main() {
	const url string = "http://localhost:8080" // local index.php to push request

	response, err := http.Get(url)
	check.IsErr(err)
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	check.IsErr(err)
	fmt.Println("This is the Internal of the WebPage \n", string(data))
}

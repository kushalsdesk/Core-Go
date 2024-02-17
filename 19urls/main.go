package main

import (
	"fmt"
	"net/url"

	check "files.in/utils/service"
)

const myurl string = "http://localhost:8080/learn?coursename=php&project=laravel"

func main() {
	fmt.Println("This is the URL handling")
	fmt.Println(myurl)

	// parsing the url

	result, err := url.Parse(myurl)

	check.IsErr(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// better format for getting the queries

	qparams := result.Query()

	fmt.Printf("The Type of the param is %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Params is: ", val)
	}

	// now lets construct an URL
	// by passing the reference of the url
	partsOfURL := &url.URL{
		Scheme:  "http",
		Host:    "localhost",
		Path:    "/learn",
		RawPath: "user=kushal",
	}

	anotherURL := partsOfURL.String()
	fmt.Println("The new Constructued URL is", anotherURL)
}

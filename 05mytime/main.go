package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study ")

	presentTime := time.Now()
	fmt.Println(presentTime)

	/* The default format might not be useful for all
	 but to get the understanble version of the current time there is a syntax, that may seem overwhelming where we put the format by passing a proper date as an arguement
	 need to pass only the needed value, like pass the date just to get the date, or pass time just to get time, but lets look at the key-value pair
			{
				date = ["01-02-2006", "01/02/2006"],
				time = ["3:04:05pm", "15:04:05"],
				day = "Monday"
			},
			arrays contains inner-formats,
	*/

	fmt.Println(presentTime.Format("01-02-2006 3:04:05pm Monday"))

	// creating the timestamps from any of the values

	createDate := time.Date(2024, time.July, 28, 22, 30, 0, 0, time.UTC)

	fmt.Println(createDate)
	//again need to format the data created
	fmt.Println(createDate.Format("01/02/2006 Monday"))
}

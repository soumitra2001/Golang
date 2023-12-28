package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang")

	currentDate := time.Now()
	fmt.Println("Current Date: ", currentDate.Format("01-02-2006 15:04:05 Monday")) // this the standard layout to format

	createdDate := time.Date(2024, time.May, 15, 11, 0, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

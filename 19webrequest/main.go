package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"

func main() {
	fmt.Println("My Web Request")

	responce, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer responce.Body.Close()

	fmt.Printf("Type of responce: %T\n", responce)
	dataByte, err := io.ReadAll(responce.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataByte)
	fmt.Println(content)
}

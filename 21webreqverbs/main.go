package main

import (
	"fmt"
	"io"
	"net/http"
	// "strings"
)

func main() {
	fmt.Println("Welcome to Web Request in golang")

	PerformGetRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:3000/get"

	responce, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()
	fmt.Println(responce.StatusCode)
	// fmt.Println(responce.ContentLength)

	// var responceString strings.Builder
	content, _ := io.ReadAll(responce.Body)
	// byteCount, _ := responceString.Write(content)
	// fmt.Println(byteCount)
	// fmt.Println(responceString.String())
	fmt.Println(string(content))
}

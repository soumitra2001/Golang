package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Request in golang")

	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostformRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:3000/get"

	fmt.Println(url)
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

func PerformPostRequest() {
	const myurl = "http://localhost:3000/post"

	data := strings.NewReader(`
		{
			"Name":"Supriya",
			"email":"sg@go.dev"
		}
	`)

	responce, err := http.Post(myurl, "application/json", data)

	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()
	content, _ := io.ReadAll(responce.Body)

	fmt.Println(string(content))
}

func PerformPostformRequest() {
	const myurl = "http://localhost:3000/postform"

	formData := url.Values{}

	formData.Add("Country", "India")
	formData.Add("Name", "SG")

	responce, err := http.PostForm(myurl, formData)
	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	content, _ := io.ReadAll(responce.Body)

	fmt.Println(string(content))
}

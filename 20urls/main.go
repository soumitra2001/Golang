package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://linkedin.com:3400/in/itsmesupriya?filter=all"

func main() {
	fmt.Println("Welcome to URLs in Golang")
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The Type of query params are: %T\n", qparams)

	// fmt.Println(qparams["filter"])

	for key, val := range qparams {
		fmt.Printf("Key: %v and Value: %v\n", key, val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "http",
		Host:     "github.com",
		Path:     "soumitra01",
		RawQuery: "tab=repositories",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}

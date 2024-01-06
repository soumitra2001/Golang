package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json in Go")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{Name: "ReactJs", Price: 299, Platform: "LearnCodeOnline.in", Password: "abc323", Tags: []string{"React", "Js"}},
		{Name: "Golang", Price: 999, Platform: "YouTube.com", Password: "kuchbhi", Tags: nil},
		{Name: "Javascript", Price: 599, Platform: "LearnCodeOnline.in", Password: "jbk726", Tags: []string{"Functional", "Dynamic", "Web", "Js"}},
	}

	jsonData, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)
}

func DecodeJson() {
	jsonData := []byte(
		`{
			"coursename": "Javascript",
			"Price": 599,
			"website": "LearnCodeOnline.in",
			"tags": [
					"Functional",
					"Dynamic",
					"Web",
					"Js"
			]
		}`,
	)

	isValidJson := json.Valid(jsonData)

	var courses course

	if isValidJson {
		fmt.Println("This a valid json data")
		json.Unmarshal(jsonData, &courses)
		fmt.Printf("Value of courses is: %#v\n", courses)
	} else {
		fmt.Println("Invalis json data")
	}

	var courseData map[string]interface{}

	json.Unmarshal(jsonData, &courseData)

	// fmt.Printf("%#v\n", courseData)

	for k, v := range courseData {
		fmt.Printf("Key is %v and Value is %v and Type %T\n", k, v, v)
	}
}

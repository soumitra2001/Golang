package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Json in Go")
	EncodeJson()
}

func EncodeJson() {
	myCourses := []Course{
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

type Course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

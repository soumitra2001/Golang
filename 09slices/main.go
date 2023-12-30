package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Golang")

	var nameList = []string{"Manju", "Akash"}

	nameList[0] = "Jit"
	nameList = append(nameList, "Supriya", "Sanju")
	fmt.Println("Values of nameList is: ", nameList)
	fmt.Printf("Type of nameList: %T \n", nameList)

	nameList = append(nameList[0:3])
	fmt.Println(nameList)

	scores := make([]int, 3)

	scores[0] = 87
	scores[1] = 28
	scores[2] = 96
	// scores[3] = 48

	scores = append(scores, 273, 97, 67)
	fmt.Println(scores)

	sort.Ints(scores)
	fmt.Println(scores)

	fmt.Println("Is scores are sorted: ", sort.IntsAreSorted(scores))

	// How to remove a value from Slice based on index

	languages := []string{"Java", "Go", "Pyhon"}

	languages = append(languages, "C++", "Javascript")

	fmt.Println("Values of languages: ", languages)

	var index = 3
	languages = append(languages[:index], languages[index+1:]...)

	fmt.Println("Values of languages: ", languages)

}

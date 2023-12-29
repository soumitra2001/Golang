package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in Golang")

	var langList [4]string

	langList[0] = "Bengali"
	langList[1] = "English"
	langList[3] = "Spanish"

	fmt.Println("Language list is: ", langList)
	fmt.Println("Size of langList is: ", len(langList))
	fmt.Println("langList 2nd index value: ", langList[2])

	fruitList := [3]string{"Apple", "Orange", "Banana"}

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Size of fruitList is: ", len(fruitList))
	fmt.Printf("FruitList type is  %T", fruitList)
}

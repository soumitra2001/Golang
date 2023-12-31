package main

import "fmt"

func main() {

	fmt.Println("switch case in Golang")

	var num int = 2

	switch num {
	case 1:
		fmt.Println("value is: ", num)

	case 2:
		fmt.Println("Value is: ", num)
		fallthrough
	case 3:
		fmt.Println("Value is: ", num)
	default:
		fmt.Println("value is: ", num)
	}
}

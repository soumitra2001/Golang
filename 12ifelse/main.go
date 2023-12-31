package main

import "fmt"

func main() {
	userCount := 23

	if userCount < 10 {
		fmt.Println("User count less than 10")
	} else if userCount > 10 {
		fmt.Println("Usre count greater than 10")
	} else {
		fmt.Println("User count exact 10")
	}
}

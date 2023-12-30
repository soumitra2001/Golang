package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Map in Golang")

	langs := make(map[string]string)

	langs["JS"] = "Javascript"
	langs["PY"] = "Python"
	langs["RB"] = "Ruby"

	fmt.Println("Values of langs is: ", langs)
	fmt.Println("JS shorts for ", langs["JS"])

	delete(langs, "RB")
	fmt.Println("Value of langs is: ", langs)

	for key, value := range langs {
		fmt.Printf("For key %v value is %v \n", key, value)
	}

}

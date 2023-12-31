package main

import "fmt"

func main() {
	fmt.Println("Welcome to function in Go")

	Greeter()

	result := adder(3, 6)

	fmt.Println("Result is: ", result)

	total, _ := proAdder(3, 8, 5, 7, 26)

	fmt.Println("Total value is: ", total)

}

func Greeter() {
	fmt.Println("Happy new year 2024")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	sum := 0

	for i := range values {
		sum += values[i]
	}

	return sum, "Numbers added sucessfully!"
}

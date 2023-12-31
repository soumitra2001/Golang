package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thusday", "Friday"}

	fmt.Println("Days of week: ", days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for idx, value := range days {
	// 	fmt.Printf("Value of index %v is: %v \n", idx, value)
	// }

	value := 1

	// for i := 1; i <= value; i++ {
	// 	fmt.Print(i, " ")
	// }

	for value <= 10 {
		if value == 5 {
			value++
			continue
		}

		if value == 10 {
			goto sup
		}

		fmt.Println(value, " ")
		value++
	}

sup:
	fmt.Println("Hello this is the last line")

}

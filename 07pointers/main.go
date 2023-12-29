package main

import "fmt"

func main() {
	fmt.Println("Knowing all about pointers =>")

	var num *int // to create a new pointer we use * sign

	fmt.Println("Value of num ", num)

	myNum := 20

	fmt.Println("Value of myNum ", myNum)

	ptr := &myNum // to assign other val to pointer we use & sign

	fmt.Println("Memory address of ptr ", ptr)
	fmt.Println("Value of ptr ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("Noe value of myNum ", myNum)
}

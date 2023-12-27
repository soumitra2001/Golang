package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var greet string = "Welcome"
	fmt.Println(greet)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your name: ")

	// read input using comma ok || comma error syntax
	name, _ := reader.ReadString('\n')

	fmt.Print("Hello " + name)
	fmt.Printf("Type of input variable: %T", name)
}

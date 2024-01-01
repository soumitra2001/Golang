package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to file in Golang")
	content := "Happy New Year to all, I wish this year brings you happiness."

	file, err := os.Create("mywish.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Content successfully add to file. \nLength= ", length)

	defer file.Close()

	readFile("mywish.txt")

}

func checkNilErr(err error) {

	if err != nil {
		panic(err)
	}
}

func readFile(path string) {
	dataBytes, err := os.ReadFile(path)

	checkNilErr(err)

	fmt.Println("Data in the file is: \n", string(dataBytes))
}

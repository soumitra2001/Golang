package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	msg := "Hello"

	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)

	msg = "Goodbye"
	wg.Wait()

	// go f1()
	// go f2()

}

// func f1() {
// 	x := 3
// 	x++
// 	fmt.Printf("Value of x is :%v \n", x)
// }

// func f2() {
// 	y := 5
// 	y++
// 	fmt.Printf("Value of y is :%v \n", y)
// }

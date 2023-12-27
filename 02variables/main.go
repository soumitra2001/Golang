package main

import "fmt"

const LoginToken string = "uuclfhdcd" //[Public] In Go we declare a variable public by writing variable first letter in capital

func main() {
	var username string = "supriya01"
	fmt.Println(username)
	fmt.Printf("Variable is type of: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is type of: %T \n", isLoggedIn)

	var smallVal uint8 = 255 // range from 0 to 255 other wise in general we can use int type
	fmt.Println(smallVal)
	fmt.Printf("Variable is type of: %T \n", smallVal)

	var floatVal float64 = 6588.764
	fmt.Println(floatVal)
	fmt.Printf("Variable is type of: %T \n", floatVal)

	// Different type of variable declaration
	var name = "Supriya" // now Lexer set this type as string so can not assign or change it ot other type
	fmt.Println(name)

	// declare variable using walrus operator
	userEmail := "sg@gmail.com" // we can use any type
	fmt.Println(userEmail)
	// we can use walrus operator inside the method only

	fmt.Println(LoginToken)
	fmt.Printf("Variable of type: %T \n", LoginToken)
}

package main

import "fmt"

func main() {
	fmt.Println("Struct in Golang")

	myUser := User{"Supriya", "sg@gmail.com", false, 23}

	fmt.Printf("MyUser value %+v \n", myUser) // +v is used to get detail value

	fmt.Printf("Name is %v and Email is %v \n", myUser.Name, myUser.Email)

	myUser.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool

	Age int
}

func (u User) GetStatus() {
	fmt.Println("Is active user:", u.Status)
}

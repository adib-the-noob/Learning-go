package main

import "fmt"

func main(){
	fmt.Println("Structs in GoLang")
	adib := User{
		FirstName: "Adib",
		LastName: "the Noob",
		Age: 20,
	}
	fmt.Printf("FirstName is %v and last Name is %v", adib.FirstName, adib.LastName)
}

type User struct {
	FirstName string
	LastName string
	Age int
}

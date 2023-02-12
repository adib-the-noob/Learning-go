package main

import "fmt"

func main() {
	var name string
	fmt.Println("What's your name?")
	fmt.Scan(&name)
	age := 12
	fmt.Println(name)
	fmt.Printf("Age is %v and it's type is %T", age)
}

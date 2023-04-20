package main

import "fmt"

func main() {
	myNumber := 42
	var myPointer = &myNumber
	fmt.Println(myPointer)
	fmt.Println(*myPointer)
	newResult := *myPointer + 10
	fmt.Println(newResult)
}
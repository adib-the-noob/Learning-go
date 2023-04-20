package main

import "fmt"

func addTwoNumber(number1 int, number2 int) (result int) {
	result = number1 + number2
	return
}

func main() {
	fmt.Println(addTwoNumber(1, 4))
}

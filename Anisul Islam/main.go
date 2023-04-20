package main

import "fmt"

const PI = 3.1416

func main() {
	var a = 12
	var b = "adib"
	var cars = [3]string{"adib", "the_noob", "pythonista"}
	price := [2]string{"mohammed", "adib"}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(PI)
	fmt.Println(cars)
	fmt.Print(price)
	price[1] = "mohammed_updated"
	fmt.Println(price[1])
}

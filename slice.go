package main

import "fmt"

func main() {
	myslice := []int{}
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	append()
	myslice[1] = 12
	myslice[2] = 23
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
}

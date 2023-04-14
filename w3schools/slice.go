package main
import "fmt"

func main()  {
	myslice := []string{"GO", "Slices", "are", "powerful"}
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	myNewSlice := []string {"Here", "How", "it", "is"}
	result := append(myslice, myNewSlice...)
	fmt.Println(result)
}
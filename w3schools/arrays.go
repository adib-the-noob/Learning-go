package main
import "fmt"

func main()  {
	// var fruite_list = [4]string{"Apple", "Orange", "Grape", "Mango"}
	// fmt.Printf(fruite_list[1])
	
	// Array for Integer
	// NumberList := []int{
	// 	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	// }
	// fmt.Println(NumberList)

	// array without initial value
	// withOutValueList := [5]int{}
	// fmt.Println(withOutValueList)

	// withSomeValueList := [5]int{1, 2, 3}
	// fmt.Println(withSomeValueList)
	
	// With some specific value
	array1 := [5]int{1: 10, 3: 30}
	fmt.Println(array1)
	fmt.Println(len(array1))
	
}
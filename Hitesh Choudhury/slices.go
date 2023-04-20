package main 
import "fmt"

func main() {
	fruite_list := []string{"Apple", "Orange", "Grape", "Mango"}
	fmt.Println(fruite_list)
	fmt.Println(len(fruite_list))
	fmt.Println(fruite_list[1:3])
	fmt.Println(fruite_list[1:])
	fmt.Println(fruite_list[:3])
	fmt.Println(fruite_list[1:])

	// Append
	fruite_list = append(fruite_list, "Pineapple")
	fmt.Println("New list ", fruite_list)
}
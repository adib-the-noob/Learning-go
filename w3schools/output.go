package main
import "fmt"


func main() {
	var i, j string = "Hello", "World"
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")
	fmt.Printf("The value of '%v' and type is %T \n", i , i)
	fmt.Printf("Go Syntax format %#v", i)
	fmt.Printf("The value of '%v' and type is %T ", j , j)
}
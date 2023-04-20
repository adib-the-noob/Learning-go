package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println("------------------")
	for day := range days{
		fmt.Println(days[day])
	}
	fmt.Println("------------------")

	for index, day := range days{
		fmt.Println(index, day)
	}

	fmt.Println("------------------")

	testValur := 0
	for testValur < 10 {
		if testValur == 5 {
			goto middleNumber
			
		}
		fmt.Println(testValur)
		testValur++
	}
middleNumber:
	fmt.Println("Middle Number")
}
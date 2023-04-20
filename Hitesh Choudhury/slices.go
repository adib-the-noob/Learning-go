package main

import (
	"fmt"
	"sort"
)

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

	// sort 
	highScores := []int{500, 200, 400, 300}
	fmt.Println(highScores)
	sort.Sort(sort.Reverse(sort.IntSlice(highScores)))
	fmt.Println(highScores)

	courses := []string{"Docker", "Kubernetes", "Puppet", "Terraform", "AWS", "GCP", "Azure"}
	fmt.Println(courses)
	index := 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
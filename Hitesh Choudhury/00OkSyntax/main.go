package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This contains similer like - Try catch block
func main() {
	fmt.Println("Please rate our pizza!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate our pizza from 1 to 5")

	input , _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating our pizza with", input)
	
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Thanks for rating our pizza with", numRating + 1)
	}
}
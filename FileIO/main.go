package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fileContent := "This is a test file"
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, fileContent)
	checkNilErr(err)
	fmt.Printf("Wrote %d bytes to file", length)	
	file.Close()
	readFile("test.txt")
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside file is ", string(databyte))
}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}
package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	formatedTime := presentTime.Format("01-02-2006 Monday 15:04:05")
	fmt.Println(formatedTime)

}
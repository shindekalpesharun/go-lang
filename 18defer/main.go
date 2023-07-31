package main

import (
	"fmt"
)

func main()  {
	// This put on the last on the function but its work like stack 
	defer fmt.Println("This with defer keyword")
	defer fmt.Println("This with defer keyword 2")
	fmt.Println("hi this is method")
	myDefer()
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

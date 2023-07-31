package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Welcome to a class on pointers");

	// var ptr *int
	// fmt.Println("Values of pointer is ", &ptr)

	// normal variable
	myNumber := 23
	// address store in ptr or reference
	var ptr = &myNumber
	fmt.Println("Values of actual pointer is ", ptr)
	fmt.Println("Values of actual pointer is ", *ptr)


	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)
}
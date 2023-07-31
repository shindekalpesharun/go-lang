package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Welcome to function in golang")

	// function in function is not allow
	// func greeterTwo(){
	// 	fmt.Println("Another Method")
	// }

	greeter()
	result := adder(3, 5)
	fmt.Println("result is: ", result)

	proResult, text := proAdder(2,3,8,7)
	fmt.Println("Pro result is: ",proResult)
	fmt.Println(text)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values{
		total += val
	}

	return total, "hi this is function"
}

func greeter(){
	fmt.Println("Namstey from golang")
}
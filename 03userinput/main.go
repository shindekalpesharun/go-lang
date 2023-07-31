package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	// comma okay syntax, comma error syntax
	// if don't know what type of data come then use
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for out Pizza:")

	// comma ok || error ok
	input, _:= reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
}
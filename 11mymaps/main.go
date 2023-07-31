package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Maps in golang")

	lang := make(map[string]string)

	lang["js"] = "Javascript"
	lang["rb"] = "Ruby"
	lang["py"] = "Python"

	fmt.Println("List of all lang: ", lang)
	fmt.Println("JS short for: ", lang["js"])

	// loops are interesting in golang

	for key, value := range lang{
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
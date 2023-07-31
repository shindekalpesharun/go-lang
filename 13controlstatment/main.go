package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Welcome to control statement")

	loginCount := 23
	var result string
	
	if loginCount < 10{
		result = "Regular user"
	}else if loginCount > 10{
		result = "watchout"
	}else{
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0{
		fmt.Println("Number is even")
	}else{
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10{
		fmt.Println("num is less than 10")
	}else{
		fmt.Println("num is not less than 10")
	}

	// if err != nil{
		// 
	// }
}
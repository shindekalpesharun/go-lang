package main

import (
	"fmt"
)

func main()  {
	fmt.Println("welcome to struct")

	// no inheritance in golang: No super or parent
	kalpesh:=User{"Kalpesh", "test@test.com", true, 22}
	fmt.Println(kalpesh)
	fmt.Printf("kalpesh details are: %+v\n", kalpesh)
	fmt.Printf("name is: %v and email is %v:\n", kalpesh.Name, kalpesh.Email)
}

// Uppercase in name convention means there are public variable
type User struct{
	Name string
	Email string
	Status bool
	Age int
}
package main

import "fmt"

// volus operator is not allowed
// jwtToken := 3000

// allowed
var jwtToken int = 30000

// constant
const LoginToken string = "mujnyhbtgvrfcedxs"
// capital letter of first word means is public variable

func main()  {
	// fmt.Println("Variables")
	var username string = "kalpesh"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	
	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("variable is of type: %T \n", smallValue)

	var smallFloat float32 = 255.54565454456789076549876544567897654
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	var bigFloat float64 = 255.54565454456789076549876544567897654
	fmt.Println(bigFloat)
	fmt.Printf("variable is of type: %T \n", bigFloat)
	
	// default values and some aliases
	var anotherVariable = "kalpesh"
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// garbage value 
	var anotherVariableInt int
	fmt.Println(anotherVariableInt)
	fmt.Printf("Variable is of type: %T \n", anotherVariableInt)

	// implicit type lexer
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)
}
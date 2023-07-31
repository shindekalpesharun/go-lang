package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our Pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _:= reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil{
		fmt.Println(err)
		// panic(err)
	}else{
		fmt.Println("Added 1 to your rating: ", numRating + 1)
	}
}
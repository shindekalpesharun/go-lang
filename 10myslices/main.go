package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to video on Slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	// delete
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 526
	highScores[2] = 887
	highScores[3] = 912
	// highScores[4] = 999

	highScores = append(highScores, 555, 666, 777)

	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove value from slices based on index
	var courses = []string{"react js", "js", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
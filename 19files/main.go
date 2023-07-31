package main

import(
	"fmt"
	"os"
	"io"
	"io/ioutil"
)

func main()  {
	fmt.Println("Welcome to files in golang")

	content := "This need to go in the files - shindekalpesharun"

	file, err := os.Create("./myFile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(fileName string){
	databyte, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("Text data inside file is: \n", databyte)
	fmt.Println("Text data inside file is: \n", string(databyte))
}

func checkNilError(err error)  {
	if err != nil{
		panic(err)
	}
}
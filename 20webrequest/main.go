package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const url = "https://lco.dev"

func main()  {
	fmt.Println("LCO web request")

	res, err := http.Get(url)
	
	if err != nil{
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", res)
	defer res.Body.Close()	// caller's responsibilty to close the connection

	databyte, err := ioutil.ReadAll(res.Body)
	if err != nil{
		panic(err)
	} 

	content := string(databyte)
	fmt.Println(content)
}
package main 

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"net/url"
)

func main(){
	fmt.Println("Welcome to web verb video")
	// performGetRequest()
	performPostJsonRequest()
}

func performGetRequest()  {
	const baseUrl = "http://localhost:1111/get"
	res, err := http.Get(baseUrl)
	if err != nil{
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("Status code: ", res.StatusCode)
	fmt.Println("Content Length is: ", res.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(res.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}

func performPostJsonRequest(){
	const baseUrl = "http://localhost:1111/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "Let's go with golang",
			"price" : 0,
			"platform" : "learncodeonline.in"
		}	
	`)

	res, err := http.Post(baseUrl, "application/json", requestBody)
	if err != nil{
		panic(err)
	}
	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}

func performPostFormRequest()  {
	const baseUrl = "http://localhost:1111/postform"

	// formdata

	data := url.Values{}
	data.Add("firstName", "kalpesh")
	data.Add("lastName", "shinde")
	data.Add("email", "shindekalpesharun@gmail.com")

	res, err := http.PostForm(baseUrl, data)
	if err != nil{
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
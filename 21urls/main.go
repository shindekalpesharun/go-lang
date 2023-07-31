package main

import(
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gsfsdf745fg"

func main()  {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myUrl)

	// parsing
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query param are: %T\n", qparams)
	fmt.Println("The type of query param are: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for val, _ := range qparams {
		fmt.Println("Param is: ", val)
	}

	partOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=hitesh",
	}

	anotherUrl := partOfUrl.String()
	fmt.Println(anotherUrl)
}
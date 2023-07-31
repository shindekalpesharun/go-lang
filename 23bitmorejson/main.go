package main

import(
	"fmt"
	"encoding/json"
)

type course struct{
	Name string `json:"coursename"`
	Price int
	Platform string `json:"website"`
	Password string	`json:"-"`
	Tags []string	`json:"tags,omitempty"`
}

func main()  {
	fmt.Println("Welcome to JSON Video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson(){
	lcoCourses := []course{
		{
			"ReactJs Bootcamp",
			299,
			"learncodeonline.in",
			"test",
			[]string{"web-dev", "js"},
		},
		{
			"MERN Bootcamp",
			199,
			"learncodeonline.in",
			"test",
			[]string{"full-stack", "js"},
		},
		{
			"Angular Bootcamp",
			299,
			"learncodeonline.in",
			"test",
			nil,
		},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJs Bootcamp",
			"Price": 299,
			"website": "learncodeonline.in",
			"tags": ["web-dev", "js"]
		}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid{
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	}else{
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Println("%#v\n", myOnlineData)

	for k, v := range myOnlineData{
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
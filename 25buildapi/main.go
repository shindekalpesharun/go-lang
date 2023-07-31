package main

import(
	"fmt"
	"net/http"
	"encoding/json"
)

// model for course - file
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// fake DB
var course []course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}


func main()  {
	fmt.Println("hi")
}

// controllers - files

// serve home route

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encoder(course)
}

func getOneCourse(w http.ResponseWriter,r *http.Request){
	
}
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Model for course - file

type Course struct {
	CourseID    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"course_price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

// in-line DB
var courses []Course

// Middleware, Utility functions - file
func (c *Course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

func main() {
}

// Controllers -file

// server home route

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome To Build API in golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

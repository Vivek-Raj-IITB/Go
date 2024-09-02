package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/exp/rand"
)

// Model for course -file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname  string `json:"fullname"`
	Website   string `json:"website"`
	AvatarUrl string `json:"avatar_url"`
}

// fake DB
var courses []Course

// middleware, helper - file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Server Started..")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "Golang", CoursePrice: 299,
		Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev", AvatarUrl: "XXXXXXXXXXXXXXX"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 199,
		Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev", AvatarUrl: "XXXXXXXXXXXXXXX"}})
	// litens to a port
	// routing
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	fmt.Println("Listening to port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

// server home route

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	// params := r.URL.Query()
	// courseId := params.Get("id")
	// or
	params := mux.Vars(r)
	courseId := params["id"]

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == courseId {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - data
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data Inside Json")
		return
	}

	// generate course id and append course into courses
	// rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	// or
	params := mux.Vars(r)
	courseId := params["id"]

	// loop through courses, find matching id and return the response
	for index, course := range courses {
		if course.CourseId == courseId {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = courseId
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course with given id")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)
	courseId := params["id"]

	// loop through courses, find matching id and return the response
	for index, course := range courses {
		if course.CourseId == courseId {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course is deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("No Course with given id")
	return
}

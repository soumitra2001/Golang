package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Building apis in Golang")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("Get")

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "React Js", CoursePrice: 299, Author: &Author{FullName: "Supriya Ghorai", Website: "supriya.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Java with Spring", CoursePrice: 599, Author: &Author{FullName: "Supriya Ghorai", Website: "codewithme.in"}})

	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	r.HandleFunc("/courses", deleteAllCourse).Methods("DELETE")
	r.HandleFunc("/course/price/{id}", updateCoursePrice).Methods("PUT")

	log.Fatal(http.ListenAndServe(":4000", r))
}

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// create some middleware, helper ->

func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

// controllers =>

// serve home
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Online Course Platform</h1>"))
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
	params := mux.Vars(r)

	// loop through the courses, find matching id and return responce

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course was found with the id: " + params["id"])
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// fetch the request body data, check the data is valid or not
	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	fmt.Println(err)
	if r.Body == nil {
		json.NewEncoder(w).Encode("No course data found, please provide some data")
		return
	} else if course.isEmpty() {
		json.NewEncoder(w).Encode("Not a valid course ot create")
		return
	} else if !isUnique(course) {
		json.NewEncoder(w).Encode("This course already present")
		return
	}

	// generate unique id, string
	// add course to our fake DB

	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode("Your course added successfully")
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application-json")

	// fetch the id from request
	params := mux.Vars(r)

	// loop,if get id, remove course, add course from request body with this id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var myCourse Course

			json.NewDecoder(r.Body).Decode(&myCourse)

			if r.Body == nil {
				json.NewEncoder(w).Encode("No course data found, please provide some data")
				return
			} else if myCourse.isEmpty() {
				json.NewEncoder(w).Encode("Not a valid course ot create")
				return
			} else if !isUnique(myCourse) {
				json.NewEncoder(w).Encode("This course already present")
				return
			}

			myCourse.CourseId = params["id"]
			courses = append(courses, myCourse)

			json.NewEncoder(w).Encode(myCourse)
			return
		}
	}

	json.NewEncoder(w).Encode("Please provide a valid id to update course")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application-json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted with Id: " + params["id"])
			return
		}
	}

	json.NewEncoder(w).Encode("Invalid course id")
	return
}

func deleteAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete all course")
	var msg string
	json.NewDecoder(r.Body).Decode(&msg)

	if msg == "supriya" {
		courses = []Course{}
		json.NewEncoder(w).Encode(courses)
		return
	}

	json.NewEncoder(w).Encode("Only authenticate user can delete all courses")
	return
}

func updateCoursePrice(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course price")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			var price int
			json.NewDecoder(r.Body).Decode(&price)

			course.CoursePrice = price
			courses = append(courses[:index], courses[index+1:]...)
			courses = append(courses, course)
			json.NewEncoder(w).Encode("Course price updated")
			return
		}
	}

	json.NewEncoder(w).Encode("Please provide a valid id to update")
	return
}

func isUnique(c Course) bool {
	for _, course := range courses {
		if course.CourseName == c.CourseName {
			return course.Author.FullName != c.Author.FullName
		}
	}

	return true
}

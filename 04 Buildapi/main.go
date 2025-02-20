// CRUD of couses, use helper for uniq validation
// Used slices as DB is not intgrated here
// Used gorilla/mux for routing

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

// Models Start
type Course struct {
	CourseId    int     `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname`
	Website  string `json:"website"`
}

// Models End

// Fake DB Start
var courses []Course

// Fake DB End

// Helpers Start
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// Helper End

func main() {
	fmt.Println("Course API module")
	r := mux.NewRouter()

	// Seeding data to the slice
	courses = append(courses, Course{CourseId: 1, CourseName: "ROR", CoursePrice: 299, Author: &Author{Fullname: "Radhika R Parmar", Website: "www.google.com"}})
	courses = append(courses, Course{CourseId: 2, CourseName: "Python", CoursePrice: 199, Author: &Author{Fullname: "Radhika R Parmar", Website: "www.google.com"}})

	// Routings
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	// Listen to a port
	log.Fatal(http.ListenAndServe(":3000", r))
}

// Controller
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API Module.</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab the id of course
	params := mux.Vars(r)

	// loop through the array and compare the id and return it
	for _, course := range courses {
		if strconv.Itoa(course.CourseId) == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with Given id")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add cours")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Invalid body")
		return
	}

	// Decode the json value
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	for _, obj := range courses {
		if obj.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Duplicated course")
			return
		}
	}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	// Generate random courseid
	RandomInteger := rand.Intn(100)
	course.CourseId = RandomInteger
	// Append new course to course slice
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update cours")
	w.Header().Set("Content-Type", "application/json")

	// Get id from the request
	params := mux.Vars(r)

	// loop , remove the id and add with same id
	for index, course := range courses {
		if strconv.Itoa(course.CourseId) == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId, _ = strconv.Atoi(params["id"])
			courses = append(courses, course)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course")
	w.Header().Set("Content-Type", "application/json")

	// Get id
	params := mux.Vars(r)

	// loop trough slice and remove and return
	for index, course := range courses {
		if strconv.Itoa(course.CourseId) == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(courses)
	return
}

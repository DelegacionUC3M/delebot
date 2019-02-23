package main

import (
	"fmt"
)

// Delegate has all the info required to contact a delegate
type Delegate struct {
	Name    string
	Surname string
	NIA     string
}

// ParseCourse returns the course queried as an int
func ParseCourse(course string) int {

	switch course {
	case "primero":
	case "1":
		return 1
	case "segundo":
	case "2":
		return 2
	case "tercero":
	case "3":
		return 3
	case "cuarto":
	case "4":
		return 4
	}

	// None of the courses match
	return -1
}

// CourseQuery returns the query depending on the course provided
func CourseQuery(course int) string {
	query := fmt.Sprintf("SELECT name, surname, nia FROM delegates WHERE course=%d", course)

	return query
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"

	"github.com/brianvoe/gofakeit/v6"
)

type Student struct {
	ID               int
	Name             string
	Cohort           Cohort
	CompletedCourses []Course
}

type Cohort struct {
	ID                  int
	Name                string
	RequiredCourses     []Course
	RequiredDepartments []string
}

type Course struct {
	ID               int
	Name             string
	Department       string
	Instructors      []Instructor
	Prerequisites    []Course
	MaximumClassSize int
}

type Instructor struct {
	ID   int
	Name string
}

type MeetingTime struct {
	ID        int
	Name      string
	Day       string
	StartTime int
}

type Class struct {
	Course     Course
	Instructor Instructor
	Students   []Student
}

type Schedule struct {
	ID        int
	Matrix    [][]Class
	Conflicts int
}

func (schedule *Schedule) Print() {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintf(w, "==================\n")
	fmt.Fprintf(w, "Schedule: %v; Conflicts: %v\n", schedule.ID, schedule.Conflicts)
	for i, row := range schedule.Matrix {
		if i == 0 {
			tempLen := len(row)
			for j := 1; j <= tempLen; j++ {
				fmt.Fprintf(w, "\t%v", j)
			}
			fmt.Fprintf(w, "\n")
		}
		for _, class := range row {
			fmt.Fprintf(w, "\t%v (%v/%v)", class.Course.Name, len(class.Students), class.Course.MaximumClassSize)
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
}

type CommandLinePrintable interface {
	Print()
}

type School struct {
	Name         string `fake:"{company}"`
	MeetingTimes []MeetingTime
	Courses      []Course
	Students     []Student
	Instructors  []Instructor
	Cohorts      []Cohort
}

func CreateSchedule(school School) {
	schedule := Schedule{
		ID:     1,
		Matrix: [][]Class{},
	}
	for _, c := range school.Courses {
		schedule.Matrix = append(schedule.Matrix, []Class{})
		totalStudents := len(school.Students)
		requiredSections := totalStudents / c.MaximumClassSize
		studentsPerSection := totalStudents / requiredSections
		if totalStudents%c.MaximumClassSize > 0 {
			requiredSections += 1
			studentsPerSection += 1
		}
		classes := make([]Class, requiredSections)
		for x := 0; x < requiredSections; x++ {
			class := Class{
				Course:   c,
				Students: []Student{},
			}

			classes[x] = class
		}
	}
	schedule.Print()
}

func NewStudentsList(n int) []Student {
	if n == 0 {
		n = rand.Intn(80)
	}
	var students []Student
	for len(students) < n {
		students = append(students, Student{
			Name: gofakeit.Name(),
		})
	}
	return students
}

func NewCourseList(n int) []Course {
	var courses []Course
	for i := 0; i < n; i++ {
		courses = append(courses, Course{
			Name:             fmt.Sprintf("%v %v %v", gofakeit.Adjective(), gofakeit.Animal(), gofakeit.JobTitle()),
			MaximumClassSize: gofakeit.RandomInt([]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}),
		})
	}
	return courses
}

func main() {
	school := School{
		Name:     gofakeit.Dinner(),
		Courses:  NewCourseList(12),
		Students: NewStudentsList(80),
	}
	CreateSchedule(school)
}

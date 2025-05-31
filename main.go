package main

import (
	"fmt"
)

// структури - це визначенний користувачем тип, який дозволяє комбінувати елементи різних типів в одит тип.
// struct is value type. NOT REFERENCE TYPE
type Student struct {
	Name    string
	Age     int
	Groupe  string
	Avarage float64
}
type School struct {
	Students []Student
}

func main() {
	student1 := Student{
		Name:    "Jo",
		Age:     23,
		Groupe:  "IT-52",
		Avarage: 52.2,
	}

	student2 := Student{
		Name: "Qwer",
		Age:  23,
		//group, avg -> null
	}

	var student3 Student
	student3.Name = "Andrew"
	student3.Age = 18
	student3.Groupe = "IT-912"
	student3.Avarage = 10.5

	fmt.Printf("Student 1: %+v\n", student1)
	fmt.Printf("Student 1: %+v\n", student2)
	fmt.Printf("Student 1: %+v\n", student3)

	fmt.Printf("Student1 name: %s\n", student1.Name)

	school := School{
		Students: []Student{student1, student2, student3},
	}

	studentsSlice := []Student{student1, student2, student3}

	school2 := School{
		Students: studentsSlice,
	}

	fmt.Printf("School: %v\n", school)
	fmt.Printf("School: %v\n", school2)
}

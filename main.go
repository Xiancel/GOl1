package main

import "fmt"

type Studen struct {
	Name    string
	Age     int
	Group   string
	Avarage float64
}

func main() {
	a := Studen{Name: "Shanks", Age: 39}
	b := &a //b - вказівник на a
	b.Name = "Bob"

	fmt.Println(a.Name)
	fmt.Println(b.Name)
}

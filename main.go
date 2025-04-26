package main

import "fmt"

func main() {
	const Pi = 3.14159
	const MaxUsers = 50

	const (
		StatusOk           = 200
		StatusNotFound     = 404
		StatusInternalEror = 500
	)

	fmt.Print(StatusOk)
}

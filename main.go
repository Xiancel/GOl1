package main

import (
	"fmt"
)

// custome type
type names string
type km float64

func main() {
	var jrl names = "Joseph"
	var jr2 names = "Dio"
	fmt.Println(jrl)
	fmt.Println(jr2)
	var section km = 64.5
	fmt.Println(section)
}

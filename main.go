package main

import "fmt"

func main() {
	value := 55

	if true {
		value := 20
		fmt.Print(value)
	}

	fmt.Print(value)
}

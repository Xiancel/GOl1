package main

import "fmt"

func main() {
	PrintAnything(42)
	PrintAnything(52.2)
	PrintAnything("Welcome to the club body")
	PrintAnything(true)
	PrintAnything([]int{1, 52, 78, 9})
	PrintAnything('H')

}

func PrintAnything(value interface{}) {
	fmt.Printf("Value: %v, Type: %T \n", value, value)
}

package main

import (
	"fmt"
)

func main() {
	name := []string{"Jessie Pinkman", "Walter White", "Gustavo Fring"}

	for index, name := range name {
		fmt.Printf("Index %d: %s\n", index, name)
	}

	for _, name := range name {
		fmt.Printf("Hellow - %s\n", name)
	}
}

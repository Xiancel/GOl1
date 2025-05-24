package main

import (
	"fmt"
)

func main() {
	names := make([]string, 3, 5) //["","","","",""]
	names[0] = "Oleg"
	names[1] = "Oleksandr"
	names[2] = "Joseph"
	fmt.Println(names)
	fmt.Println("Lenght", len(names))
	fmt.Println("Capacity", cap(names))

	s := make([]int, 0, 3)
	fmt.Printf("Slice: %v, Lenght: %d, Capacity: %d\n", s, len(s), cap(s))

	fmt.Println(" --- Післе циклу ---")
	for i := 1; i <= 7; i++ {
		s = append(s, i)
		fmt.Printf("Slice: %v, Lenght: %d, Capacity: %d\n", s, len(s), cap(s))
	}
}

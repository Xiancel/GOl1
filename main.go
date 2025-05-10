package main

import (
	"fmt"
)

func main() {
	var a int
	for i := 0; i <= 3; i++ {
		a = a + 2
		fmt.Printf("Iteration: %d \n", i)
	}

	fmt.Println(a)
}

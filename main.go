package main

import (
	"fmt"
)

func main() {
	marks := map[string]int{
		"Andrew":     52,
		"Gustav":     42,
		"Starii Bog": 99,
	}

	for name, mark := range marks {
		fmt.Printf("%s: %d\n", name, mark)
	}

	for name := range marks {
		fmt.Println("Name: ", name)
		if name == "Gustav" {
			break
		}
	}

outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("Stop outerLoop")
				break outerLoop
			}
			fmt.Printf("i: %d, j: %d \n", i, j)
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Пране число: ", i)
	}

	for i := 0; i < 10; i++ {
		if i > 3 && i < 7 {
			continue
		}
		fmt.Println("число: ", i)
	}

}

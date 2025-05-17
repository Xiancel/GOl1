package main

import (
	"fmt"
)

func main() {
	x := 20
	func() {
		fmt.Println("Hellow from anonymous func", x)
	}()

	func(name string) {
		fmt.Println("Hellow from anonymous func", name)
	}("GoLang")

	square := func(n int) int {
		return n * n
	}

	fmt.Println(square(8))
}

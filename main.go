package main

import (
	"fmt"
)

func main() {
	names := []string{"Walter", "Jesse"}
	names = append(names, "Oleg")
	fmt.Println(names)

	set2Names := []string{"Law", "Zoro", "Shanks"}
	names = append(names, set2Names...)
	fmt.Println(names)

	fmt.Println(" --- Після видалення --- ")
	names = removeElement(names, 2)
	fmt.Println(names)

	names = removeElement(names, 2)
	fmt.Println(names)
}

func removeElement(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func removeFastElement(slice []string, index int) []string {
	slice[index] = slice[len(slice)-1]

	return slice[:len(slice)-1]
}

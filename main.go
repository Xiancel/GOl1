package main

import (
	"fmt"
)

func main() {
	var rating int
	fmt.Scanln(&rating)
	switch rating {
	case 5:
		fmt.Println("very Nice :))")
	case 4:
		fmt.Println("nice :)")
	case 3:
		fmt.Println("good :|")
	case 1, 2:
		fmt.Println("BAd :(")
	default:
		fmt.Println("некоректна оцінка")
	}
}

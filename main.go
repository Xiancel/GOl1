package main

import (
	"fmt"
)

func main() {
	//слайс
	//var ім'я []тип
	//var number []int

	arr := [5]int{1, 2, 3, 4, 5}

	slice1 := arr[1:4]
	fmt.Println(slice1)

	slice2 := arr[:3]
	slice3 := arr[2:]
	slice4 := arr

	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
}

package main

import (
	"fmt"
)

func output(a int) {
	fmt.Println(a)
}
func main() {
	var a int

	fmt.Scanln(&a)

	output(a)
}

// func functionName (par1 type1, par2 type2, ...)return_type{
// 	//function body
// 	return value
// }

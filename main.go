package main

import (
	"fmt"
)

func main() {
	password := "i123nqwe"
	if len(password) >= 8 {
		fmt.Println("Password has successful leng")
	} else {
		fmt.Println("Password no has successful leng")
	}

}

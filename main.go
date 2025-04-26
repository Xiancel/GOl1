package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int8 = 122
	var b float32 = 2.15

	var str string = "Hellow 1231232131231232131231231231231231231231"
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(str))
}

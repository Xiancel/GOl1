package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	currentHour := time.Now().Hour()
	fmt.Println(currentHour)

	if currentHour >= 10 && currentHour < 17 {
		fmt.Println("зараз день")
	} else {
		fmt.Println("зараз ніч")
	}

	if rand.Intn(2) == 0 {
		fmt.Println("Орел")
	} else {
		fmt.Println("Решка")
	}

	satusCode := 200
	if satusCode >= 200 && satusCode < 300 {
		fmt.Println("OK")
	} else {
		fmt.Println("Nook")
	}
}

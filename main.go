package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(divide(10, 2))

	q, r := divideWithRemainder(13, 5)
	fmt.Printf("13 / 5 = %d with remainder %d\n", q, r)

	number, _, _ := getDate()
	fmt.Println("Number: ", number)

	_, _, isValid := getDate()
	fmt.Println("Correct: ", isValid)

	res, err := divideErr(10, 2)
	if err != nil {
		fmt.Println("Errors: ", err)
	} else {
		fmt.Println("Result: ", res)
	}

	res, err = divideErr(10, 0)
	if err != nil {
		fmt.Println("Errors: ", err)
	} else {
		fmt.Println("Result: ", res)
	}
}

func divide(a, b float64) (res float64) {
	res = a / b
	return
}

func divideWithRemainder(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func getDate() (int, string, bool) {
	return 42, "hellow", true
}

func divideErr(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

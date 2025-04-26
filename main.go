package main

import (
	"fmt"
	"math"
)

func main() {
	a := 20
	b := 11

	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	dif := a - b
	fmt.Printf("%d - %d = %d\n", a, b, dif)

	prod := a * b
	fmt.Printf("%d * %d = %d\n", a, b, prod)

	quo := a / b
	fmt.Printf("%d / %d = %d\n", a, b, quo)

	rem := a % b
	fmt.Printf("%d %% %d = %d\n", a, b, rem)

	FloatQ := float64(a) / float64(b)
	fmt.Printf("%d / %d = %.3f\n", a, b, FloatQ)

	fmt.Print(math.Pow(float64(a), 2))
}

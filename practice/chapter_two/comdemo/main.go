package main

import (
	"fmt"
)

func main() {
	var (
		a = 2
		b = 10
	)
	fmt.Println("False =", a == b)
	fmt.Println("True =", a != b)
	fmt.Println("False =", a > b)
	fmt.Println("True =", a < b)
	fmt.Println("False =", a >= b)
	fmt.Println("True =", a <= b)
}

package main

import "fmt"

func main() {
	var a, b float32

	a = 25
	b = 15

	add := a + b
	fmt.Println("Addition = ", add)

	sub := a - b
	fmt.Println("Subtaction = ", sub)

	multiplication := a * b
	fmt.Println("Multiplication = ", multiplication)

	division := a / b
	fmt.Println("Division = ", division)
}

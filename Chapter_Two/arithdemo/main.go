package main

import "fmt"

func main() {
	var a, b int

	a = 30
	b = 15

	add := a + b
	fmt.Printf("%d + %d = %d  \n", a, b, add)

	sub := a - b
	fmt.Printf("%d - %d = %d  \n", a, b, sub)

	multiplication := a * b
	fmt.Printf("%d * %d = %d \n", a, b, multiplication)

	division := a / b
	fmt.Printf("%d / %d = %d  \n", a, b, division)

}

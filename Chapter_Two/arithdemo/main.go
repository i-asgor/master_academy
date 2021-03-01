package main

import "fmt"

func main() {
	var a, b int

	a = 55
	b = 15

	add := a + b
	fmt.Printf("%d + %d = %d  \n", a, b, add)

	sub := a - b
	fmt.Printf("%d - %d = %d  \n", a, b, sub)

	multiplication := a * b
	fmt.Printf("%d * %d = %d \n", a, b, multiplication)

	division := float32(a) / float32(b)
	fmt.Printf("%d / %d = %.2f  \n", a, b, division)

	modulus := a % b
	fmt.Printf("%d %% %d = %d", a, b, modulus)

}

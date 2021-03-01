package main

import "fmt"

func main() {
	a := 4

	// increment
	fmt.Printf("a = %d\n", a)
	a = a + 1
	fmt.Printf("a + 1 = %d\n", a)
	a++
	fmt.Printf("a ++ = %d\n", a)

	fmt.Println("Decrement")
	// decrement
	fmt.Printf("a = %d\n", a)
	a = a - 1
	fmt.Printf("a - 1 = %d\n", a)
	a--
	fmt.Printf("a -- = %d\n", a)

}

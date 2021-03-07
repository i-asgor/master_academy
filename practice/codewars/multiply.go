package main

import "fmt"

func Multiply(a, b int) int {
	return a * b

}

func main() {
	multiplication := Multiply(10, 2)
	fmt.Println(multiplication)
}

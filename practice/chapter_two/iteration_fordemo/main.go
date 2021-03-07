package main

import "fmt"

func main() {
	var n int = 10
	for i := 0; i <= n; i++ {
		fmt.Println("value i =", i)
	}
	fmt.Println("Decrement value print")
	for j := 10; j >= 0; j-- {
		fmt.Println("value j =", j)
	}
}

package main

import "fmt"

func main() {
	var (
		a = 5
		b = 8
	)
	if a > b || b-a < a {
		fmt.Println("conditional a>b || b-a<a")
	} else {
		fmt.Println("Else loop executed")
	}
}

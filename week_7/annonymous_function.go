package main

import "fmt"

func main() {
	a := func(l int, b int) (r int) {
		r = l*b
		return
	}(20, 30)
	fmt.Println(a) 
}
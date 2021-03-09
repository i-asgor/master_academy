package main

import (
	"fmt"
)

func main() {
	fmt.Println(Opposite(25))
}

func Opposite(value int) int {
	return -value
}

package main

import "fmt"

func Solution(word string) string {
	runes := []rune(word)
	fmt.Println(runes)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(Solution("this is asgor"))
}

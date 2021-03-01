package main

import "fmt"

func main() {
	selected := 2

	switch selected {
	case 0:
		fmt.Println("Selected = 0")
	case 1:
		fmt.Println("Selected = 1")
	case 2:
		fmt.Println("Selected = 2")
	case 3:
		fmt.Println("Selected = 3")
	default:
		fmt.Println("Default")
	}
}

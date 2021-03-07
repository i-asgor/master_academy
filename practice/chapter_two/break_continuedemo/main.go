package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 2 {
			break
		}
		fmt.Println("value of i = ", i)
	}

	for j := 11; j <= 20; j++ {
		if j == 13 {
			continue
		}
		fmt.Println("value of j = ", j)
	}
}

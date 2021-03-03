package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 10; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
			// fmt.Println(i)
		}

	}
	fmt.Println(sum)
}

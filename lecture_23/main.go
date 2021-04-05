package main

import "fmt"

func main() {
	result := Sum(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(result)
}

func Sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}

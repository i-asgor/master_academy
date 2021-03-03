package main

import "fmt"

func SquareSum(numbers [3]int) int {
	sum := 0
	for _, val := range numbers {
		s := val ^ 2
		fmt.Println(s)
		sum = sum + s
	}
	return sum
}

func main() {
	var myArr = [3]int{1, 2, 3}
	fmt.Println(SquareSum(myArr))
}

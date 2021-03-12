package main

import "fmt"

func main() {
	fmt.Println(maximumSubarraySum([]int{ -2, 1, 3, 4, -1, 2, 1, -5, 4}))
}

func maximumSubarraySum(numbers []int) int {
	max := 0
	result := 0
	for _, i := range numbers {
		max += i
		fmt.Println(max)
		if max < 0 {
			max = 0
		} else {
			if max > result {
				result = max
			}
		}
	}
	return result
}
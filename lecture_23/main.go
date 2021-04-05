package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("All Int value : ", result)

	// result2 := Sum2(1, 2, 3, 4, 5, "6", 7, 8, "9", 10)
	// fmt.Println("All type Number value : ", result2)

	strSlc := []string{"10", "20", "30"}
	var res int64
	for i, slc := range strSlc {

		res += Sum2(slc)
		fmt.Println(i, res)
	}
	fmt.Println("str number add : ", res)
}

func Sum2(nums ...interface{}) int64 {
	var res int64
	for _, n := range nums {
		strval := fmt.Sprintf(`%v`, n)
		nint, _ := strconv.ParseInt(strval, 10, 64)
		res += nint
	}
	return res
}

func Sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	result := Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("All Int value : ", result)

	// result2 := Sum2(1, 2, 3, 4, 5, "6", 7, 8, "9", 10)
	// fmt.Println("All type Number value : ", result2)

	strSlc := []interface{}{"10", "20", "30"}

	result2 := Sum2(strSlc...)
	fmt.Println("interface value : ", result2)
	// var res int64
	// for i, slc := range strSlc {

	// 	res += Sum2(slc)
	// 	fmt.Println(i, res)
	// }
	// fmt.Println("str number add : ", res)

	stooges := [...]string{"asgor", "nasarul", "mahmud"}
	bs, err := json.Marshal(stooges)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(bs)

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

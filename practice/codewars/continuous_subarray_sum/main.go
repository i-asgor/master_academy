package main
import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if k != 0 {
			sum %= k
		}
		prev, ok := m[sum]
		if ok {
			if i-prev > 1 {
				return true
			}
		} else {
			m[sum] = i
		}
	}
	return false
}

func main(){
	fmt.Println(checkSubarraySum(25))
}
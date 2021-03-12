package main

import "fmt"

func PositiveSum(numbers []int) int {
	 sum := 0
	 for i := 0; i < len(numbers); i++ {
     if(numbers(i) > 0) {
     sum += numbers[i];
     }
    }
    return sum;
}

func main(){
	primes := [...]int{1,2,3}
	fmt.Println(PositiveSum(primes))
}
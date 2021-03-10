package main

import "fmt"

func SquareSum(numbers []int) int{
	sum :=0
	
	for _,v := range numbers{
	 sum += (v*v)
	}
	return sum
}

func main(){
	primes := [...]int{1,2,3}
	fmt.Println(SquareSum(primes))
}
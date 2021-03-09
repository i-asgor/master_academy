package main

import "fmt"

func IsDivisible(n, x, y int) bool {
    if n%x ==0 && n%y==0{
      return true
    }
  return false
}

func main(){
	fmt.Println(IsDivisible(35,5,8))
}
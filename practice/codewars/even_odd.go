package main  
  
import (  
 "fmt"  
)  
  
func main() {  
 var numbs = []int{51, 4, 5, 6, 24, 11}  
 for _, number := range numbs {  
  if number%2 == 0 {  
   fmt.Printf("%v is a Even Number\n", number)  
  } else {  
   fmt.Printf("%v is a Odd Number\n", number)  
  }  
 }  
} 
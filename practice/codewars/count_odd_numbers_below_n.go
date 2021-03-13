package main
import "fmt"

func OddCount(n int) int{
  for i:=1;i<=n;i++{
    if i%2==1{
      return i
    }
  }
  return i
}

func main(){
	fmt.Println(OddCount(7))
}
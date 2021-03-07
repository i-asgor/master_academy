package main

import "fmt"
//example-1
/*func add(x int,y int)int{
  r:=x+y
  return r
}*/

//example-2
/*func add(x,y int) int{
	r := x+y
	return r
}*/

//example-3 
/*func add(x,y int) (r int){
	r = x+y
	return r
}*/

//example-4
/*func add(x,y int) (r int){
	r = x+y
	return
}*/

func add(x,y int)int{
	return x+y
}

/*func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}*/

func main(){
	x:=add(40,30)
	//a,p := rectangle(10,20)
	fmt.Println(x)

}
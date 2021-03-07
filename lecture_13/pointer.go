package main

import "fmt"

func update(x *int){
	fmt.Println("update function x memory adress =",x)
	*x = *x+10
}

func main(){
 var a int
 var b *int

  fmt.Println("a value is = ",a)
  fmt.Println("a memory address is = ", &a)

  fmt.Println("b value is = ",b)
  fmt.Println("b memory address is = ", &b)
  a=10 //assignment
  b=&a //referencing
  fmt.Println("a value is = ",a)//accessing
  fmt.Println("b value is = ",b) // print a memory address
  fmt.Println("b dereferencing value is = ",*b) // Print a value

   update(&a)
   fmt.Println("Update function value update = ",a)
}
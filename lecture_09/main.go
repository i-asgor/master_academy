package main

import "fmt"

func main(){
 fmt.Println(true && true) 
 fmt.Println(true && false) 
 fmt.Println(true || true) 
 fmt.Println(true || true) 
 fmt.Println(!true) 

 	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var chr rune 
	chr = 'A'
	fmt.Println(chr)
	fmt.Printf("%c",chr)

}
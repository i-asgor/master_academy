package main
import (
	"fmt"
	"reflect"
)

func main(){
	var a [3]string
	
	fmt.Println(a)
	fmt.Println(len(a))
	
	var b [3]string
	b[0] = "asgor"
	b[1] = "abdullah"
	b[2] = "jony"
	fmt.Println("")
	fmt.Println("all array output show")
	fmt.Println(b)
	fmt.Println("")
	fmt.Println("single array output show")
	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])
	fmt.Println("")
	//short hand declaration
	fmt.Println("//short hand declaration and used ... array not defined")
	c := [...]string{"abdullah","asgor","jony","ismile"}
	fmt.Println(c)
	fmt.Printf("%T\n",c)
	fmt.Println(len(c))
	s := reflect.TypeOf(c).Kind().String()
	fmt.Println(s)
	
	//slice
	d :=c[0:2]
	fmt.Println(d)
	fmt.Printf("%T\n",d)
	
	
}
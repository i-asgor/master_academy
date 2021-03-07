package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X,Y float64
} 

func (v Vertex) Absolute() float64{
	return math.Sqrt(v.X*v.Y+v.X*v.Y)
}

func main(){
	v := Vertex{2,4}
	fmt.Println(v.Absolute())	
}
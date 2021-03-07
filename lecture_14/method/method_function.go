package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Absolute(v Vertex) float64{
	return math.Sqrt(v.X*v.Y + v.X*v.Y)
}

func main(){
	v := Vertex{2,4}
	fmt.Println(Absolute(v))
}
package main

import (
	"fmt"
	"master_academy/lecture_28/vatcalculator"
)

func main() {
	i := vatcalculator.InclusiveTax(200, 15)
	fmt.Println("Inclusive tax: ", i)
	
	
	e := vatcalculator.ExclusiveTax(200, 15)
	fmt.Println("Exclusive tax : ", e)
}

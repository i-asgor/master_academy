package main

import (
	"fmt"
	"strings"
)

func MakeUpperCase(str string) string{
	return strings.ToUpper(str)
}

func main(){
	fmt.Println(MakeUpperCase("asgor"))
}
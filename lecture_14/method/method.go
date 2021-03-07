package main

import "fmt"

type Doctor struct{
	Name string
	Education string
	Age int
	Salary float32
}

//method
func (d Doctor)Speak(){
	fmt.Println(d.Name," can speak")	
}

func (d Doctor)Sallery(){
	fmt.Println(d.Name,"Salary is",d.Salary)
}
func main(){
	//var d = Doctor{"Tarik","MBBS",30,50000.00}
	//var d = Doctor{Name:"Tarek",Education:"MBBS",Age:35,Salary:50000.00}
	var d Doctor
	d.Name = "Tarek"
	d.Education = "MBBS & FPPS"
	d.Age = 30
	d.Salary = 60000
	fmt.Println(d.Name,d.Education,d.Age)
	d.Speak()
	d.Sallery()	
}
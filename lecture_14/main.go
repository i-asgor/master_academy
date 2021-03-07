package main

import "fmt"

type doctor struct{
	Name string
	Education string
	Age int
	Salary float32
}

func (d doctor)Speak(){
	fmt.Println(d.Name, "can speak")
}

func main(){
	//var d = doctor{"Tarek","MBBS", 30,50000}
	var d doctor
	d.Name = "Tarek"
	d.Education = "MBBS"
	d.Age = 30
	d.Salary = 50000
	fmt.Println(d.Name,d.Education,d.Age,d.Salary)
}
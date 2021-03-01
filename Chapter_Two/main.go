package main

import "fmt"

func main() {
	var str string
	var n, m int
	var mn float32

	str = "Md. Asgor"
	n = 10
	m = 20
	mn = 25.3

	fmt.Println("value of str", str)
	fmt.Println("value of n", n)
	fmt.Println("value of m", m)
	fmt.Println("value of mn", mn)

	fmt.Println("===========================")
	var city string = "Dhaka"
	var x int = 100

	fmt.Println("value of city : ", city)
	fmt.Println("value of x : ", x)

	fmt.Println("===========================")

	country := "Bangladesh"
	de := 15

	fmt.Println("values of country : ", country)
	fmt.Println("values of de : ", de)
	fmt.Println("===========================")

	var (
		name  string
		email string
		age   int
	)
	name = "Asgor"
	email = "asgor.ice@gmail.com"
	age = 30

	fmt.Println("my name is : ", name)
	fmt.Println("my email address : ", email)
	fmt.Println(name, " age is : ", age)
}

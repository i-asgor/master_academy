package main

import "fmt"

/* type Book struct {
	Title string
	Author string
	ISBN string
	Price float32
	pages int
} */

func main(){
	//var b1 Book// dot notation
	//b1.Title = "AN INTRODUCTION TO PROGRAMMING IN GO"
	//b1.Author = "CALEB DOXSEY"
	//b1.ISBN = "978-1478355823"
	//b1.Price =10.5
	//b1.pages =165
	
b1 := struct {
	Title string
	Author string
	ISBN string
	Price float32
	pages int
}{
	Title : "AN INTRODUCTION TO PROGRAMMING IN GO",
	Author : "CALEB DOXSEY",
	ISBN : "978-1478355823",
	Price : 10.5,
	pages : 165,
}
	
	fmt.Println(b1)
	fmt.Println(b1.Title)
	fmt.Println(b1.Author)
	fmt.Println(b1.ISBN)
	fmt.Println(b1.Price)
	fmt.Println(b1.pages)
	
	var w1 Weight
	fmt.Println(w1,name)

}
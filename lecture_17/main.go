package main

import (
	"fmt"
	"net/http"
)

//type Handler interface{
//	ServeHTTP(ResponseWriter, *Request)
//}

func main() {

	//var name Datatype
	//var x string
	//var handler func(ResponseWriter,*Request)
	http.HandleFunc("/", home)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Welcome to my first golang webpage with Coding Bootcamp</h1>`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to Coding Bootcamp contact page`)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to Coding Bootcamp about page`)
}

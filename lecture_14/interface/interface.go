package main

import "fmt"

type Email interface{
	Write(string, string, string)
	Send()
	Read()
}

type Test struct{
	To string
	From string
	Subject string
	Message string
}

func (t Test)Write(to string, from string, subject string){
	fmt.Println(to,from, subject)
}


func (t Test) Send(){
	fmt.Println(t.To,"email Sent")
}

func (t Test) Read(){
	fmt.Println(t.From, "Received from")
}

func main(){
	//var e Email
	//fmt.Println(e)
	var t Test
	t.To="asgor.ice@gmail.com"
	t.From="asgor@padakhep.org"
	t.Subject = "Test mail"
	t.Message = "Hello this is a test mail"
	t.Write(t.To, t.From,t.Subject)
}
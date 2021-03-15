package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/mateors/mcb"
)

var db *mcb.DB

type RequstTable struct {
	ID      string `json:"aid"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Type    string `json:"type"`
	Status  int    `json:"status"`
}

func init() {

	db = mcb.Connect("localhost", "asgor", "iasgor")
	res, err := db.Ping()
	if err != nil {

		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Println(res, err)

}

func main() {

	//How to insert into couchbase bucket
	var myData RequstTable

	form := make(url.Values, 0)
	form.Add("bucket", "master_academy") 
	form.Add("aid", "request::3")               
	form.Add("name", "Md shumon")
	form.Add("company", "Master Academy")
	form.Add("email", "shumon@gmail.com")
	form.Add("type", "request") 
	form.Add("status", "1")

	p := db.Insert(form, &myData)    
	fmt.Println("Status:", p.Status) 


}
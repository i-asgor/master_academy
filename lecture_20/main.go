package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql" //underscore means building package or building method handle
)

var db *sql.DB
var err error

func init() {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err = sql.Open("mysql", "root:iasgor@tcp(127.0.0.1:3306)/hosting_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	// defer db.Close()

	fmt.Println("db connection successfully")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/request", request)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer((http.Dir("assets")))))
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

func features(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

func docs(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

func request(w http.ResponseWriter, r *http.Request) {

	//method 1
	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")

	// fmt.Println(name, company, email)

	// fmt.Fprintf(w, "Received %s %s %s", name, company, email) //response

	//method 2
	// r.ParseForm()
	// for key, val := range r.Form {
	// 	fmt.Println(key, val) // key = string and value = slice of string
	// }
	qs := "INSERT INTO `request` (`id`,`name`,`company`,`email`,`status`) VALUES (NULL, '%s', '%s','%s','1')"
	sql := fmt.Sprintf(qs, name, company, email)
	insert, err := db.Query(sql)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	fmt.Fprintf(w, "ok")
}

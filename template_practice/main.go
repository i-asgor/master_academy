package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	ptmp, err := template.ParseFiles("one.html", "two.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(os.Stdout, "asgor")
}

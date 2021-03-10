package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	ptmp, err := template.ParseFiles("two.html", "one.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(os.Stdout, nil)
}

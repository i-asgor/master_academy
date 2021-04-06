package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(dir)

	posf, err := os.Create("asgor.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	n, err := posf.Write([]byte("This is a text file"))
	fmt.Println(n, err)

}

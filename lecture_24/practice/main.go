package main

import (
	"fmt"
	"os"
)

func main() {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(dir)

	//posf.Close()
	CreateFile("asgorfunc.txt", "This is text file with using another function")

}

func CreateFile(fileName, content string) bool {
	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer posf.Close()

	_, err = posf.Write([]byte(content))
	// fmt.Println(n, err)
	if err != nil {
		return false
	}
	return true
}

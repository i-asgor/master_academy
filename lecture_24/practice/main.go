package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}

	// fmt.Println(dir)

	//posf.Close()
	// isErr := CreateFile("asgorfunc.txt", "This is text file with using another function")
	// fmt.Println(isErr)

	// fi, err := os.Stat("asgor.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(fi.IsDir())
	// fmt.Println(fi.ModTime().Date())
	// fmt.Println(fi.Name())
	// fmt.Println(fi.Size())

	// How to make folder
	// err := os.Mkdir("master_academy", 0777)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// base := filepath.Base("./")
	// base := filepath.Base("./master_academy")
	base := filepath.Base(dir)
	relativepath := filepath.Join("master_academy")
	fullpath, _ := filepath.Abs("master_academy")
	fmt.Println(base)
	fmt.Println(relativepath, "\n", fullpath)

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

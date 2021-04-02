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

	// isErr := createFile("master_academy.txt", "hello golang is awesome")
	// fmt.Println(isErr)

	//how to make a folder
	// err := os.Mkdir("master_academy", 0777)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	base := filepath.Base(dir)
	//relativePath := filepath.Join("master_academy")
	absolutePath, _ := filepath.Abs("master_academy")

	newPath := filepath.Join(absolutePath, "..", "..")
	fmt.Println(base)
	//fmt.Println(relativePath)
	fmt.Println(absolutePath)
	fmt.Println(newPath)

	//folderPath := fmt.Sprintf(`%s`,newPath)
	os.Mkdir(newPath, 777)

}

func createFile(fileName, content string) bool {
	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer posf.Close()

	_, err = posf.Write([]byte(content))
	//fmt.Println(n, err)
	if err != nil {
		return false
	}

	return true
}

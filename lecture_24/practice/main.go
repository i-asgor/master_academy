package main

import (
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

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
	// base := filepath.Base(dir)
	// relativepath := filepath.Join("master_academy")
	// fullpath, _ := filepath.Abs("master_academy")

	// D:\GALANG\src\master_academy\lecture_24\practice\master_academy
	// newpath := filepath.Join(fullpath, "..", "..", "lecture_24_backup")
	// fmt.Println(base)
	// fmt.Println(relativepath, "\n", fullpath, "\n", newpath)

	// os.Mkdir("D:\\TEST", 777)
	// os.Rename("D:\\TEST", "D:\\TEST_01")

	// f := excelize.NewFile()
	// // Create a new sheet.
	// index := f.NewSheet("Sheet2")
	// // Set value of a cell.
	// f.SetCellValue("Sheet2", "A2", "Hello world.")
	// f.SetCellValue("Sheet1", "B2", 100)
	// // Set active sheet of the workbook.
	// f.SetActiveSheet(index)
	// // Save spreadsheet by the given path.
	// if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }

	// f, err := excelize.OpenFile("Book1.xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// // Get value from cell by given worksheet name and axis.
	// cell, err := f.GetCellValue("Sheet1", "B2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(cell)
	// // Get all the rows in the Sheet1.
	// rows, err := f.GetRows("Sheet2")
	// for _, row := range rows {
	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }

	categories := map[string]string{
		"A2": "Small", "A3": "Normal", "A4": "Large",
		"B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{
		"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	if err := f.AddChart("Sheet1", "E1", `{
        "type": "col3DClustered",
        "series": [
        {
            "name": "Sheet1!$A$2",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$2:$D$2"
        },
        {
            "name": "Sheet1!$A$3",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$3:$D$3"
        },
        {
            "name": "Sheet1!$A$4",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$4:$D$4"
        }],
        "title":
        {
            "name": "Fruit 3D Clustered Column Chart"
        }
    }`); err != nil {
		fmt.Println(err)
		return
	}
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

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

package compare

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strings"
)

func readFile(file string) map[string]bool {
	result := make(map[string]bool)
	f, err := excelize.OpenFile(file)
	if err != nil {
		panic(err)
	}
	//cellVal, err := f.GetCellValue("Sheet1", "B2")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(cellVal + " XXXXXXXXX")
	// Get all the rows in the vegan section.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		//for _, colCell := range row {
		//	fmt.Print(colCell, "\t")
		//}
		v := strings.ReplaceAll(row[1], " ", "")
		if v == "" {
			break
		}
		_, ok := result[v]
		if ok {
			fmt.Println("重复项 ： ", v)
		}
		result[v] = true
	}
	return result
}
func CompareFile(file1, file2 string) {
	book1 := readFile(file1)
	book2 := readFile(file2)
	i := 0
	for k := range book1 {
		_, ok := book2[k]
		if !ok {
			i++
			fmt.Printf("%s,", k)
			//fmt.Println("boo1 多余: ", k)
		}
	}
	fmt.Printf(";总数量%d", i)

	//for k := range book2 {
	//	_, ok := book1[k]
	//	if !ok {
	//		fmt.Println("book2 多余: ", k)
	//	}
	//}
}

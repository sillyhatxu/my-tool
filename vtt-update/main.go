package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	moduleName = "vtt-update"
	pattern    = "\\d+" //反斜杠要转义
)

func isNumber(input string) bool {
	result, _ := regexp.MatchString(pattern, input)
	return result
}

func writeFile(content string, filePathName string) {
	file, err := os.Create(filePathName) //创建文件
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.WriteString(file, content) //写入文件(字符串)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	//var wireteString = "测试n"
	//var filename = "./output1.txt"
	//var f *os.File
	//var err1 error

	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	//if checkFileIsExist(filename) { //如果文件存在
	//	f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
	//	fmt.Println("文件存在")
	//} else {
	//	f, err1 = os.Create(filePathName) //创建文件
	//	fmt.Println("文件不存在")
	//}
	//check(err1)
	//n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	//check(err1)
	//fmt.Printf("写入 %d 个字节n", n)

	///*****************************  第二种方式: 使用 ioutil.WriteFile 写入文件 ***********************************************/
	//var d1 = []byte(wireteString)
	//err2 := ioutil.WriteFile("./output2.txt", d1, 0666) //写入文件(字节数组)
	//check(err2)
	//
	///*****************************  第三种方式:  使用 File(Write,WriteString) 写入文件 ***********************************************/
	//f, err3 := os.Create("./output3.txt") //创建文件
	//check(err3)
	//defer f.Close()
	//n2, err3 := f.Write(d1) //写入文件(字节数组)
	//check(err3)
	//fmt.Printf("写入 %d 个字节n", n2)
	//n3, err3 := f.WriteString("writesn") //写入文件(字节数组)
	//fmt.Printf("写入 %d 个字节n", n3)
	//f.Sync()
	//
	///***************************** 第四种方式:  使用 bufio.NewWriter 写入文件 ***********************************************/
	//w := bufio.NewWriter(f) //创建新的 Writer 对象
	//n4, err3 := w.WriteString("bufferedn")
	//fmt.Printf("写入 %d 个字节n", n4)
	//w.Flush()
	//f.Close()
}

func replaceFile(filePath string, fileName string) {
	file, err := os.Open(fmt.Sprintf("%s/%s", filePath, fileName))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var buffer bytes.Buffer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content := scanner.Text()
		content = strings.ReplaceAll(content, "WEBVTT", "")
		content = strings.ReplaceAll(content, "<v Instructor>", "")
		content = strings.ReplaceAll(content, "<v Narrator>", "")
		content = strings.ReplaceAll(content, "<v instructor>", "")
		content = strings.ReplaceAll(content, "<v Presenter>", "")
		content = strings.ReplaceAll(content, "<v ->", "")
		content = strings.ReplaceAll(content, "</v>", "")
		content = strings.ReplaceAll(content, "<v Clement>", "")
		if isNumber(content) {
			continue
		} else if strings.Contains(content, "-->") {
			continue
		}
		if len(content) > 0 && content[len(content)-1:] != "." {
			content += " "
		}
		content = strings.ReplaceAll(content, "\n", "")
		content = strings.ReplaceAll(content, "\r", "")
		content = strings.ReplaceAll(content, ".", ". ")
		buffer.WriteString(content)
	}
	fmt.Println(buffer.String())
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	writeFile(buffer.String(), fmt.Sprintf("%s/new/%s", filePath, fileName))
}

func main() {
	pwd, _ := os.Getwd()
	currentFolder := fmt.Sprintf("%s/%s", pwd, moduleName)
	fmt.Println(currentFolder)
	files, _ := ioutil.ReadDir(currentFolder)
	for _, f := range files {
		fileName := f.Name()
		if fileName == "main.go" || fileName == "backups" || fileName == "new" {
			continue
		}
		replaceFile(currentFolder, fileName)
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	moduleName = "vtt-update"
)

func main() {
	pwd, _ := os.Getwd()
	currentFolder := fmt.Sprintf("%s/%s", pwd, moduleName)
	fmt.Println(currentFolder)
	files, _ := ioutil.ReadDir(currentFolder)
	for _, f := range files {
		fmt.Println(f.Name())
	}
	//
	//fmt.Println(filePath)
	////创建单词=>频次map
	//frequencyForWord := map[string]int{} // 与:make(map[string]int)相同
	////读取文件内容
	//for _, filename := range commandLineFiles([]string{filePath}) {
	//	//更新每个单词的频次
	//	updateFrequencies(filename, frequencyForWord)
	//}
	////打印单词=>频次
	//reportByWords(frequencyForWord)
	////反转单词=>频次 为 频次=>单词（多个）
	//wordsForFrequency := invertStringIntMap(frequencyForWord)
	////打印频次=>单词（多个）
	//reportByFrequency(wordsForFrequency)
}

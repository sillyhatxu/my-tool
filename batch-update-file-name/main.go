package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func TwoNumberSum(array []int, target int) []int {
	check := make(map[int]int)
	for _, value := range array {
		index, ok := check[value]
		if ok {
			return []int{array[index], value}
		}
		array[target-value] = index
	}
	return []int{-1, -1}
}

func main() {
	pathname := "/Users/shikuanxu/Downloads/新阴阳魔界"
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		panic(err)
	}
	for _, fi := range rd {
		if fi.IsDir() {
			continue
		} else {
			newName := fi.Name()
			//newName = strings.ReplaceAll(newName, "英语听力 - Lesson ", "")
			//newName = newName[0:2] + ".mp3"
			//newName = strings.ReplaceAll(newName, "[BD影视分享bd-film.cc]The.Boys.2019.", "")
			//newName = strings.ReplaceAll(newName, ".HD1080P.官方中字", "")
			//newName = strings.ReplaceAll(newName, "[BD影视分享bd-film.cc]黑袍纠察队.第二季.The.Boys.", "")
			//newName = strings.ReplaceAll(newName, ".HD720P.中字.人人.", "")
			//newName = newName[0:6] + ".mp4"
			newName = strings.ToUpper(newName[23:29]) + ".mp4"
			fmt.Println(newName)
			oldPath := fmt.Sprintf("%s/%s", pathname, fi.Name())
			newPath := fmt.Sprintf("%s/%s", pathname, newName)
			fmt.Println(fmt.Sprintf("oldPath : %s", oldPath))
			fmt.Println(fmt.Sprintf("newPath : %s", newPath))
			err := os.Rename(oldPath, newPath)
			if err != nil {
				panic(err)
			}
		}
	}
}

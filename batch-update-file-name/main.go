package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	pathname := "/Users/shikuanxu/Downloads/xxxxxxx/xxxxxxx"
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		panic(err)
	}
	for _, fi := range rd {
		if fi.IsDir() {
			continue
		} else {
			if !strings.Contains(fi.Name(), "mp41") {
				continue
			}
			newName := strings.ReplaceAll(fi.Name(), " ", "")
			newName = newName[:len(newName)-1]
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

package main

import (
	"fmt"
	"os"
)

/**
 * @Project GoDemo
 * @File    08_fileWriteAdd.go
 * @Author  Augus Lee
 * @Date    2022/10/17 17:30
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {
	f, err := os.OpenFile("../file/lines.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}

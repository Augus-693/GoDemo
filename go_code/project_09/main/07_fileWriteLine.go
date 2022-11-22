package main

import (
	"fmt"
	"os"
)

/**
 * @Project GoDemo
 * @File    07_fileWriteLine.go
 * @Author  Augus Lee
 * @Date    2022/10/17 17:28
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {
	f, err := os.Create("../file/lines.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.",
		"It is easy to learn Go."}
	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}

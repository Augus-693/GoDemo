package main

import (
	"fmt"
	"io"
	"os"
)

/**
 * @Project GoDemo
 * @File    02_fileReadAll.go
 * @Author  Augus Lee
 * @Date    2022/10/17 13:49
 * @Version 1.0
 */

func main() {
	/*
		//使用 os.ReadFile 直接指定文件名读取
		content, err := os.ReadFile("./test.txt")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(content))
		//文件的 Open 和 Close 被封装到 ReadFile 函数内部
		//所以不需要显式的 Open 文件和显式的 Close 文件
	*/

	/*
		//使用 ioutal.ReadFile 直接指定文件名读取
		//As of Go 1.16, this function simply calls os.ReadFile.
		content, err := ioutil.ReadFile("./test.txt")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(content))
	*/

	/*
		//先创建句柄再读取: os.Open
		file, err := os.Open("test.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		content, err := io.ReadAll(file)
		fmt.Println(string(content))
	*/

	//直接使用 os.OpenFile，只是要多加两个参数
	file, err := os.OpenFile("../file/test.txt", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	fmt.Println(string(content))

}

package main

import (
	"fmt"
	"os"
)

/**
 * @Project GoDemo
 * @File    05_fileWriteString.go
 * @Author  Augus Lee
 * @Date    2022/10/17 17:24
 * @GoVersion go1.19.2
 * @Version 1.0
 */

/**
 * @Author
 * @Description //TODO
 * @Date 14:48 2022/10/18
 * @Param
 * @return
 **/

func main() {
	//创建文件
	//如果这个文件已经存在，那么 create 函数将截断这个文件。该函数返回一个文件描述符
	f, err := os.Create("../file/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//将字符串写入文件
	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

package main

import (
	"fmt"
	"os"
)

/**
 * @Project GoDemo
 * @File    01_fileOpenAndCloseDemo.go
 * @Author  Augus Lee
 * @Date    2022/10/17 12:34
 * @Version 1.0
 */

func main() {
	//概念说明：file的叫法
	//1.file 叫 file对象
	//2.file 叫 file指针
	//3.file 叫 file文件句柄
	//打开文件
	file, err := os.Open("./src/go_code/project_09/file/test.txt") //相对路径：用Goland运行
	//file, err := os.Open("../file/test.txt") //相对路径：用Terminal运行
	//file, err := os.Open("E:\\Code\\GoDemo\\src\\go_code\\project_09\\file\\test.txt") //绝对路径
	if err != nil {
		fmt.Println("open file failed:", err)
	}
	//输出文件,file就是一个指针 *File
	fmt.Printf("file = %v\n", file)
	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file failed:", err)
	}
}

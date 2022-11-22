package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
 * @Project GoDemo
 * @File    10_fileCopy.go
 * @Author  Augus Lee
 * @Date    2022/10/18 12:29
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open srcFile failed:", err)
	}
	defer srcFile.Close()
	//通过 srcFile ，获取到 Reader
	reader := bufio.NewReader(srcFile)
	//打开 dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open dstFile failed:", err)
		return
	}
	//通过 dstFile , 获取到 Writer
	writter := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writter, reader)
}

func main() {
	//将 flower.jpg 文件拷贝到 grass.jpg
	//调用 CopyFile 完成文件拷贝
	srcFile := "../file/flower.jpg"
	dstFile := "../file/grass.jpg"
	_, err := CopyFile(srcFile, dstFile)
	if err != nil {
		fmt.Printf("Copy Successfully\n")
	} else {
		fmt.Printf("Copy Failed: %v\n", err)
	}
}

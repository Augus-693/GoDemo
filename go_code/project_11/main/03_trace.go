package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

/**
 * @Project GoDemo
 * @File    03_trace.go
 * @Author  Augus Lee
 * @Date    2022/10/19 10:35
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {

	//创建trace文件
	f, err := os.Create("../file/trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//server
	fmt.Println("Hello World")
}

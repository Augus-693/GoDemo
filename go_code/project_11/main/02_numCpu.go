package main

import (
	"fmt"
	"runtime"
)

/**
 * @Project GoDemo
 * @File    02_numCpu.go
 * @Author  Augus Lee
 * @Date    2022/10/18 20:07
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {
	//获取当前系统CPU数量
	num := runtime.NumCPU()
	fmt.Println("当前系统CPU数量为：", num)
	runtime.GOMAXPROCS(num) //设置num-1的CPU运行go程序
}

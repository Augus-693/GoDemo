package main

import (
	"fmt"
	"unsafe"
)

/**
 * @Project GoDemo
 * @File 05_integer.go
 * @Author Augus Lee
 * @Date 2022/9/20 17:31
 * @Version 1.0
 */

func main() {
	var n1 = 100 // ? n1 是什么类型
	//这里我们给介绍一下如何查看某个变量的数据类型
	//fmt.Printf() 可以用于做格式化输出。
	fmt.Printf("n1 的 类型 %T \n", n1)

	var n2 int64 = 10
	//unsafe.Sizeof(n1) 是unsafe包的一个函数，可以返回n1变量占用的字节数
	fmt.Printf("n2 的 类型 %T  n2占用的字节数是 %d ", n2, unsafe.Sizeof(n2))
}

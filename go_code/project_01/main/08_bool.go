package main

import (
	"fmt"
	"unsafe"
)

/**
 * @Project GoDemo
 * @File 08_bool.go
 * @Author Augus Lee
 * @Date 2022/9/20 18:59
 * @Version 1.0
 */

func main() {
	var b = false
	fmt.Println("b = ", b)
	//bool类型占用存储空间是一个字节
	fmt.Println("b的占用空间为", unsafe.Sizeof(b))
}

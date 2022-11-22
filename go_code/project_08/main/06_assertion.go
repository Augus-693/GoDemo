package main

import "fmt"

/**
 * @Project GoDemo
 * @File    06_assertion.go
 * @Author  Augus Lee
 * @Date    2022/10/16 13:16
 * @Version 1.0
 */

var x interface{}

func main() {
	var b float32 = 1.1
	x = b
	y := x.(float32)
	fmt.Printf("y 的类型是 %T 值是 %v\n", y, y)
	//类型断言，带检测机制
	var b2 float32 = 2.1
	x = b2
	if y, ok := x.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf("y 的类型是 %T 值是 %v\n", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("后续代码")
}

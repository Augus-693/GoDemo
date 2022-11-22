package main

import "fmt"

/**
 * @Project GoDemo
 * @File 03_logicalOperatorDetail.go
 * @Author Augus Lee
 * @Date 2022/9/26 15:12
 * @Version 1.0
 */

/*
逻辑运算符的使用细节
*/
//声明一个测试函数
func test() bool {
	fmt.Println("test.......")
	return true
}
func main() {
	var i int = 10
	//因为 i < 9 为 false，所以后面的 tset() 就不执行
	if i < 9 && test() {
		fmt.Println("ok")
	}
	////因为 i > 9 为 true，所以后面的 tset() 执行
	if i > 9 && test() {
		fmt.Println("ok......")
	}
	//因为 i > 9 为 true，所以后面的 tset() 就不执行
	if i > 9 || test() {
		fmt.Println("hello")
	}
	//因为 i < 9 为 false，所以后面的 tset() 执行
	if i < 9 || test() {
		fmt.Println("hello......")
	}
}

package main

import "fmt"

/**
 * @Project GoDemo
 * @File 05_anonymousFuncDemo.go
 * @Author Augus Lee
 * @Date 9/29/2022 8:37 PM
 * @Version 1.0
 */

func main() {
	//在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1 = ", res1)
	//将匿名函数赋给一个变量(函数变量)，再通过该变量来调用匿名函数
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(90, 30)
	fmt.Println("res2 = ", res2)
	res3 := a(80, 30)
	fmt.Println("res3 = ", res3)
}

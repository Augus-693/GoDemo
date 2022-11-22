package main

import "fmt"

/**
 * @Project GoDemo
 * @File 04_assigningOperator.go
 * @Author Augus Lee
 * @Date 2022/9/26 16:04
 * @Version 1.0
 */
/*
赋值运算符
*/
func main() {
	var i int
	i = 10 //基本赋值
	fmt.Println("i = ", i)

	//有两个变量a和b，对其值进行替换，并打印
	a := 9
	b := 10
	fmt.Printf("交换前：a = %v,b = %v\n", a, b)
	//定义一个临时变量t
	t := a
	a = b
	b = t
	fmt.Printf("交换后：a = %v,b = %v\n", a, b)
	//不设置临时变量替换值
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("再次交换后：a = %v,b = %v\n", a, b)
	//复合赋值操作
	i += 7
	fmt.Printf("i = %v\n", i)
}

package main

import "fmt"

/**
 * @Project GoDemo
 * @File 11_newDemo.go
 * @Author Augus Lee
 * @Date 2022/10/10 13:46
 * @Version 1.0
 */

func main() {
	num1 := 100
	fmt.Printf("num1的类型：%T num1的值 = %v num1的地址：%v\n", num1, num1, &num1)

	num2 := new(int)
	*num2 = 100
	fmt.Printf("num2的类型：%T num2的值 = %v num2的地址：%v\nnum2指针指向的值 = %v\n", num2, num2, &num2, *num2)
}

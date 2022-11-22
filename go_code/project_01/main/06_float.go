package main

import (
	"fmt"
	"reflect"
)

/**
 * @Project GoDemo
 * @File 06_float.go
 * @Author Augus Lee
 * @Date 2022/9/20 17:40
 * @Version 1.0
 */

func main() {
	var price float32 = 5.83
	fmt.Println("price = ", price)

	var num1 float32 = -0.00089
	var num2 float64 = -5151515.05
	fmt.Println("num1 = ", num1, " num2 = ", num2)

	var num3 float32 = -123.0000901 //精度损失
	var num4 float64 = -123.0000901 //精度没有损失
	fmt.Println("num3 = ", num3, " num4 = ", num4)

	//Golang浮点型默认声明为float64 类型
	var num5 = 0.01
	fmt.Println("num5的数据类型是", reflect.TypeOf(num5))
	fmt.Printf("num5的数据类型是 %T \n", num5)

	//浮点型常量表现形式
	//十进制形式
	num6 := 5.20
	num7 := .132 //相当于0.132
	fmt.Println("num6 = ", num6, " num7 = ", num7)
	//科学计数法形式
	num8 := 1.455112e3
	num9 := 1.455112e3
	fmt.Println("num8 = ", num8, " num9 = ", num9)

}

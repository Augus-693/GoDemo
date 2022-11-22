package main

import (
	"fmt"
	"strconv"
)

/**
 * @Project GoDemo
 * @File 12_stringTran.go
 * @Author Augus Lee
 * @Date 2022/9/23 14:05
 * @Version 1.0
 */

func main() {

	var num1 int = 99
	var num2 float64 = 45.652
	var b bool = true
	var myChar byte = 'h'
	var str string

	//方式1： 使用 fmt.Sprintf转换
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T, str = %q\n", str, str)
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T, str = %q\n", str, str)
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T, str = %q\n", str, str)
	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T, str = %q\n", str, str)

	//方式2：使用 strconv 包的函数
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type %T, str = %q\n", str, str)
	str = strconv.Itoa(num1) //直接将int转为string
	fmt.Printf("str type %T, str = %q\n", str, str)
	str = strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Printf("str type %T, str = %q\n", str, str)
	str = strconv.FormatBool(b)
	fmt.Printf("str type %T, str = %q\n", str, str)
	str = strconv.FormatInt(int64(myChar), 10)
	fmt.Printf("str type %T, str = %q\n", str, str)
}

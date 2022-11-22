package main

import (
	"fmt"
	"strconv"
)

/**
 * @Project GoDemo
 * @File 13_stringTran2.go
 * @Author Augus Lee
 * @Date 2022/9/23 14:52
 * @Version 1.0
 */

func main() {
	var str1 string = "true"
	var b bool
	//strconv.ParseBool(str)函数会返回两个值(value bool, err error)
	//因为只要获取value，所以用_忽略err
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("b type %T, b = %v\n", b, b)

	var str2 string = "123465789"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T, n1 = %v\n", n1, n1)
	fmt.Printf("n1 type %T, n2 = %v\n", n2, n2)

	var str3 string = "135.5475"
	var f float64
	f, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f type %T, f = %v\n", f, f)
}

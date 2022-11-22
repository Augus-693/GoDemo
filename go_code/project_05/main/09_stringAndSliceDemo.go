package main

import "fmt"

/**
 * @Project GoDemo
 * @File 09_stringAndSliceDemo.go
 * @Author Augus Lee
 * @Date 2022/10/12 16:03
 * @Version 1.0
 */

func main() {

	//string 也可以进行切片处理
	str := "hello&Golang"
	slice := str[6:]
	fmt.Println("slice: ", slice) //Golang

	//string -> []byte -> 修改 -> 重写转成 string
	//只能处理英文和数字
	arr1 := []byte(str)
	arr1[0] = 'z'
	str1 := string(arr1)
	fmt.Println("str1: ", str1)

	//string -> []rune -> 修改 -> 重写转成 string
	//可以处理英文，数字和中文
	arr2 := []rune(str)
	arr2[0] = '李'
	str2 := string(arr2)
	fmt.Println("str2: ", str2)
}

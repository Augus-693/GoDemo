package main

import (
	"fmt"
)

/**
 * @Project GoDemo
 * @File 06_forDemo.go
 * @Author Augus Lee
 * @Date 9/28/2022 5:50 PM
 * @Version 1.0
 */

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello Golang", i)
	}

	//将变量初始化和变量迭代写到其他位置
	j := 0
	for j < 10 {
		fmt.Println("Hello Golang", j)
		j++
	}

	//等价for ; ; {} 是一个无限循环， 通常需要配合break 语句使用
	k := 0
	for {
		if k < 10 {
			fmt.Println("Hello Golang", k)
		} else {
			break
		}
		k++
	}

	//Golang 提供for-range 的方式，可以方便遍历字符串和数组
	//字符串遍历方式1——传统方式
	var str string = "hello,golang!"
	for l := 0; l < len(str); l++ {
		fmt.Printf("%c \n", str[l])
	}

	//字符串遍历方式1——传统方式
	str1 := "hello,golang语言"
	str2 := []rune(str1) //把str转成[]rune 切片解决中文乱码问题
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}

	//字符串遍历方式2——for-range
	str3 := "abc~ok"
	for index, val := range str3 {
		fmt.Printf("index = %d, val = %c \n", index, val)
	}
}

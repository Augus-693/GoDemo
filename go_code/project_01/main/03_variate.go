package main

import "fmt"

/**
 * @Project GoDemo
 * @File 03_variate.go
 * @Author Augus Lee
 * @Date 2022/9/20 15:44
 * @Version 1.0
 */

func main() {
	//Golang的变量使用方式

	//第一种，指定变量类型，声明后不赋值，使用默认值（0）
	var i int //声明变量
	i = 10    //变量赋值
	fmt.Println("i = ", i)

	//第二种指定值自动判断变量类型
	var num = 10.11
	fmt.Println("num = ", num)

	//第三种：省略 var。注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误
	//下面的方式等价于 var name string    name = "tom"
	name := "tom"
	fmt.Println("name = ", name)
}

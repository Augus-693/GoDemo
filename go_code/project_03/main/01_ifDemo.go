package main

import "fmt"

/**
 * @Project GoDemo
 * @File 01_ifDemo.go
 * @Author Augus Lee
 * @Date 2022/9/27 14:21
 * @Version 1.0
 */

func main() {
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanf("%d", &age)
	if age >= 18 {
		fmt.Println("你的年龄大于18岁，需要对自己的行为负责")
	}

	//golang支持在if中，直接定义一个变量
	if i := 10; i > 0 {
		fmt.Println("i是一个正数")
	}

}

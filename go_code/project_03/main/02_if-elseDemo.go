package main

import "fmt"

/**
 * @Project GoDemo
 * @File 02_if-elseDemo.go
 * @Author Augus Lee
 * @Date 2022/9/27 14:29
 * @Version 1.0
 */
func main() {
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanf("%d", &age)
	if age >= 18 {
		fmt.Println("你的年龄大于18岁，需要对自己的行为负责")
	} else {
		fmt.Println("你的年龄未满18周岁")
	}
}

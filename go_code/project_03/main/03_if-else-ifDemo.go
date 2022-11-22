package main

import "fmt"

/**
 * @Project GoDemo
 * @File 03_if-else-ifDemo.go
 * @Author Augus Lee
 * @Date 2022/9/27 14:34
 * @Version 1.0
 */

func main() {
	var score int
	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)

	if score == 100 {
		fmt.Println("非常棒的成绩")
	} else if score >= 80 && score <= 99 {
		fmt.Println("成绩优秀")
	} else if score >= 60 && score < 80 {
		fmt.Println("成绩中等")
	} else if score >= 0 && score < 60 {
		fmt.Println("成绩不合格")
	} else {
		fmt.Println("数据非法")
	}
}

package main

import "fmt"

/**
 * @Project GoDemo
 * @File 04_loopNesting.go
 * @Author Augus Lee
 * @Date 2022/9/27 14:44
 * @Version 1.0
 */

func main() {
	var second float64
	fmt.Println("请输入秒数")
	fmt.Scanln(&second)

	if second <= 8 {
		fmt.Println("恭喜你进入决赛")
		var gender string
		fmt.Println("请输入性别：")
		fmt.Scanln(&gender)
		if gender == "男" {
			fmt.Println("男子组")
		} else {
			fmt.Println("女子组")
		}
	} else {
		fmt.Println("很遗憾，你已被淘汰")
	}
}

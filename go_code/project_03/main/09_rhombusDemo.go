package main

import "fmt"

/**
 * @Project GoDemo
 * @File 09_rhombusDemo.go
 * @Author Augus Lee
 * @Date 9/29/2022 8:19 PM
 * @Version 1.0
 */

func main() {
	// 长
	var x int
	// 宽
	var y int
	// 行数
	fmt.Scanf("请输入菱形的长和宽：%d %d\n", &x, &y)
	row := 1
	for row <= y {
		// 计算每行得星星数
		count := 0
		if row <= (y/2 + 1) {
			count = (2*row - 1)
		} else {
			count = 2*(y-row) + 1
		}
		row++
		text := ""
		// 算出显示星星的范围
		star_min := ((x - count) / 2) + 1
		star_max := ((x - count) / 2) + count
		for index := 1; index <= x; index++ {
			if index >= star_min && index <= star_max {
				text += "*"
			} else {
				text += " "
			}
		}
		fmt.Println(text)
	}
}

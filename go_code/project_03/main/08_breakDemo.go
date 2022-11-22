package main

import "fmt"

/**
 * @Project GoDemo
 * @File 08_breakDemo.go
 * @Author Augus Lee
 * @Date 9/28/2022 8:44 PM
 * @Version 1.0
 */
func main() {
	//指定标签形式使用 break
lable1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break lable1
			}
			fmt.Println("j = ", j)
		}
	}
}

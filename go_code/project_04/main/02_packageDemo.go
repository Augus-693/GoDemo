package main

import (
	"fmt"
	"github.com/GoDemo/go_code/project_04/util"
)

/**
 * @Project GoDemo
 * @File 02_packageDemo.go
 * @Author Augus Lee
 * @Date 9/28/2022 9:27 PM
 * @Version 1.0
 */
func main() {
	var n1 float64
	var n2 float64
	var operator byte
	fmt.Println("输入数值:（格式：数字 符号 数字）")
	fmt.Scanf("%f %c %f\n", &n1, &operator, &n2)
	result := util.Cal(n1, n2, operator)
	fmt.Printf("result = %.2f\n", result)
}

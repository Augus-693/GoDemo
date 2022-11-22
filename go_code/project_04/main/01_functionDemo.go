package main

import "fmt"

/**
 * @Project GoDemo
 * @File 01_functionDemo.go
 * @Author Augus Lee
 * @Date 9/28/2022 9:10 PM
 * @Version 1.0
 */

func cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号错误")
	}
	return res
}

func main() {
	var n1 float64
	var n2 float64
	var operator byte
	fmt.Println("输入数值:（格式：数字 符号 数字）")
	fmt.Scanf("%f %c %f\n", &n1, &operator, &n2)
	result := cal(n1, n2, operator)
	fmt.Printf("result = %.2f\n", result)
}

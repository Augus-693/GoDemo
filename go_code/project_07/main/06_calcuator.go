package main

import "fmt"

/**
 * @Project GoDemo
 * @File 06_calcuator.go
 * @Author Augus Lee
 * @Date 2022/10/14 13:45
 * @Version 1.0
 */

type Calcuator struct {
	Num1 float64
	Num2 float64
}

// 实现形式1：用四个方法完成
func (calcuator Calcuator) getSum() float64 {
	return calcuator.Num1 + calcuator.Num2
}
func (calcuator *Calcuator) getSub() float64 {
	return calcuator.Num1 - calcuator.Num2
}
func (calcuator *Calcuator) getMul() float64 {
	return calcuator.Num1 * calcuator.Num2
}
func (calcuator *Calcuator) getDiv() float64 {
	return calcuator.Num1 / calcuator.Num2
}

// 实现形式2：用一个方法完成
func (calcuator *Calcuator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 - calcuator.Num2
	case '*':
		res = calcuator.Num1 * calcuator.Num2
	case '/':
		res = calcuator.Num1 / calcuator.Num2
	}
	return res
}

func main() {
	var calcuator Calcuator
	calcuator.Num1 = 10.2
	calcuator.Num2 = 5.1
	fmt.Printf("%.2f %.2f %.2f %.2f\n", calcuator.getSum(), calcuator.getSub(), calcuator.getMul(), calcuator.getDiv())

	res := calcuator.getRes('*')
	fmt.Printf("%.2f\n", res)
}
